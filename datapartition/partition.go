package datapartition

import (
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

func Start() {

}
