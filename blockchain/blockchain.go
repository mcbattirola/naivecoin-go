package blockchain

import (
	"fmt"
	"math"
	"strings"
)

const (
	// BlockGenerationInterval defines how often a block should be found, in SECONDS
	BlockGenerationInterval = 10

	// DifficultyAdjustmentInterval defines how often the difficulty should adjust to the increasing or decreasing network hashrate, in BLOCKS
	DifficultyAdjustmentInterval = 10
)

var genesisBlock = block{
	Index:      0,
	Hash:       "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
	Timestamp:  1465154705,
	Data:       "test genesis block",
	Nonce:      1,
	Difficulty: 1,
}

var blockchain = []block{genesisBlock}

// GetBlockchain returns current blockchain
func GetBlockchain() []block {
	return blockchain
}

func getAdjustedDifficulty(latestBlock block, blockchain []block) int {
	prevAdjustmentBlock := blockchain[len(GetBlockchain())-DifficultyAdjustmentInterval]

	timeExpected := BlockGenerationInterval * DifficultyAdjustmentInterval
	timeTaken := latestBlock.Timestamp - prevAdjustmentBlock.Timestamp

	if timeTaken < int64(timeExpected/2) {
		return prevAdjustmentBlock.Difficulty + 1
	} else if timeTaken > int64(timeExpected*2) {
		return prevAdjustmentBlock.Difficulty - 1
	} else {
		return prevAdjustmentBlock.Difficulty
	}

}

func getBinaryRepresentation(inputString string) string {
	binString := ""
	for _, char := range inputString {
		binString = fmt.Sprintf("%s%b", binString, char)
	}
	return binString
}

func getChainDifficulty(chain []block) float64 {
	difficulty := 0.0
	for _, block := range chain {
		difficulty += math.Pow(2.0, float64(block.Difficulty))
	}

	return difficulty
}

func getDifficulty(blockchain []block) int {
	latestBlock := getLatestBlock()
	if latestBlock.Index%DifficultyAdjustmentInterval == 0 && latestBlock.Index != 0 {
		return getAdjustedDifficulty(latestBlock, blockchain)
	} else {
		return latestBlock.Difficulty
	}
}

func getLatestBlock() block {
	return GetBlockchain()[len(blockchain)-1]
}

func findValidBlock(index int64, previousHash string, timestamp int64, data string, difficulty int) block {
	var nonce int32
	for {
		hash := calculateHash(index, previousHash, timestamp, data, difficulty, nonce)

		if hashMatchesDifficulty(hash, difficulty) {
			return block{
				Index:        index,
				Hash:         hash,
				PreviousHash: previousHash,
				Timestamp:    timestamp,
				Data:         data,
				Difficulty:   difficulty,
				Nonce:        nonce,
			}
		}
		nonce++
	}

}

func hashMatchesDifficulty(hash string, difficulty int) bool {
	hashInBinary := getBinaryRepresentation(hash)
	requiredPrefix := strings.Repeat("0", difficulty)

	return strings.HasPrefix(hashInBinary, requiredPrefix)
}

func isValidChainGenesisBlock(blockchainToValidate []block) bool {
	if len(blockchainToValidate) < 1 {
		return false
	}

	genesisToValidate := blockchainToValidate[0]

	return genesisToValidate.Index == 0 &&
		genesisToValidate.Hash == genesisBlock.Hash &&
		genesisToValidate.Timestamp == genesisBlock.Timestamp &&
		genesisToValidate.Data == genesisBlock.Data
}

func isValidChain(blockchainToValidate []block) bool {
	if !isValidChainGenesisBlock(blockchainToValidate) {
		return false
	}

	for i := 1; i < len(blockchainToValidate); i++ {
		if !isValidNewBlock(blockchainToValidate[i], blockchainToValidate[i-1]) {
			return false
		}
	}

	return true
}

func replaceChain(newBlocks []block) {
	if isValidChain(newBlocks) && getChainDifficulty(newBlocks) > getChainDifficulty(GetBlockchain()) {
		blockchain = newBlocks
		// TODO broadcastLatest()
	}
}
