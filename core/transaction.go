package core

import (
	"fmt"

	"github.com/PRACTICING-GO/blockchain/crypto"
)

type Transaction struct {
	data []byte // Placeholder for transaction data

	Validator crypto.PublicKey
	Signature crypto.Signature
}

func (tx *Transaction) Sign(privkey crypto.PrivateKey) error {
	sig, err := privkey.Sign(tx.data)
	if err != nil {
		return err
	}
	tx.Signature = sig
	tx.Validator = privkey.PublicKey()
	return nil
}

func (tx *Transaction) Verify() error {
	var zeroSig crypto.Signature
	if tx.Signature == zeroSig {
		return fmt.Errorf("missing signature")
	}

	if !tx.Signature.Verify(tx.Validator, tx.data) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
