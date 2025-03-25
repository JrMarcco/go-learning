package rpc

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type KVStoreService struct {
	data   map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		data:   make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (k *KVStoreService) Get(ket string, val *string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if v, ok := k.data[ket]; ok {
		*val = v
		return nil
	}

	return errors.New("not found")
}

func (k *KVStoreService) Set(kv [2]string, _ *struct{}) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	key, val := kv[0], kv[1]

	if ov := k.data[key]; ov != val {
		for _, fn := range k.filter {
			fn(key)
		}
	}

	k.data[key] = val
	return nil
}

func (k *KVStoreService) Watch(timeout int, changedKey *string) error {
	id := fmt.Sprintf("wwatch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	k.mu.Lock()
	k.filter[id] = func(key string) {
		ch <- key
	}
	k.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		return errors.New("time out")
	case key := <-ch:
		*changedKey = key
		return nil
	}
}
