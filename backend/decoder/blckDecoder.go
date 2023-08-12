package blckDecoder

import (
	file "SuperHack2023/server/api/decode/contracts"
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BLCKDecoder struct {
	nodeLookup      map[int]string
	chainNameLookup map[string]int
}

func NewBLCKDecoder(chainIDmap map[int]string, chainNameMap map[string]int) (decoder *BLCKDecoder) {
	decoder = &BLCKDecoder{
		nodeLookup:      chainIDmap,
		chainNameLookup: chainNameMap,
	}
	return
}

func (blckDcdr *BLCKDecoder) GetBytesFromTx(chainID int, txHash string) (data []byte, err error) {

	nodeEndpoint, isOk := blckDcdr.nodeLookup[chainID]

	if !isOk {
		err = errors.New("no nodeendpoint has been registered for the given chainID")
		return
	}

	blockchainClient, err := ethclient.Dial(nodeEndpoint)

	if err != nil {
		return
	}

	theTx, _, err := blockchainClient.TransactionByHash(context.Background(), common.HexToHash(txHash))

	if err != nil {
		return
	}

	data = theTx.Data()

	return
}

func (blckDcdr *BLCKDecoder) BytesToFile(data []byte, outputFilePath string) (err error) {

	targetFile, err := os.Create(outputFilePath)

	if err != nil {
		return
	}

	writer := bufio.NewWriter(targetFile)

	_, err = writer.Write(data)

	if err != nil {
		return
	}

	err = writer.Flush()

	return
}

func (blckDcdr *BLCKDecoder) GetFileContractInstance(scAddress string, chainName string) (caller *file.BlockchainFileCaller, err error) {

	chainID, isOk := blckDcdr.chainNameLookup[chainName]

	if !isOk {
		err = errors.New("no node endpoint has been registered for the given chainID")
		return
	}

	nodeEndpoint, isOk := blckDcdr.nodeLookup[chainID]

	if !isOk {
		err = errors.New("no node endpoint has been registered for the given chainID")
		return
	}

	blockchainClient, err := ethclient.Dial(nodeEndpoint)

	caller, err = file.NewBlockchainFileCaller(common.HexToAddress(scAddress), blockchainClient)
	if err != nil {
		return
	}
	return
}

func (blckEncd *BLCKDecoder) GetFileCheckSumSHA256(fileBuffer []byte) (checksum string, err error) {

	hash := sha256.New()

	_, err = io.Copy(hash, bytes.NewBuffer(fileBuffer))

	if err != nil {
		return
	}

	checksum = hex.EncodeToString(hash.Sum(nil))

	return
}
