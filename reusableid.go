package reusableid

import (
	"sync"
)

type IdGenerator struct {
	sync.Mutex
	currentId uint64
	reusables map[uint64]bool
}

func NewIdGenerator() *IdGenerator {
	ret := &IdGenerator{
		currentId: 0,
		reusables: make(map[uint64]bool),
	}
	return ret
}

func (idGenerator *IdGenerator) Next() (id uint64) {
	idGenerator.Lock()
	defer idGenerator.Unlock()

	for key := range idGenerator.reusables {
		delete(idGenerator.reusables, key)
		return key
	}

	currentId := idGenerator.currentId
	idGenerator.currentId += 1
	return currentId
}

func (idGenerator *IdGenerator) Recycle(id uint64) {
	idGenerator.Lock()
	defer idGenerator.Unlock()
	idGenerator.reusables[id] = true
}
