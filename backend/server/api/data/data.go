package data

import (
	"SuperHack2023/server/responses"
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileName := c.Param("fileName")

		// Sanity Check
		if !strings.Contains(fileName, ".blck") {
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		// Open the target file
		targetFile, err := os.Open("./uploads/processed/" + fileName)

		if err != nil {
			responses.FailResponse(c, http.StatusInternalServerError)
			return
		}

		defer targetFile.Close()

		// Initialize the buffered reader
		reader := bufio.NewReader(targetFile)

		buffer := []byte{}

		for {
			byteValue, readErr := reader.ReadByte()

			// Write it to the buffer
			buffer = append(buffer, byteValue)

			if readErr != nil {
				break
			}

		}

		responses.SuccessResponse(c, buffer)
	}
}
