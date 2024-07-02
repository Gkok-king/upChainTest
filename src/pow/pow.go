package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	nickname := "Qi"
	nonce := 0
	// 计算4个零开头的哈希值
	fmt.Println("开始计算满足4个零开头的哈希值...")
	start := time.Now()
	hashValue1, content1, nonce1 := calculateHash(nickname, 4, nonce)
	// str := strings.Repeat("0", 6)
	// fmt.Printf("noce值2112: %s\n", str)
	duration1 := time.Since(start)
	fmt.Printf("noce值: %s\n", nonce1)
	fmt.Printf("花费时间: %s\n", duration1)
	fmt.Printf("Hash 内容: %s\n", content1)
	fmt.Printf("Hash 值: %s\n", hashValue1)
	nonce  = nonce1
	// 计算5个零开头的哈希值
	fmt.Println("开始计算满足5个零开头的哈希值...")
	start = time.Now()
	hashValue5, content5, nonce2 := calculateHash(nickname, 5, nonce)
	duration5 := time.Since(start)
	fmt.Printf("nonce值: %s\n", nonce2)
	fmt.Printf("花费时间: %s\n", duration5)
	fmt.Printf("Hash 内容: %s\n", content5)
	fmt.Printf("Hash 值: %s\n", hashValue5)
}

func calculateHash(nickname string, len int, nonce int) (string, string, int) {
	for {
		nonce++
		content := nickname + strconv.Itoa(nonce)	
		hash := sha256.Sum256([]byte(content))
		hashValue := hex.EncodeToString(hash[:])
		str := strings.Repeat("0", len)
		if hashValue[:len]== str {
			fmt.Printf("Hash 值: %s\n", str)
			return hashValue, content, nonce
		}
	}
}
