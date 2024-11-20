package BLC

import (
	"bytes"
	"crypto/sha256"
	"nft/server/utils"
	"strconv"
	"time"
)

type Block struct {
	Height        int64  // 区块的高度
	Timestamp     int64  // 时间戳，创建区块时的时间
	PrevBlockHash []byte // 父区块的Hash
	Data          []byte // 交易数据
	Hash          []byte // 当前区块的Hash
}

func (block *Block) SetHash() {
	// Height 转为 [] byte
	heightBytes := utils.IntToHex(block.Height)
	// Timestamp 转为 [] byte
	timeString := strconv.FormatInt(block.Timestamp, 2) // 2是二进制
	timeBytes := []byte(timeString)
	// 拼接所有属性
	headers := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	// 生成 Hash
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Hash:          nil,
	}
	block.SetHash()
	return block
}

// CreateGenesisBlock 生成创世区块
func CreateGenesisBlock(data string) *Block {
	block := NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return block
}
