package cache

import "sync"

type SyncPool struct {
	classes     []sync.Pool
	classesSize []int
	minSize     int
	maxSize     int
}

func NewSyncPool(minSize, maxSize, factor int) *SyncPool {
	n := 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		n++
	}

	pool := &SyncPool{
		make([]sync.Pool, n),
		make([]int, n),
		minSize, maxSize,
	}
	n = 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		pool.classesSize[n] = chunkSize
		pool.classes[n].New = func(size int) func() interface{} {
			return func() interface{} {
				buf := make([]byte, size)
				return &buf
			}
		}(chunkSize)
		n++
	}
	return pool
}

func (pool *SyncPool) Alloc(size int) []byte {
	if size <= pool.maxSize {
		for i := 0; i < len(pool.classesSize); i++ {
			men := pool.classes[i].Get().(*[]byte)
			return (*men)[:size]
		}
	}
	return make([]byte, size)
}

func (pool *SyncPool) Free(men []byte) {
	if size := cap(men); size <= pool.maxSize {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= size {
				pool.classes[i].Put(&men)
				return
			}
		}
	}
}
