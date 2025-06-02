package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce         int      // (隨機數) 根據規則計算的答案
	previousHash  string   // 前一個區塊的雜湊值
	timestamp     int64    // 區塊的時間戳
	stransactions []string // 區塊中的交易列表
}

// Print 方法用於輸出區塊的詳細信息
func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("stransactions: %s\n", b.stransactions)
}

func newBlock(nonce int, previousHash string) *Block {
	b := &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
	return b
}

// init 函數用於初始化
func init() {
	log.SetPrefix("INFO: ")
}

// main 函數是程式的入口點
func main() {
	log.Println("Starting blockchain...")

	b := newBlock(777, "init hash")
	b.Print()
}
