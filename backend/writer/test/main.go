package main

import (
	blckWriter "SuperHack2023/writer"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	wrtr, _ := blckWriter.NewBLCKWriter("YOURKEY", "https://polygon-mumbai.gateway.tenderly.co")

	//TODO error handling
	root := "./input"
	files, _ := ioutil.ReadDir(root)

	txHashes := []common.Hash{}

	sort.Slice(files, func(i, j int) bool {
		splittedI := strings.Split(files[i].Name(), ".")
		splittedJ := strings.Split(files[j].Name(), ".")

		//TODO error handling
		numI, _ := strconv.Atoi(splittedI[0])
		numJ, _ := strconv.Atoi(splittedJ[0])

		return numI < numJ
	})

	for _, file := range files {
		fmt.Println("Processing ", file.Name())

		fileAsBytes, err := wrtr.FileToByte(root + "/" + file.Name())

		if err != nil {
			log.Fatal("Error processing file ", file.Name(), " : ", err)
		}

		signedTx, err := wrtr.CreateSignedTransaction(common.HexToAddress(""), fileAsBytes)

		if err != nil {
			log.Fatal("Error processing file ", file.Name(), " : ", err)
		}

		hash, err := wrtr.SendSignedTransaction(signedTx)

		if err != nil {
			log.Fatal("Error processing file ", file.Name(), " : ", err)
		}

		txHashes = append(txHashes, hash)
		time.Sleep(time.Second * 10)
	}

	fmt.Println(txHashes)
}
