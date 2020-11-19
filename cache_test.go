package gocachebenchmarks

import (
	"crypto/rand"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	ristretto "github.com/dgraph-io/ristretto"
	hashicorp "github.com/hashicorp/golang-lru"
	koding "github.com/koding/cache"
	"github.com/muesli/cache2go"
	cmap "github.com/orcaman/concurrent-map"
	cache "github.com/patrickmn/go-cache"
)

func toKey(i int) string {
	return fmt.Sprintf("item:%d", i)
}

var testStr string

func BenchmarkCache(b *testing.B) {
	sizes := []int64{64, 1024, 10240, 1024 * 1024}
	for _, n := range sizes {
		testBytes := make([]byte, n)
		readN, err := rand.Read(testBytes)
		if readN != int(n) {
			panic(fmt.Sprintf("expect %d but got %d", n, readN))
		}
		if err != nil {
			panic(err)
		}
		testStr = string(testBytes)

		b.Run(fmt.Sprintf("Kodingcache-%d", n), benchmarkKodingCache)
		b.Run(fmt.Sprintf("HashicorpLRU-%d", n), benchmarkHashicorpLRU)
		b.Run(fmt.Sprintf("Cache2Go-%d", n), benchmarkCache2Go)
		b.Run(fmt.Sprintf("GoCache-%d", n), benchmarkGoCache)
		b.Run(fmt.Sprintf("Freecache-%d", n), benchmarkFreecache)
		b.Run(fmt.Sprintf("BigCache-%d", n), benchmarkBigCache)
		b.Run(fmt.Sprintf("Ristretto-%d", n), benchmarkRistretto)
		b.Run(fmt.Sprintf("SyncMap-%d", n), benchmarkSyncMap)
		b.Run(fmt.Sprintf("ConcurrentMap-%d", n), BenchmarkConcurrentMap)
		fmt.Println()
	}

}

func benchmarkKodingCache(b *testing.B) {
	c := koding.NewMemoryWithTTL(time.Duration(60) * time.Second)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func benchmarkHashicorpLRU(b *testing.B) {
	// c := cache2go.Cache("test")
	c, _ := hashicorp.New(10)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			value, err := c.Get(toKey(i))
			if err == true {
				_ = value
			}
		}
	})
}

func benchmarkCache2Go(b *testing.B) {
	c := cache2go.Cache("test")

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), 1*time.Minute, testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Value(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func benchmarkGoCache(b *testing.B) {
	c := cache.New(1*time.Minute, 5*time.Minute)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), testStr, cache.DefaultExpiration)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := c.Get(toKey(i))
			if found {
				_ = value
			}
		}
	})
}

func benchmarkFreecache(b *testing.B) {
	c := freecache.NewCache(1024 * 1024 * 5)

	data := []byte(testStr)
	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set([]byte(toKey(i)), data, 60)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get([]byte(toKey(i)))
			if err == nil {
				_ = value
			}
		}
	})
}

func benchmarkBigCache(b *testing.B) {
	c, _ := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 10,
	})

	data := []byte(testStr)
	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), data)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func benchmarkGCache(b *testing.B) {
	c := gcache.New(b.N).LRU().Build()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func benchmarkRistretto(b *testing.B) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(toKey(i), testStr, 1)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, ok := cache.Get(toKey(i))
			if ok {
				_ = value
			}
		}
	})

}

// No expire, but helps us compare performance
func benchmarkSyncMap(b *testing.B) {
	var m sync.Map

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(toKey(i), testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := m.Load(toKey(i))
			if found {
				_ = value
			}
		}
	})
}

func BenchmarkConcurrentMap(b *testing.B) {
	var m = cmap.New()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Set(toKey(i), testStr)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := m.Get(toKey(i))
			if found {
				_ = value
			}
		}
	})
}
