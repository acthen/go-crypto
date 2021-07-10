package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)


type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

// Only accessible in the blockchain package.
var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks - 1].hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", ""}
	newBlock.calculateHash()
	return &newBlock
}

// a function whose name starts with capital 
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock("Genesis Block"))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}