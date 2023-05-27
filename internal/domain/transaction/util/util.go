package util

import "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"

// ArrayTxMemoryToArrayValue converts an array of transactions from memory to an array of transactions values.
func ArrayTxMemoryToArrayValue(txs []*entity.Transaction) (txsValues []entity.Transaction) {
	for _, tx := range txs {
		txsValues = append(txsValues, *tx)
	}
	return
}
