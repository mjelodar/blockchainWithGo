package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/PRACTICING-GO/blockchain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (pk *PrivateKey) Sign(data []byte) (Signature, error) {
	// Implementation for signing data with the private key
	// hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, pk.key, data)
	if err != nil {
		return Signature{}, err
	}
	return Signature{r: r, s: s}, nil
}

func GeneratePrivateKey() PrivateKey {
	// Implementation for generating a private key
	key, error := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if error != nil {
		panic(error)
	}
	return PrivateKey{key: key}
}

func (pk *PrivateKey) PublicKey() PublicKey {
	// Implementation for getting the public key from the private key
	return PublicKey{key: &pk.key.PublicKey}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pk *PublicKey) toSlice() []byte {
	return elliptic.MarshalCompressed(pk.key, pk.key.X, pk.key.Y)
}

func (pk *PublicKey) Address() types.Address {
	h := sha256.Sum256(pk.toSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (s *Signature) Verify(pk PublicKey, data []byte) bool {
	return ecdsa.Verify(pk.key, data, s.r, s.s)
}
