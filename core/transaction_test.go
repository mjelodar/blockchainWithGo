package core

import (
	"testing"

	"github.com/PRACTICING-GO/blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	tx := &Transaction{
		data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	tx := &Transaction{
		data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))

	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
