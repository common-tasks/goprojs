package datapartition

import (
	"fmt"
	"sync"
)

type KeyValueStore struct {
	data  map[string]string
	mutex sync.RWMutex
}

func (kv *KeyValueStore) Get(key string) (string, bool) {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()
	value, ok := kv.data[key]
	return value, ok
}
func (kv *KeyValueStore) Set(key string, value string) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	kv.data[key] = value
}

type DataPartitions struct {
	partitions []*KeyValueStore
}

func (dp *DataPartitions) GetPartition(key string) *KeyValueStore {
	partition := dp.partitions[hashCode(key)%len(dp.partitions)]
	return partition
}
func NewDataPartition(num int) *DataPartitions {
	partitions := make([]*KeyValueStore, num)
	for i := 0; i < num; i++ {
		partitions[i] = &KeyValueStore{data: make(map[string]string)}
	}
	return &DataPartitions{partitions: partitions}
}

func hashCode(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}

func Start() {
	dataPartition := NewDataPartition(3)

	// Set keys
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		partition := dataPartition.GetPartition(key)
		partition.Set(key, value)
	}

	// Get keys
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		partition := dataPartition.GetPartition(key)
		value, _ := partition.Get(key)
		fmt.Printf("%s -> %s\n", key, value)
	}
}
