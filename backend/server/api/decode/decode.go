package decode

import (
	blckDecoder "SuperHack2023/decoder"
	file "SuperHack2023/server/api/decode/contracts"
	decodeStructs "SuperHack2023/server/api/decode/structs"
	"SuperHack2023/server/responses"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func PrivHandler(decoder *blckDecoder.BLCKDecoder) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawData, err := c.GetRawData()

		if err != nil {
			log.Println("Error getting the raw data: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		privateFileManifest := decodeStructs.PrivateFile{}

		err = json.Unmarshal(rawData, &privateFileManifest)

		if err != nil {
			log.Println("Error unmarshaling the data: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		returnBuffer := []byte{}

		for _, txInfo := range privateFileManifest.TxInfo {
			txData, txErr := decoder.GetBytesFromTx(txInfo.ChainID, txInfo.TxHash)

			if txErr != nil {
				log.Println("Unable to get the tx data: ", txErr)
				responses.FailResponse(c, http.StatusBadRequest)
				return
			}

			returnBuffer = append(returnBuffer, txData...)
		}

		//Check Checksum
		modifiedBuffer := returnBuffer[:len(returnBuffer)-1]
		checkSum, err := decoder.GetFileCheckSumSHA256(modifiedBuffer)
		if err != nil {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Unable to get GetFileCheckSumSHA256: ", err)
			return
		}
		if privateFileManifest.CheckSum != checkSum {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Unable to get same checksum:")
			return
		}

		log.Println("success")
		log.Println(modifiedBuffer)
		// Base64 encode the binary data
		//encodedData := base64.StdEncoding.EncodeToString(modifiedBuffer)
		responses.SuccessResponse(c, modifiedBuffer)
	}
}

func PublicHandler(decoder *blckDecoder.BLCKDecoder) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("I am inside")
		scAddress := c.Param("address")
		chainName := c.Param("chainName")

		log.Println(scAddress)
		log.Println(chainName)

		instance, err := decoder.GetFileContractInstance(scAddress, chainName)
		if err != nil {
			log.Println("Unable to get the instance: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		fileName, err := instance.GetFileName(nil)
		if err != nil {
			log.Println("Unable to get the fileName: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		fileType, err := instance.GetFileType(nil)
		if err != nil {
			log.Println("Unable to get the fileType: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		completeFileName := fileName + fileType

		fileCheckSum, err := instance.GetFileCheckSum(nil)
		if err != nil {
			log.Println("Unable to get the fileCheckSum: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		numTransactions, err := instance.GetNumberofTransactions(nil)
		if err != nil {
			log.Println("Unable to get the numTransactions: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		privateFileManifest := decodeStructs.PrivateFile{}

		privateFileManifest.FileType = fileType
		privateFileManifest.FileName = fileName
		privateFileManifest.CheckSum = fileCheckSum
		privateFileManifest.TxInfo, err = createTxInfoData(numTransactions, instance)
		if err != nil {
			log.Println("Unable to get createTxInfoDataa: ", err)
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		returnBuffer := []byte{}

		for _, txInfo := range privateFileManifest.TxInfo {
			txData, txErr := decoder.GetBytesFromTx(txInfo.ChainID, txInfo.TxHash)

			if txErr != nil {
				log.Println("Unable to get the tx data: ", err)
				responses.FailResponse(c, http.StatusBadRequest)
				return
			}

			returnBuffer = append(returnBuffer, txData...)
		}

		//Check checksum
		modifiedBuffer := returnBuffer[:len(returnBuffer)-1]
		checkSum, err := decoder.GetFileCheckSumSHA256(modifiedBuffer)
		if err != nil {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Unable to get GetFileCheckSumSHA256: ", err)
			return
		}
		if privateFileManifest.CheckSum != checkSum {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Unable to get same checksum: ", err)
			return
		}

		info := decodeStructs.PublicInfo{}
		info.ReturnBuffer = modifiedBuffer
		info.FileName = completeFileName

		responses.SuccessResponse(c, info)
	}
}

func createTxInfoData(numTransactions *big.Int, instance *file.BlockchainFileCaller) ([]decodeStructs.TxData, error) {

	var txDataList []decodeStructs.TxData
	num := numTransactions.Int64()

	for i := int64(0); i < num; i++ {
		var txData decodeStructs.TxData
		tx, err := instance.GetTransactionAtIndex(nil, new(big.Int).SetInt64(i))
		if err != nil {
			return nil, err
		}
		txData.ChainID, err = strconv.Atoi(tx.ChainID)
		if err != nil {
			return nil, err
		}
		txData.TxHash = tx.ChainID

		txDataList = append(txDataList, txData)
	}

	return txDataList, nil
}
