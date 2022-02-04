package main

import "fmt"

type Block struct {
	PreHash []byte
	Hash    []byte
	Data    []byte
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}
	return &block
}
func main() {
	block := NewBlock("创世币", []byte{})
	fmt.Printf("前区块哈希：%x\n\n", block.PreHash)
	fmt.Printf("当前区块哈希：%x\n\n", block.Hash)
	fmt.Printf("区块数据：%s\n\n", block.Data)
	fmt.Println("hello world")
}
