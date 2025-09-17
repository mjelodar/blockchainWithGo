package core

import (
	"crypto/sha256"

	"github.com/PRACTICING-GO/blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}
type BlockHasher struct{}

func (bh BlockHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderBytes())
	return types.Hash(h)
}
