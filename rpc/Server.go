// @Time : 2019/8/23 15:42 
// @Author : zwz
// @File : Server
// @Software: GoLand
// @Desc: to do somewhat..

package main

import (
	"demoChain/core"
	"encoding/json"
	"io"
	"net/http"
)
//blockchain ...在main 函数中初始化
var blockchain *core.Blockchain
//run ...初始化http服务
func run()  {
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriteHandler)
	http.ListenAndServe("localhost:8088",nil)
}
func blockchainGetHandler(w http.ResponseWriter,r *http.Request)  {
	bytes,error :=  json.Marshal(blockchain)
	if error !=nil  {
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter,r *http.Request)  {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w,r)
}

func main()  {
	blockchain = core.NewBlockchain()
	run()
}