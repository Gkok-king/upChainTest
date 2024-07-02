package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//第一题
	// func1()

	//第二题
	func2()
}

func func1()(string){
	nickname := "Qi"
	nonce := 0
	// 计算4个零开头的哈希值
	fmt.Println("开始计算满足4个零开头的哈希值...")
	start := time.Now()
	hashValue1, content1, nonce1 := calculateHash(nickname, 4, nonce)
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

	return hashValue1;
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

// 实践⾮对称加密 RSA 
// ❖ 先⽣成⼀个公私钥对
// ❖ ⽤私钥对符合POW⼀个昵称 + nonce 进⾏私钥签名
// ❖ ⽤公钥验证
func func2(){

	// 得到符合的POW
	var value = func1()

	//生成秘钥对
	fmt.Printf("符合POW⼀个昵称 + nonce : %s\n", value)
    privateKey, publicKey, err := creatKey(2048)
	if err != nil {
        return 
    }

    // 签名消息
    signature, err := sign(value, privateKey)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("签名值: %x\n", signature)

    // 验证签名
    err = check([]byte(value), signature, publicKey)
    if err != nil {
        fmt.Println("Verification failed:", err)
    } else {
        fmt.Println("Verification successful")
    }
}


func  creatKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, nil, err
    }
    return privateKey, &privateKey.PublicKey, nil
}
func  sign(value string, privKey *rsa.PrivateKey) ([]byte, error) {
    hashed := sha256.Sum256([]byte(value))
    return rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])

}
func  check(value []byte, signature []byte, pubKey *rsa.PublicKey) error {
    hashed := sha256.Sum256(value)
    return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signature)
}

