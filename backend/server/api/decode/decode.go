package decode

import (
	blckDecoder "SuperHack2023/decoder"
	decodeStructs "SuperHack2023/server/api/decode/structs"
	"SuperHack2023/server/responses"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PrivHandler(decoder *blckDecoder.BLCKDecoder) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawData, err := c.GetRawData()

		if err != nil {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Error getting the raw data: ", err)
			return
		}

		privateFileManifest := decodeStructs.PrivateFile{}

		err = json.Unmarshal(rawData, &privateFileManifest)

		if err != nil {
			responses.FailResponse(c, http.StatusBadRequest)
			log.Println("Error unmarshaling the data: ", err)
			return
		}

		returnBuffer := []byte{}

		for _, txInfo := range privateFileManifest.TxInfo {
			txData, txErr := decoder.GetBytesFromTx(txInfo.ChainID, txInfo.TxHash)

			if txErr != nil {
				responses.FailResponse(c, http.StatusBadRequest)
				log.Println("Unable to get the tx data: ", err)
				return
			}

			returnBuffer = append(returnBuffer, txData...)
		}

		responses.SuccessResponse(c, returnBuffer)
	}
}
