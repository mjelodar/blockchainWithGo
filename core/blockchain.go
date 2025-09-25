package core

type Blockchain struct {
	headers   []*Header
	store     storage
	validator Validator
}

func NewBlockchain(genesis *Block) *Blockchain {
	bc := &Blockchain{
		headers: make([]*Header, 0),
		store:   NewInMemoryStorage(),
	}
	bc.validator = NewBlockChainValidator(*bc)
	bc.AddBlockWithoutValidation(genesis)
	return bc
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}
	return bc.AddBlockWithoutValidation(b)
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= uint32(len(bc.headers)-1)
}

func (bc *Blockchain) AddBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return bc.store.SaveBlock(b)
}
