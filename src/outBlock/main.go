// package outblock

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"strconv"
// 	"strings"
// )

// type Block struct {
//     Index     int
//     Timestamp string
//     Data      string
//     PrevHash  string
//     Hash      string
//     Nonce     int
// }


// func calculateHash(block Block) string {
//     record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
//     h := sha256.New()
//     h.Write([]byte(record))
//     hashed := h.Sum(nil)
//     return hex.EncodeToString(hashed)
// }

// func mineBlock(difficulty int, block Block) Block {
//     for !isHashValid(block.Hash, difficulty) {
//         block.Nonce++
//         block.Hash = calculateHash(block)
//     }
//     return block
// }

// func isHashValid(hash string, difficulty int) bool {
//     prefix := strings.Repeat("0", difficulty)
//     return strings.HasPrefix(hash, prefix)
// }
