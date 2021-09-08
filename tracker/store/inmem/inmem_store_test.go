package inmem

import (
	"testing"

	"github.com/Energy8Token/go-web3/tracker/store"
)

func TestInMemoryStore(t *testing.T) {
	store.TestStore(t, func(t *testing.T) (store.Store, func()) {
		return NewInmemStore(), func() {}
	})
}
