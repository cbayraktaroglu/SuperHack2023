package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://polygon-mumbai-infura.wallet.coinbase.com")

	if err != nil {
		log.Fatal(err)
	}

	/*
		txList := []string{
			"0xe07b113180b49a33a35476f9526475c29a20ff7458c1a387b98df8655c4f3cb1",
			"0x79769390557717e2faf3b5ef07dc8d52db09f3517a8f7ab2c38978a2d9d0ed47",
			"0xb9d8c8823cc3ca0571333369dc04ccb53857ef41fe60210b9551acf482d0003a",
			"0x7be28adb06ef310591038c9dd10ae5eeaec1bd3b820d60bbac1f90805e608536",
			"0xad54ca673a2db59678df0e738f11d139c7f67a973dd0fd0970774b0fc53c510b",
			"0x14aaa5c9bb95645fa9a41f757907fe7087f3dde609719b24f33acefdb6d86de4",
			"0xb2007d863cbaf12061685799058e96f11927cf4bc0482141ee7bfe2d3b4b2918",
			"0xe9d79090914bb19fc4e963496883e5c527d4de8b1955828fccc05ee56f3cdd24",
			"0x1c6bb100f0ff9c1a34fba240e9f95da05ce83d58895c0b95fdd8e022d9f46444",
			"0x63dab991d512f2ab57c55e881b9981b53793a6b846b621c1459aa86679450dcf",
			"0x53908f1feabd12ecad4d08c47409cfa3f15a8275c45d5199536b50dde36ee233",
			"0x7116b4eae8ccfe62f6f66681626ffa46315af1e43b7bf7b351822592536d64c8",
			"0x4415091cb8c69fae90004b71ead824d12b6b4e9cce5f2238f30749906eea55a7",
			"0x986a60e1d387090fb22292037814edf670161177ab6941b6dc9146b854255d59",
			"0x246701d99d3f3f055e43bd29fd5524156e811c3c1380d2ab1892f3c7e0a0603c",
			"0x1bb7c7ae88edb05b8d4e19438e5c7bbad09755dbb72110b7a953b59016cbabf2",
			"0x6f256dfdef8efcd48bafb04b5718c6761d822e8ca7742a124c43719fbf99fa51",
			"0xd4ee2214db65eb1d82d3520eb44427264da87c895565f0f103ac32850075f6d0",
			"0xd7cfe5d59d5e038e40375a915f96675a83693882b38449adf3149252344b20da",
			"0x957a395389fa3b1bfb30e2d3f86dfb962bb3707d40b84386ad8f8368649a7122",
		}
	*/
	txList := []string{
		"0x8d2d831fb7f37430a38cccd8cf9fd7afda742d3f9ac2e8ce1d5e752578f07cbf",
		"0x114408030cfdd0e0494ccdf6e59ade4ecfa92daca35ccbbaf95ac2dbc71326d8",
		"0xc10f0e1f39975330e30a7967e03bda93b43fd52c2198154aea5743c5b698f88d",
		"0x44eae0cfa92f554aaa783caf31c661d96fd43d0f7c810653de01be9da9856328",
		"0x244da1d5ba96b6d727cb19e60e2fe10195da639293394392d2a42217ae2f39eb",
		"0x010e04812181d0609bbd2806cc02ead46ded6c57151fceadb1cb94b3a9e75d41",
		"0xf82ab979742135cf39732a36bbceb3268d64d5eb1c2986f3fdedfbec238fad12",
		"0x52f7a51ed749d3bce44bcef6cc8af97749824bc2c826fc7f300a0e460409386a",
		"0xa5ff442deb22eedff59dddaa8e74c2474714b1fb1de23f6a9e595f2fa45afab1",
		"0x0f85e39f26f542813f4cd385a7a0d3a85c7c747cb26271b7e9c8d7cceade5012",
		"0xabf20c58270624b25c817b98c1127603ae5d45431dfbeb11214501c0713146c2",
		"0xa14361ff2b6e4c181c8215f3b90f1b38d49ede1ed4ba58852411e946cb0b498b",
		"0x997778cdc40b6d4e34f305c4eda9ba263185401e5ef5effe1dcf120e54a71645",
		"0x7b99d2c844a2aa71acdac95c99504abf55a36c09be55e9afeefe5313ad71ebe2",
		"0x4cb27c578b2fea836586450ac0248fec6f5c61a103f5241962c0b3e6fb0f18ae",
		"0x961b518753b3be3839ffa9edd26a8473c574b4753219150344e3f101b556bd84",
		"0x9d900a1d9ccbfcf41f4c64d77a92a1ce05c0646a7a59e7989927d94434642436",
		"0xbedaadd4c4a568565b9f8abc4204a1eaad4c0dac0b5bd5f2adbd3710a2f2e2af",
		"0xe37e39c2b4bcfd1beb9c0c9d43894fbeaab2cd018c2dbc6bf53af06a70a2eb9c",
		"0xe730ac9e247cd18b286ea852e943dcd93f422b50034f4e25cda8231867f857cf",
	}

	targetFile, err := os.Create("./output/index.html")

	if err != nil {
		log.Print("Error while crating the file: ", err)
		return
	}

	writer := bufio.NewWriter(targetFile)

	for _, hex := range txList {
		theTx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hex))

		if err != nil {
			log.Fatal(err)
		}

		nn, err := writer.Write(theTx.Data())

		log.Println(nn, " ", err)
	}

	writer.Flush()
}
