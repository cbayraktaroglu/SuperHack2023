package blckDecoder

import (
	"bufio"
	"context"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BLCKDecoder struct {
	nodeLookup map[int]string
}

func NewBLCKDecoder(chainIDmap map[int]string) (decoder *BLCKDecoder) {
	decoder = &BLCKDecoder{
		nodeLookup: chainIDmap,
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
