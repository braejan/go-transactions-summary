package entity_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewTxFile tests the NewTxFile function.
func TestNewTxFile(t *testing.T) {
	// When calling NewTxFile with a valid file path
	// Then it should return a new TxFile instance.
	file := entity.NewTxFile("txns.csv", "/path/to/txns.csv", "hash", 10)
	assert.Equal(t, "txns.csv", file.Name)
	assert.Equal(t, "/path/to/txns.csv", file.Path)
	assert.Equal(t, "hash", file.Hash)
	assert.Equal(t, int64(10), file.Lines)
}
