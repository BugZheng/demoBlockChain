// @Time : 2019/8/23 15:26 
// @Author : zwz
// @File : main
// @Software: GoLand
// @Desc: to do somewhat..

package main2

import "demoChain/core"

func main2()  {
	bc := core.NewBlockchain()
	bc.SendData("Send 666 to MyDataList")
	bc.SendData("Send 777 to MyDataList")
	bc.Print()
}