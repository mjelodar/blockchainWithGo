package core

import (
	"github.com/PRACTICING-GO/blockchain/crypto"
	"github.com/PRACTICING-GO/blockchain/types"

	"io"
)

type Transaction struct {
	data []byte // Placeholder for transaction data

	Validator crypto.PublicKey
	Signature crypto.Signature
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	// Implement encoding logic here
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	// Implement decoding logic here
	return nil
}
