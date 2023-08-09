package blckEncoder

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strconv"
)

type BLCKEncoder struct {
	encodeBuffer      []byte
	currentBufferSize int
	partNumber        int
	sizeLimit         int
	outputFile        *os.File
}

func NewBLCKEncoder(limit int) (encoder *BLCKEncoder, err error) {
	// Sanity check
	if limit <= 0 {
		err = errors.New("invalid size limit")
		return
	}

	encoder = &BLCKEncoder{
		encodeBuffer:      []byte{},
		currentBufferSize: 0,
		partNumber:        0,
		sizeLimit:         limit,
	}

	return
}

func (blckEncd *BLCKEncoder) GetEncodeBuffer() (buff []byte) {
	buff = blckEncd.encodeBuffer
	return
}

func (blckEncd *BLCKEncoder) GetCurrentBufferSize() (size int) {
	size = blckEncd.currentBufferSize
	return
}

func (blckEncd *BLCKEncoder) GetCurrentPartNumber() (number int) {
	number = blckEncd.partNumber
	return
}

func (blckEncd *BLCKEncoder) GetCurrentPartNumberString() (number string) {
	number = strconv.Itoa(blckEncd.partNumber)
	return
}

func (blckEncd *BLCKEncoder) GetSizeLimit() (limit int) {
	limit = blckEncd.sizeLimit
	return
}

func (blckEncd *BLCKEncoder) Reset() {
	blckEncd.currentBufferSize = 0
	blckEncd.encodeBuffer = []byte{}
	blckEncd.partNumber = 0
	return
}

func (blckEncd *BLCKEncoder) ResetBuffer() {
	blckEncd.currentBufferSize = 0
	blckEncd.encodeBuffer = []byte{}
	return
}

func (blckEncd *BLCKEncoder) IncrementPartNumber() {
	blckEncd.partNumber += 1
	return
}

func (blckEncd *BLCKEncoder) GetFileCheckSumSHA256(filePath string) (checksum string, err error) {
	targetFile, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer targetFile.Close()

	hash := sha256.New()

	_, err = io.Copy(hash, targetFile)

	if err != nil {
		return
	}

	checksum = hex.EncodeToString(hash.Sum(nil))

	return
}

func (blckEncd *BLCKEncoder) ConvertFile(inputPath string, outputPath string) (err error) {
	// First reset the stats
	blckEncd.Reset()

	// Open the target file
	targetFile, err := os.Open(inputPath)

	if err != nil {
		return
	}

	defer targetFile.Close()

	// Initialize the buffered reader
	reader := bufio.NewReader(targetFile)

	//Reading Loop
	for {
		byteValue, readErr := reader.ReadByte()

		// Write it to the buffer
		blckEncd.encodeBuffer = append(blckEncd.encodeBuffer, byteValue)
		// Since the above function only reads 1 byte increse the counter by one
		blckEncd.currentBufferSize += 1

		if blckEncd.currentBufferSize >= blckEncd.sizeLimit || readErr != nil {

			blckEncd.outputFile, err = os.Create(outputPath + "/" + blckEncd.GetCurrentPartNumberString() + ".blck")

			if err != nil {
				return
			}

			writer := bufio.NewWriter(blckEncd.outputFile)

			_, err = writer.Write(blckEncd.encodeBuffer)

			if err != nil {
				return
			}

			writer.Flush()

			blckEncd.ResetBuffer()
			blckEncd.IncrementPartNumber()

			if readErr != nil {
				break
			}
		}
	}

	return
}
