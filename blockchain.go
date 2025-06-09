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

type Blockchain struct {
	transactionPool []string // 交易池
	chain           []*Block // 區塊鏈
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("Chain: %d \n", i)
		block.Print()
	}
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := newBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// init 函數用於初始化
func init() {
	log.SetPrefix("INFO: ")
}

// main 函數是程式的入口點
func main() {
	log.Println("Starting blockchain...")

	blockChain := NewBlockchain()
	blockChain.Print()
}
