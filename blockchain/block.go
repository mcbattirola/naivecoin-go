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
	Difficulty   int    `json:"difficulty"`
	Nonce        int32  `json:"nonce"`
}

func calculateHash(index int64, previousHash string, timestamp int64, data string, difficulty int, nonce int32) string {
	blockString := fmt.Sprintf("%d%s%d%s%d%d", index, previousHash, timestamp, data, difficulty, nonce)
	bytes := sha256.Sum256([]byte(blockString))
	return fmt.Sprintf("%x", bytes)
}

// GenerateNextBlock returns a block out of the input data,
// given the current state of the blockchain
func GenerateNextBlock(data string, nonce int32) block {
	previousBlock := getLatestBlock()

	nextIndex := previousBlock.Index + 1
	nextTimestamp := time.Now().UnixNano()
	nextHash := calculateHash(nextIndex, previousBlock.Hash, nextTimestamp, data, 0, nonce)

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
	if newBlock.Hash != calculateHash(newBlock.Index, newBlock.PreviousHash, newBlock.Timestamp, newBlock.Data, 0, 0) {
		return false
	}
	if !isValidTimestamp(newBlock, previousBlock) {
		return false
	}

	return true
}

func isValidTimestamp(newBlock block, previousBlock block) bool {
	return (previousBlock.Timestamp-60 < newBlock.Timestamp) && (newBlock.Timestamp-60 < time.Now().UnixNano())
}
