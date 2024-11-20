package BLC

type BlockChain struct {
	Blocks []*Block
}

// CreateBlockChain 创建带有创世区块的区块链
func CreateBlockChain() *BlockChain {
	// 创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.......")
	return &BlockChain{[]*Block{genesisBlock}}
}

// 增加区块到区块链里面
func (block *BlockChain) AddBlockToBlockChain() {
	
}
