package core

import (
	"testing"
	"time"

	"github.com/PRACTICING-GO/blockchain/crypto"
	"github.com/PRACTICING-GO/blockchain/types"
	"github.com/stretchr/testify/assert"
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

func TestSignBlock(t *testing.T) {
	block := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, block.Sign(privKey))
	assert.NotNil(t, block.Signature)
}

func TestVerifyBlock(t *testing.T) {
	block := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, block.Sign(privKey))
	assert.Nil(t, block.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	block.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, block.Verify())
	// Tamper with the block to make the signature invalid
	block.Header.Height = 1
	assert.NotNil(t, block.Verify())
}
