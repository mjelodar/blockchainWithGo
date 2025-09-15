package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/PRACTICING-GO/blockchain/types"
)

func randomBlock(height uint32) *Block {
	Header := &Header{
		Version:       1,
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
		PrevBlockHash: types.RandomHash(),
		MerkleRoot:    types.RandomHash(),
		Nonce:         0,
	}
	tx := Transaction{
		data: []byte("foo"),
	}
	return &Block{
		Header:       Header,
		Transactions: []Transaction{tx},
	}
}

func TestHashBlock(t *testing.T) {
	block := randomBlock(0)
	fmt.Println(block.Hash(BlockHasher{}))
}
