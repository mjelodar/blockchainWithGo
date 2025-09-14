package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/PRACTICING-GO/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func test_header_encode_decode(t *testing.T) {
	// Create a new header
	header := Header{
		Version:       1,
		Height:        10,
		TimeStamp:     uint64(time.Nanosecond),
		PrevBlockHash: types.RandomHash(),
		Nonce:         1234567890,
	}

	// Encode the header to binary
	var buf []byte
	writer := bytes.NewBuffer(buf)
	assert.Nil(t, header.EncodeBinary(writer))

	// Decode the header from binary
	var decodedHeader Header
	reader := bytes.NewReader(writer.Bytes())
	assert.Nil(t, decodedHeader.DecodeBinary(reader))
	assert.Equal(t, header, decodedHeader)

}

func test_block_encode_decode(t *testing.T) {
	// Create a new block with a header and transactions
	block := Block{
		Header: Header{
			Version:       1,
			Height:        10,
			TimeStamp:     uint64(time.Nanosecond),
			PrevBlockHash: types.RandomHash(),
			Nonce:         1234567890,
		},
		Transactions: []Transaction{{}, {}}, // Add dummy transactions
	}

	// Encode the block to binary
	var buf []byte
	writer := bytes.NewBuffer(buf)
	assert.Nil(t, block.EncodeBinary(writer))

	// Decode the block from binary
	var decodedBlock Block
	reader := bytes.NewReader(writer.Bytes())
	assert.Nil(t, decodedBlock.DecodeBinary(reader))
	assert.Equal(t, block.Header, decodedBlock.Header)
	assert.Equal(t, block, decodedBlock)
	fmt.Printf("Decoded Block: %+v\n", decodedBlock.DecodeBinary(reader))
}

func TestBlockHash(t *testing.T) {
	// Create a new block with a header and transactions
	block := Block{
		Header: Header{
			Version:       1,
			Height:        10,
			TimeStamp:     uint64(time.Nanosecond),
			PrevBlockHash: types.RandomHash(),
			Nonce:         1234567890,
		},
		Transactions: []Transaction{{}, {}}, // Add dummy transactions
	}

	// Calculate the hash of the block
	hash := block.Hash()

	// Verify that the hash is not zero
	assert.False(t, hash.IsZero(), "Block hash should not be zero")
	fmt.Printf("Block Hash: %s\n", hash.String())
}
