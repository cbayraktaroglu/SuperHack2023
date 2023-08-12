package router

import (
	blckDecoder "SuperHack2023/decoder"
	"SuperHack2023/server/api/data"
	"SuperHack2023/server/api/decode"
	"SuperHack2023/server/api/encode"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	endpointMap := map[int]string{

		//OPTIMISM
		blckDecoder.OP_MAINNET: "https://rpc.ankr.com/optimism",
		blckDecoder.OP_GOERLI:  "https://goerli.optimism.io",

		//BASE
		blckDecoder.BASE_MAINNET: "https://base.blockpi.network/v1/rpc/public",
		blckDecoder.BASE_GOERLI:  "https://base-goerli.blockpi.network/v1/rpc/public",

		//MODE
		blckDecoder.MODE_TESTNET: "https://sepolia.mode.network",

		//ZORA
		blckDecoder.ZORA_MAINNET: "https://rpc.zora.energy",
		blckDecoder.ZORA_TESTNET: "https://testnet.rpc.zora.co/",

		//POLYGON
		blckDecoder.POLYGON_MUMBAI: "https://polygon-mumbai-infura.wallet.coinbase.com",
	}

	chainNameMap := map[string]int{
		//OPTIMISM
		"optimism_mainnet": blckDecoder.OP_MAINNET,
		"optimism_goerli":  blckDecoder.OP_GOERLI,

		//BASE
		"base_mainnet": blckDecoder.BASE_MAINNET,
		"base_goerli":  blckDecoder.BASE_GOERLI,

		//MODE
		"mode_testnet": blckDecoder.MODE_TESTNET,

		//ZORA
		"zora_mainnet": blckDecoder.ZORA_MAINNET,
		"zora_testnet": blckDecoder.ZORA_TESTNET,

		//POLYGON
		"polygon_mumbai": blckDecoder.POLYGON_MUMBAI,
	}

	decoder := blckDecoder.NewBLCKDecoder(endpointMap, chainNameMap)

	// Endpoints
	router.POST("/encode", encode.Handler())
	router.GET("/data/:fileName", data.Handler())
	router.POST("/decodePrivate", decode.PrivHandler(decoder))
	router.POST("/decodePublic/:address/:chainName", decode.PublicHandler(decoder))

	return router
}
