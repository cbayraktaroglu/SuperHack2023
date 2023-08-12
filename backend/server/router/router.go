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
		blckDecoder.OP_MAINNET:     "https://polygon-mumbai-infura.wallet.coinbase.com",
		blckDecoder.OP_GOERLI:      "https://polygon-mumbai-infura.wallet.coinbase.com",
		blckDecoder.POLYGON_MUMBAI: "https://polygon-mumbai-infura.wallet.coinbase.com",
	}

	chainNameMap := map[string]int{
		"optimism": blckDecoder.OP_MAINNET,
		"goerli":   blckDecoder.OP_GOERLI,
		"polygon":  blckDecoder.POLYGON_MUMBAI,
	}

	decoder := blckDecoder.NewBLCKDecoder(endpointMap, chainNameMap)

	// Endpoints
	router.POST("/encode", encode.Handler())
	router.GET("/data/:fileName", data.Handler())
	router.POST("/decodePrivate", decode.PrivHandler(decoder))
	router.POST("/decodePublic/:address/:chainName", decode.PublicHandler(decoder))

	return router
}
