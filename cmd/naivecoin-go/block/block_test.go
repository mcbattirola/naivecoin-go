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
	newBlock := generateNextBlock(newBlockData)

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

func TestIsValidChain(t *testing.T) {
	emptyChain := make([]block, 0)
	if isValidChain(emptyChain) {
		t.Errorf("Expected a empty chain to be invalid")
	}

	chainWithInvalidGenesisBlock := make([]block, 1)
	chainWithInvalidGenesisBlock[0] = block{Index: 1}
	if isValidChain(chainWithInvalidGenesisBlock) {
		t.Errorf("Expected a chain with genesis block with index = 1 to be invalid.")
	}

	chainWithInvalidGenesisBlock[0] = block{
		Index: 0,
		Hash:  "",
	}
	if isValidChain(chainWithInvalidGenesisBlock) {
		t.Errorf("Expected a chain with genesis block with hash different from '%s' to be invalid.", genesisBlock.Hash)
	}

	chainWithInvalidGenesisBlock[0] = block{
		Index:     0,
		Hash:      genesisBlock.Hash,
		Timestamp: 0,
	}
	if isValidChain(chainWithInvalidGenesisBlock) {
		t.Errorf("Expected a chain with genesis block with hash different from '%d' to be invalid.", genesisBlock.Timestamp)
	}

	chainWithInvalidGenesisBlock[0] = block{
		Index:     0,
		Hash:      genesisBlock.Hash,
		Timestamp: genesisBlock.Timestamp,
		Data:      "not the genesis block Data",
	}
	if isValidChain(chainWithInvalidGenesisBlock) {
		t.Errorf("Expected a chain with genesis block with hash different from '%s' to be invalid.", genesisBlock.Data)
	}

	if !isValidChain(blockchain) {
		t.Errorf("Expected base blockchain to be valid.")
	}

	chainWithUnorderedIndexes := make([]block, 3)
	chainWithUnorderedIndexes[0] = genesisBlock
	chainWithUnorderedIndexes[1] = generateNextBlock("second block")

	newBlockTimestamp := time.Now().UnixNano()
	newBlockData := "new data"

	chainWithUnorderedIndexes[2] = block{
		Index:        0,
		Hash:         calculateHash(0, chainWithUnorderedIndexes[1].Hash, newBlockTimestamp, newBlockData),
		PreviousHash: chainWithUnorderedIndexes[1].Hash,
		Timestamp:    newBlockTimestamp,
		Data:         newBlockData,
	}

	if isValidChain(chainWithUnorderedIndexes) {
		t.Errorf("Expected a chain without ordered indexes to be invalid.")
	}

	chainWithIncorrectHashes := make([]block, 2)
	chainWithIncorrectHashes[0] = genesisBlock
	chainWithIncorrectHashes[1] = block{
		Index:        1,
		Hash:         calculateHash(0, chainWithIncorrectHashes[0].Hash, newBlockTimestamp, newBlockData),
		PreviousHash: chainWithIncorrectHashes[0].Hash + "something",
		Timestamp:    newBlockTimestamp,
		Data:         newBlockData,
	}

	if isValidChain(chainWithIncorrectHashes) {
		t.Errorf("Expected a chain to have each block's previous hash to aways match the preivou's hash.")
	}
}
