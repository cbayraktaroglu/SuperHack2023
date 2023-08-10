package encode

import (
	blckEncoder "SuperHack2023/encoder"
	"SuperHack2023/server/responses"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			responses.FailResponse(c, http.StatusBadRequest)
			return
		}

		// Get the block size
		blockSizeString := form.Value["blockSize"]
		blockSize, err := strconv.Atoi(blockSizeString[0])

		if err != nil {
			log.Println("Cannot convert the block size: ", err)
			responses.FailResponse(c, http.StatusInternalServerError)
		}

		// Get the files from the form
		files := form.File["file"]

		// Uploaded file name
		targetFileName := ""

		// Receive the file
		for _, file := range files {
			// Create a unique file name
			targetFileName = filepath.Join("./uploads/main", file.Filename)

			// Create the file
			outFile, err := os.Create(targetFileName)
			if err != nil {
				log.Println("A: ", err)
				responses.FailResponse(c, http.StatusInternalServerError)
				return
			}
			defer outFile.Close()

			// Open the uploaded file
			uploadedFile, err := file.Open()
			if err != nil {
				log.Println("B: ", err)
				responses.FailResponse(c, http.StatusInternalServerError)
				return
			}
			defer uploadedFile.Close()

			// Copy the uploaded file's contents to the new file
			_, err = io.Copy(outFile, uploadedFile)
			if err != nil {
				log.Println("C: ", err)
				responses.FailResponse(c, http.StatusInternalServerError)
				return
			}
		}

		// Convert the file into blobs
		encoder, err := blckEncoder.NewBLCKEncoder(blockSize)

		if err != nil {
			log.Println("Cannot initialize the block encoder: ", err)
			responses.FailResponse(c, http.StatusInternalServerError)
		}

		checkSum, err := encoder.GetFileCheckSumSHA256(targetFileName)

		if err != nil {
			log.Println("Error while calculating the checksum: ", err)
			responses.FailResponse(c, http.StatusInternalServerError)
		}

		data, err := encoder.ConvertFileWOutput(targetFileName, "./uploads/processed")

		if err != nil {
			log.Println("Error while processing the file: ", err)
			responses.FailResponse(c, http.StatusInternalServerError)
		}

		responses.SuccessResponse(c, map[string]interface{}{"Checksum": checkSum, "FileList": data})
		return
	}
}
