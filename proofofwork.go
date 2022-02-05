package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	targetStr := "f000000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	block := pow.block
	hash := [32]byte{}
	// 拼装数据
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PreHash,
			block.MerkleRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(tmp, []byte{})
		// 做哈希运算
		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！%x\n", hash)
			break
		} else {
			nonce++
			fmt.Println(nonce)
		}
	}

	return hash[:], nonce
	//return []byte("helloworld"), 10

}
