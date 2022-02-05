package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 4. 引入区块链
type BlockChain struct {
	//blocks []*Block
	db     *bolt.DB
	tail   []byte
}

// 5. 定义一个区块链
const blockChainDb = "blockchain.db"
const blockBucket = "blockBucket"

func NewBlockChain() *BlockChain {
	// 创建一个创世块

	var lastHash []byte
	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket失败")
			}
			genesis := GenesisBlock()
			bucket.Put(genesis.Hash, genesis.toByte())
			bucket.Put([]byte("lastHashKey"), genesis.Hash)
			lastHash = genesis.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}

func GenesisBlock() *Block {
	return NewBlock("Go创世块", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	//lastBlock := bc.blocks[len(bc.blocks)-1]
	//prevHash := lastBlock.Hash
	//block := NewBlock(data, prevHash)
	//bc.blocks = append(bc.blocks, block)
}
