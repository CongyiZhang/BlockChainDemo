package main

import "crypto/sha256"

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
