package types

import (
	"crypto/rand"
	"encoding/hex"
)

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic("HashFromBytes: input byte slice must be exactly 32 bytes long")
	}

	var value [32]uint8

	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func (h *Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h *Hash) ToSlice() []byte {
	slice := make([]byte, 32)
	for i := 0; i < 32; i++ {
		slice[i] = h[i]
	}
	return slice
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

func RandomByte(size int) []byte {
	if size <= 0 {
		panic("RandomByte: size must be greater than 0")
	}

	b := make([]byte, size)
	rand.Read(b)
	return b
}

func RandomHash() Hash {
	return HashFromBytes(RandomByte(32))
}
