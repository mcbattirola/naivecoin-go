package transactions

import (
	"crypto/sha256"
	"fmt"
)

type TxOut struct {
	Adress string
	Amount int
}

type TxIn struct {
	TxOutID    string
	TxOutIndex int64
	Signature  string
}

type Transaction struct {
	ID     string
	TxIns  []TxIn
	TxOuts []TxOut
}

func getTransactionID(trans Transaction) string {
	var txInContent string

	for _, txIn := range trans.TxIns {
		txInContent += fmt.Sprintf("%s%d", txIn.TxOutID, txIn.TxOutIndex)
	}

	for _, txOut := range trans.TxOuts {
		txInContent += fmt.Sprintf("%s%d", txOut.Adress, txOut.Amount)
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(txInContent)))
}
