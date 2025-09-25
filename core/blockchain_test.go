package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewBlockchainWithGenesisBlock() *Blockchain {
	return NewBlockchain(randomBlock(0))
}

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchainWithGenesisBlock()
	assert.NotNil(t, bc)
	assert.Equal(t, uint32(0), bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := NewBlockchainWithGenesisBlock()
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockchainWithGenesisBlock()
	block1 := randomBlock(1)

	assert.Nil(t, bc.AddBlock(block1))
	assert.Equal(t, uint32(1), bc.Height())
	assert.True(t, bc.HasBlock(1))
}
