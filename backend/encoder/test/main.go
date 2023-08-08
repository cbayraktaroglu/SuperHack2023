package main

import (
	blckEncoder "SuperHack2023/encoder"
	"log"
)

func main() {
	encoder, _ := blckEncoder.NewBLCKEncoder(10000)
	err := encoder.ConvertFile("./target/optum.gif", "./output")
	log.Print(err)
}
