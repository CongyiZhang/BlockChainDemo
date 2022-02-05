package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	Version    uint64
	PreHash    []byte
	MerkleRoot []byte
	TimeStamp  uint64
	Difficulty uint64
	Nonce      uint64
	Hash       []byte
	Data       []byte
}

// 实现一个辅助函数 uint64 -> []byte

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return []byte{}
}

// 2. 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PreHash:    preBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func (block *Block) toByte() []byte {
	return []byte{}
}

// 3. 生成哈希

func (block *Block) SetHash() {
	//var blockInfo []byte
	//blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	//blockInfo = append(blockInfo, block.PreHash...)
	//blockInfo = append(blockInfo, block.MerkleRoot...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	//blockInfo = append(blockInfo, block.Data...)

	temp := [][]byte{
		Uint64ToByte(block.Version),
		block.PreHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	blockInfo := bytes.Join(temp, []byte{})

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
