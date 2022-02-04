package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreHash []byte
	Hash    []byte
	Data    []byte
}

// 2. 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}
	block.SetHash()
	return &block
}

// 3. 生成哈希

func (block *Block) SetHash() {
	blockInfo := append(block.PreHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// 4. 引入区块链
type BlockChain struct {
	blocks []*Block
}

// 5. 定义一个区块链

func NewBlockChain() *BlockChain {
	// 创建一个创世块
	genesis := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesis},
	}
}

func GenesisBlock() *Block {
	return NewBlock("Go创世块", []byte{})
}
func main() {
	bc := NewBlockChain()
	for i, block := range bc.blocks {
		//block := NewBlock("创世币", []byte{})
		fmt.Printf("====== 当前区块高度 %d======\n", i)
		fmt.Printf("前区块哈希：%x\n\n", block.PreHash)
		fmt.Printf("当前区块哈希：%x\n\n", block.Hash)
		fmt.Printf("区块数据：%s\n\n", block.Data)
		fmt.Println("hello world")
	}

}
