package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("cache size check", func(t *testing.T) {
		cache := NewCache(2)
		cache.Set("key1", 1)
		cache.Set("key2", 2)
		cache.Set("key3", 3)
		_, exist := cache.Get("key1")
		require.False(t, exist)
	})

	t.Run("unused element deletion", func(t *testing.T) {
		cache := NewCache(3)
		cache.Set("key1", 1)
		cache.Set("key2", 2)
		cache.Set("key3", 3)

		cache.Get("key1")
		cache.Set("key2", 22)
		cache.Set("key1", 11)
		cache.Get("key1")
		cache.Get("key1")
		cache.Set("key1", 111)
		cache.Get("key2")

		cache.Set("key4", 4)
		_, exist := cache.Get("key3")
		require.False(t, exist)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Parallel()
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
