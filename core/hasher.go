package core

import "github.com/PRACTICING-GO/blockchain/types"

type Hasher[T any] interface {
	Hash(T) types.Hash
}
