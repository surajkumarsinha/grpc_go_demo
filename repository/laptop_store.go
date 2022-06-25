package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/surajkumarsinha/go_grpc_demo/pb/messages"
)
var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *messages.Laptop) error
	Find(id string) (*messages.Laptop, error)
}

// In memory Laptop Store
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data map[string]*messages.Laptop
}

// DB Laptop Store
type DbLaptopStore struct {
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*messages.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *messages.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	} 

	// Deep copy of the laptop object
	other := &messages.Laptop{}
	err := copier.Copy(other, laptop)

	if err != nil {
		return fmt.Errorf("Cannot copy laptop data %w", err)
	}

	store.data[other.Id] = other
	return nil
}

// Find finds a laptop by ID
func (store *InMemoryLaptopStore) Find(id string) (*messages.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deepCopy(laptop)
}


func deepCopy(laptop *messages.Laptop) (*messages.Laptop, error) {
	other := &messages.Laptop{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}