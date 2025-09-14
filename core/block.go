package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"github.com/PRACTICING-GO/blockchain/crypto"
	"github.com/PRACTICING-GO/blockchain/types"
	"io"
)

type Header struct {
	Version       uint32 `json:"version"`
	Height        uint32 `json:"height"`
	TimeStamp     uint64 `json:"timestamp"`
	PrevBlockHash types.Hash
	MerkleRoot    types.Hash
	Nonce         uint64 `json:"nonce"`
}

type Block struct {
	Header       *Header       `json:"header"`
	Transactions []Transaction `json:"transactions"`
	Validator    crypto.PublicKey
	Signature    crypto.Signature
	hash         types.Hash
}

func (b *Block) Hash(hashAlgorithm Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hashAlgorithm.Hash(b)
	}
	return b.hash
}

func (b *Block) Encode(writer io.Writer, enc Encoder[Block]) error {
	return enc.Encode(writer, *b)
}

func (b *Block) Decode(reader io.Reader, dec Decoder[Block]) error {
	return dec.Decode(reader, *b)
}
