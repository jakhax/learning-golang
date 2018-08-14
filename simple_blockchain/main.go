package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// block - items in a blockchain
type Block struct {
	Index     int    //position of data record on blockchain
	Timestamp string // auto time of when data is added to the blockchain
	BPM       int    // beats per minute - heartrate
	Hash      string //SHA256 identifier of this data record
	PrevHash  string //SHA256 identifier of the previous data on the block chain
}
type Blockchain []Block

func calculateHash(b Block) {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.BPM) + b.PrevHash
	h := sha256.New()
	h.write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

//generate new Block function
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.string()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock, nil
}

func main() {

}
