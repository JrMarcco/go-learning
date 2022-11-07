package once

import (
	"sync"
	"testing"
)

func TestOnceDL(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
		})
	})
}
