// @Time : 2019/8/23 15:11 
// @Author : zwz
// @File : Blockchain
// @Software: GoLand
// @Desc: to do somewhat..

package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}
//NewBlockchain ...创建新的区块链
func NewBlockchain() *Blockchain{
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}
//SendData ...将数据写进区块链
func (bc *Blockchain)SendData(data string)  {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock,data)
	bc.AppendBlock(&newBlock)
}
//AppendBlock ...添加区块
func (bc *Blockchain)AppendBlock(newBlock *Block)  {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks,newBlock)
		return
	}
	if isVaild(*newBlock,*bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	}else{
		log.Fatal("无效的区块")
	}
}

func (bc *Blockchain)Print()  {
	for _,block :=range bc.Blocks  {
		fmt.Printf("Index: %d\n",block.Index)
		fmt.Printf("PreHash: %s\n",block.PrevBlockHash)
		fmt.Printf("CurrentHash: %s\n",block.Hash)
		fmt.Printf("Timestamp: %d\n",block.Timestamp)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Println()
	}
}
//判断是否合法
func isVaild(newBlock ,oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if caculateHash(newBlock) != newBlock.Hash {
	   return false
	}
	return true
}
