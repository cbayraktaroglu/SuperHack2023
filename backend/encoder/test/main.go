package main

import (
	blckEncoder "SuperHack2023/encoder"
	"log"
)

func main() {
	encoder, _ := blckEncoder.NewBLCKEncoder(62750)
	checksum, _ := encoder.GetFileCheckSumSHA256("./target/optum.gif")
	log.Println("File checksum:", checksum)

	checksum, _ = encoder.GetFileCheckSumSHA256("./target/uploaded.gif")
	log.Println("File checksum:", checksum)

	//err := encoder.ConvertFile("./target/optum.gif", "./output")
	//log.Print(err)
}
