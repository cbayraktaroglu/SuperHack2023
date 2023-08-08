package main

import (
	blckEncoder "SuperHack2023/encoder"
	"log"
)

func main() {
	encoder, _ := blckEncoder.NewBLCKEncoder(100)
	err := encoder.ConvertFile("./target/index.html", "./output")
	log.Print(err)
}
