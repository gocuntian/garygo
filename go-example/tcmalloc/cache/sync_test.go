package cache

import (
	"testing"

	"github.com/funny/utest"
)

func Test_SyncPool_AllocSmall(t *testing.T) {
	pool := NewSyncPool(128, 1024, 2)
	mem := pool.Alloc(64)
	utest.EqualNow(t, len(mem), 64)
	utest.EqualNow(t, cap(mem), 128)
	pool.Free(mem)
}
