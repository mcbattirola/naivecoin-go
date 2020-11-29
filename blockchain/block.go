package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type block struct {
	Index        int64  `json:"index"`
	Hash         string `json:"hash"`
	PreviousHash string `json:"previous_hash"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
}

func calculateHash(index int64, previousHash string, timestamp int64, data string) string {
	blockString := fmt.Sprintf("%d%s%d%s", index, previousHash, timestamp, data)
	bytes := sha256.Sum256([]byte(blockString))
	return fmt.Sprintf("%x", bytes)
}

// GenerateNextBlock returns a block out of the input data,
// given the current state of the blockchain
func GenerateNextBlock(data string) block {
	previousBlock := getLatestBlock()

	nextIndex := previousBlock.Index + 1
	nextTimestamp := time.Now().UnixNano()
	nextHash := calculateHash(nextIndex, previousBlock.Hash, nextTimestamp, data)

	return block{
		Index:        nextIndex,
		Hash:         nextHash,
		PreviousHash: previousBlock.Hash,
		Timestamp:    nextTimestamp,
		Data:         data,
	}
}

func isValidNewBlock(newBlock block, previousBlock block) bool {
	if newBlock.Index != previousBlock.Index+1 {
		return false
	}
	if newBlock.PreviousHash != previousBlock.Hash {
		return false
	}
	if newBlock.Hash != calculateHash(newBlock.Index, newBlock.PreviousHash, newBlock.Timestamp, newBlock.Data) {
		return false
	}

	return true
}
