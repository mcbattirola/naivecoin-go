package blockchain

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	now := time.Now().UnixNano()

	hash := calculateHash(0, "", now, "test genesis block")

	if hash == "" {
		t.Errorf("Hash should not be empty")
	}

	newDataHash := calculateHash(0, "", now, "test genesis block - with a new data")
	if hash == newDataHash {
		t.Errorf("Expected different hashes with different data, got same value: %s", newDataHash)
	}
}

func TestGenerateNextBlock(t *testing.T) {
	latestBlock := getLatestBlock()

	newBlockData := "my new block"
	newBlock := GenerateNextBlock(newBlockData)

	if newBlock.Index != latestBlock.Index+1 {
		t.Errorf("New block's index must be the latest block's index plus one, but found latest index %d and new index %d.", newBlock.Index, latestBlock.Index)
	}

	if newBlock.PreviousHash != latestBlock.Hash {
		t.Errorf("New block's previous hash doesnt match lates block hash.")
	}

	if newBlock.Data != newBlockData {
		t.Errorf("New block's data must match input. Expected %s, got %s", newBlockData, newBlock.Data)
	}

}
