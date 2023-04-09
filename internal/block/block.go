package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	var record string
	record = fmt.Sprint(block.Index) + block.Timestamp + fmt.Sprint(block.BPM) + block.PrevHash

	hash := sha256.New()
	hash.Write([]byte(record))

	hashsum := hash.Sum(nil)
	return hex.EncodeToString(hashsum)
}
