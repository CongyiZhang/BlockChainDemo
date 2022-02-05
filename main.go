package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("新块")
	for i, block := range bc.blocks {
		//block := NewBlock("创世币", []byte{})
		fmt.Printf("====== 当前区块高度 %d======\n", i)
		fmt.Printf("前区块哈希：%x\n\n", block.PreHash)
		fmt.Printf("当前区块哈希：%x\n\n", block.Hash)
		fmt.Printf("区块数据：%s\n\n", block.Data)
	}

}
