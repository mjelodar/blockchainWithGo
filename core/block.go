package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/PRACTICING-GO/blockchain/crypto"
	"github.com/PRACTICING-GO/blockchain/types"
)

type Header struct {
	Version       uint32 `json:"version"`
	Height        uint32 `json:"height"`
	TimeStamp     int64  `json:"timestamp"`
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

func (b *Block) Sign(privkey crypto.PrivateKey) error {
	sig, err := privkey.Sign(b.HeaderBytes())
	if err != nil {
		return err
	}
	b.Signature = sig
	b.Validator = privkey.PublicKey()
	return nil
}

func (b *Block) Verify() error {
	var zeroSig crypto.Signature
	if b.Signature == zeroSig {
		return nil // Genesis block
	}

	if !b.Signature.Verify(b.Validator, b.HeaderBytes()) {
		return fmt.Errorf("invalid block signature")
	}

	return nil
}

func (blk *Block) HeaderBytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(blk.Header); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
