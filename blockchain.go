package main

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

func (bc *BlockChain) AddBlock(data string)  {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}