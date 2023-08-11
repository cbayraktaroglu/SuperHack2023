package main

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os"
	"strings"

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
	/*
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
	*/
	//txList := strings.Split("0xc090ccfaab38f374ce2109d1f84534fde6354e782714fac15b60179ed82bf16f 0x4bb0f8e3bf27be18ac66520bc8c329f0501cfda7320623c1369b54d42d580487 0x359bccd3aa8e57a852b8f72329896d2db7967432f6d7171795d8c4ca210f8a55 0x812c7c613f48d369118aecec6af7dd6bf77258ab8cecd952bec38628fac8292f 0xc31d225218908eed76e14cf38533c6d8f1bfbbca091579011a5f887d6775bb00 0x31fb2c3093fd7294e83ac96f2d71afdcef5cdf920ba7aa5e7f72ed51066b8a91 0xd1e77cdaaed335eace6ad3b70e5dd34710fd1ff9f9d6c1b7334c2cc1ba7ae259 0x530ebccfdbda93200b99a1b880c57cb762aa7492bb7286b96ebf0ab2c6895b36 0x26b424804ddba9b3c1e829f272083be1e45b67614d87e0be0e65a922f1219742 0xb3ffb2191739848437e3c3a6ee4efdbf30ac128c4622ab2e22ba52a0eaa7c6ee 0xbbd0da01a5902f48257d70187aade31702fe66cf72e97043ce8bb6bf281a845c 0xba52a322d8e8b0e41fe95ef60c96d744951e45f6758bc197ad45c65971701c59 0xf4d99472ceb46e8440d0995d3b2ea83ef2db6a0d9c89076fd18db7704bf7a2d3 0xc67f1c3cc3b41881a565e5542e1218f77b44a1d5df290275aaa10db7d0716d95 0xbe5793412b510a9349a75d0f83d08ba6bc1032090f5ddf7cbcf2799c4425a097 0x37345059d89f217f8afee4a0b465196058367ca5cefad48df666e3c343bfcc72 0xe105181f3f036a7ca3140bf23291212cf6f159d5b4d2b8c7117900c0fec463b7 0x615b6c1a4d4604080f25147b83ead6b63a5c8f2d4136696d46904a445c017e12 0x6709a3f181dd07fc76cc5898db380fd981bd54ae1ba2411643f8f0be28a9dd5d 0xcd44b044ac862d45282aa95a6f86eb3945aa33d73dae09604a0750df54863464 0x868e977a091bc25be9e08d66bd27520254c89d03c54ff71e1ae73dab8a40a32c 0xf4cdc81da5079c92f6a06c88ea21d498911d4560f04d4363db649d6dd7cc567c 0x1c7c589fe1867471c9d0aa636f422a9b92203da5396e5786806ab25bbd1104d7 0x6cda67255cdf5c3a067eb3bcee4eb9e18a90d8e7b887df1d43388b79ad563408 0x7da2717cc3ef7f51ee1e5c656d9ca54413ec0e6d1f38136a987c3f42447ef929 0x2f61e59f7db93d8e9c5a4276e4e8930dabb69947c3689c21cab220c27453e24a", " ")
	//txList := strings.Split("0x234a25ce95a322f6098c8710399ee28d5e47679bdcad6704a12d0d9f06221d1d 0xc337867462cd56c6da258a2af90beb0aa5d3ae79abe48ac4f1b8c22924bb9214 0x454daed821b217ff2af4176258151e992d78902a3f96f0e62373aa6759a6754d 0x020be3516828b9b908f6aec02cd092b26d7388bdc475ff2abd2c1f2cfc1325a0 0x832c9ea37e13be6cf89da5ff6e617ce6b58348666228ec0703b0211ab10473ac", " ")
	//txList := strings.Split("0x17419ccdd4e3842b2f23fcbd4070a324e187396402dc3437448d1419aa28e3e8 0x7b319bfb0b41b0d078dc24152bd5ea12e91cebd5eb63940c56671034eb019639 0xfd862300d126c52f3f8e63f4f67e67020390a2830b22def3586e928c563140ed 0xa65a6b57d843426d829310cb48d211003c8af3de328ea2d89d335205002ea5d4 0x2bc0a74da6d845faca6a5bdcd9d4e2c4b23a22a686d09d16e1853b5435a0d542", " ")
	txList := strings.Split("0x251decded9409a3687817eb497d38dc83c34fb6a8f8588dfe71a2f2d8405acb9 0x89901f80db6639cdc5f70b0cc8b0ebeed2b55c47fba26514ce249dc609be0d18 0xfdfc589ea532dee1942f2ea37f2fdcbcfb578f551c6dc769e8b1dffdbc58fbd7 0xa33cae70c0758a422119f7ca91bb4bef794d5eddfa505ccb133eadc90cf5d9c1 0x179b05a2321f8686d0e7bbcd6f8afa9c133411084d541a9c0bcc5d85681c4528", " ")
	log.Print(len(txList))

	targetFile, err := os.Create("./output/uploaded.gif")

	if err != nil {
		log.Print("Error while crating the file: ", err)
		return
	}

	writer := bufio.NewWriter(targetFile)
	txDataBuffer := bytes.NewBuffer([]byte{})

	for _, hex := range txList {
		theTx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hex))

		if err != nil {
			log.Fatal(err)
		}

		txDataBuffer.Write(theTx.Data())
	}

	nn, err := writer.Write(txDataBuffer.Bytes()[:len(txDataBuffer.Bytes())-1])
	log.Println(nn, " ", err)
	writer.Flush()
}