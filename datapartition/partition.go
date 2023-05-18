package datapartition

import (
	"sync"
)

/*
To simulate the data partitioning
*/

type Store struct {
	DataSet   map[int]string
	storeLock sync.RWMutex
}

type Partitions struct {
	store []Store
}

func (s *Partitions) NumberOfStores() int {
	count := len(s.store)
	return count
}
func (store *Store) Add(id int, data string) {
	store.storeLock.RLocker().Lock()
	defer store.storeLock.RLocker().Unlock()
	store.DataSet[id] = data
}
func (store *Store) Get(id int) (string, bool) {
	store.storeLock.RLock()
	defer store.storeLock.RUnlock()
	str, ok := store.DataSet[id]
	return str, ok
}
func CreatePartitions(count int) *Partitions {
	stores := make([]*Store, count)

	for i := 0; i < count; i++ {
		stores[i]= &Store{DataSet: make(map[int]string)}
	}
}
