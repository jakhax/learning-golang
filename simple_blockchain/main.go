package main

// block - items in a blockchain
type Block struct {
	Index     int    //position of data record on blockchain
	Timestamp string // auto time of when data is added to the blockchain
	BPM       int    // beats per minute - heartrate
	Hash      string //SHA256 identifier of this data record
	Prevhash  string //SHA256 identifier of the previous data on the block chain
}
type Blockchain []Block
