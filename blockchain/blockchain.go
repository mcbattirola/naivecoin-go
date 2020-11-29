package blockchain

var genesisBlock = block{
	Index:     0,
	Hash:      "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
	Timestamp: 1465154705,
	Data:      "test genesis block",
}

var blockchain = []block{genesisBlock}

func GetBlockchain() []block {
	return blockchain
}

func getLatestBlock() block {
	return GetBlockchain()[len(blockchain)-1]
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
	if isValidChain(newBlocks) && len(newBlocks) > len(GetBlockchain()) {
		blockchain = newBlocks
		// broadcastLatest()
	}
}
