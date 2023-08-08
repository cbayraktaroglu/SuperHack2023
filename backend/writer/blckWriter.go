package blckWriter

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BLCKWriter struct {
	privKey    *ecdsa.PrivateKey
	pubKey     *ecdsa.PublicKey
	pubAddress common.Address
	client     *ethclient.Client
}

func NewBLCKWriter(privateKey string, nodeUrl string) (writer *BLCKWriter, err error) {

	writer = &BLCKWriter{}

	// Connect to the client
	writer.client, err = ethclient.Dial(nodeUrl)

	if err != nil {
		return
	}

	// Initialize the private key
	writer.privKey, err = crypto.HexToECDSA(privateKey)

	if err != nil {
		return
	}

	// Initialize the public key
	publicKey := writer.privKey.Public()
	writer.pubKey = publicKey.(*ecdsa.PublicKey)

	writer.pubAddress = crypto.PubkeyToAddress(*writer.pubKey)

	return
}

func (blckWrtr *BLCKWriter) getAccountNonce() (nonce uint64, err error) {
	nonce, err = blckWrtr.client.PendingNonceAt(context.Background(), blckWrtr.pubAddress)
	return
}

func (blckWrtr *BLCKWriter) getSuggestedGasPrice() (price *big.Int, err error) {
	price, err = blckWrtr.client.SuggestGasPrice(context.Background())
	return
}

func (blckWrtr *BLCKWriter) getGasLimitEstimation(multiplier float64, toAddress common.Address, data []byte) (gasLimit uint64, err error) {
	gasEstimation, err := blckWrtr.client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})

	if err != nil {
		return
	}

	gasLimit = uint64(float64(gasEstimation) * multiplier)

	return
}

func (blckWrtr *BLCKWriter) CreateSignedTransaction(toAddress common.Address, data []byte) (signedTx *types.Transaction, err error) {
	// Get the nonce
	nonce, err := blckWrtr.getAccountNonce()

	if err != nil {
		return
	}

	// Get gas price
	gasPrice, err := blckWrtr.getSuggestedGasPrice()

	if err != nil {
		return
	}

	// Get the gas limit
	gasLimit, err := blckWrtr.getGasLimitEstimation(1.3, toAddress, data)

	if err != nil {
		return
	}

	// Set the value
	value := big.NewInt(0)

	// Get chainID
	chainID, err := blckWrtr.client.NetworkID(context.Background())

	if err != nil {
		return
	}

	// Create unsigned TX
	unsignedTx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// Sign the TX
	signedTx, err = types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), blckWrtr.privKey)

	return
}

func (blckWrtr *BLCKWriter) SendSignedTransaction(signedTx *types.Transaction) (txHash common.Hash, err error) {
	err = blckWrtr.client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		return
	}

	txHash = signedTx.Hash()

	return
}

func (blckWrtr *BLCKWriter) FileToByte(filePath string) (output []byte, err error) {
	// Init the output
	output = []byte{}

	// Open the file
	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	// Initialize the reader
	reader := bufio.NewReader(file)

	for {
		byteValue, readErr := reader.ReadByte()

		if readErr != nil {
			break
		}

		output = append(output, byteValue)
	}

	return
}
