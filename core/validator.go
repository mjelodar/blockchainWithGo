package core

import "fmt"

type Validator interface {
	ValidateBlock(block *Block) error
}

type BlockChainValidator struct {
	bc Blockchain
}

func NewBlockChainValidator(bc Blockchain) *BlockChainValidator {
	return &BlockChainValidator{bc: bc}
}

func (v *BlockChainValidator) ValidateBlock(block *Block) error {
	if v.bc.HasBlock(block.Header.Height) {
		return fmt.Errorf("block at height %d already exists", block.Header.Height)
	}

	if block.Verify() != nil {
		return fmt.Errorf("invalid block signature")
	}
	return nil
}
