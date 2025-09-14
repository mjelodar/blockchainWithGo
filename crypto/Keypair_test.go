package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	// Generate a new private key
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	// address := publicKey.Address()

	msg := []byte("Hello, Blockchain!")
	signature, _ := privateKey.Sign(msg)
	// if err != nil {
	// 	t.Fatalf("Failed to sign message: %v", err)
	// }
	// fmt.Printf("Signature: r=%s, s=%s\n", signature.r.String(), signature.s.String())

	b := signature.Verify(publicKey, msg)
	assert.True(t, b, "Signature verification failed")
}
