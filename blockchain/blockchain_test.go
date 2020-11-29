package blockchain

import (
	"testing"
	"time"
)

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
	chainWithUnorderedIndexes[1] = GenerateNextBlock("second block")

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

func TestHashMatchesDificulty(t *testing.T) {

	if !hashMatchesDificulty(string([]byte{0}), 1) {
		t.Errorf("Expected string '0' to match dificulty 1")
	}

	if hashMatchesDificulty(string([]byte{0, 65, 66, 67, 0, 0}), 2) {
		t.Errorf("Expected string '0' NOT to match dificulty 1")
	}

	if hashMatchesDificulty(string([]byte{0, 0, 0, 65, 65}), 5) {
		t.Errorf("Expected string prefixed with '000' NOT to match dificulty 5")
	}

	if !hashMatchesDificulty(string([]byte{0, 0, 0}), 2) {
		t.Errorf("Expected string prefixed with '000' to match dificulty 2")
	}

	if !hashMatchesDificulty(string([]byte{0, 0, 0}), 3) {
		t.Errorf("Expected string prefixed with '000' to match dificulty 3")
	}

}
