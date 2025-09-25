package core

type storage interface {
	SaveBlock(block *Block) error
}

type inMemoryStorage struct {
}

func NewInMemoryStorage() *inMemoryStorage {
	return &inMemoryStorage{}
}

func (s *inMemoryStorage) SaveBlock(block *Block) error {
	return nil
}
