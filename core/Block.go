// @Time : 2019/8/23 14:56 
// @Author : zwz
// @File : Block
// @Software: GoLand
// @Desc: to do somewhat..

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64 //区块编号
	Timestamp int64 //区块时间戳
	PrevBlockHash  string//上一个区块的哈希值
	Hash string //当前区块的哈希值
	Data string //区块数据
}
//caculateHash ...哈希算法
func caculateHash(b Block)string  {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash+b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}
//GenerateNewBlock ...创建新的块
func GenerateNewBlock(preBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index +1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = caculateHash(newBlock)
	return  newBlock
}
//GenerateGenesisBlock ...创建原始的区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	GenerateNewBlock(preBlock,"Genesis  Block")
	return preBlock
}




