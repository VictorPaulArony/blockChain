package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Id        int
	TimeStamp string
	Data      string
	PrevHash  string
	Hash      string
}
type Blockchain struct {
	block []Block
}

var block = Block{}

func Calc(block Block) string {
	res := strconv.Itoa(block.Id) + block.TimeStamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(res))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreatGenesis() Block {
	return Block{0, time.Now().String(), "Genesis Block", "", ""}
}

func CreatBlock(prevBlock Block, data string) Block {
	block.Id = prevBlock.Id + 1
	block.TimeStamp = time.Now().String()
	block.Data = data
	block.PrevHash = prevBlock.Hash
	block.Hash = Calc(block)
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.block[len(bc.block)-1]
	newBlock := CreatBlock(prevBlock, data)
	bc.block = append(bc.block, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.block); i++ {
		prevBlock := bc.block[i-1]
		currentBlock := bc.block[i]
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
		if currentBlock.Hash != Calc(currentBlock) {
			return false
		}
	}
	return true
}

func main() {
	genesisBlock := CreatGenesis()
	genesisBlock.Hash = Calc(genesisBlock)
	blockchain := Blockchain{[]Block{genesisBlock}}

	blockchain.AddBlock("second block")
	blockchain.AddBlock("next block")
	blockchain.AddBlock("next block")

	for _, block := range blockchain.block {
		fmt.Printf("index: %d\n", block.Id)
		fmt.Printf("TIMESTAMP: %s\n", block.TimeStamp)
		fmt.Printf("DATA: %s\n", block.Data)
		fmt.Printf("PREVHASH: %s\n", block.PrevHash)
		fmt.Printf("HASH: %s\n", block.Hash)
	}
}
