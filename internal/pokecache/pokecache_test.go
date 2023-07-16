package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 100
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("Expected cache to be created")
	}
}

func TestAddToCache(t *testing.T) {
	interval := time.Millisecond * 100
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("Expected to find key %s in cache", c.inputKey)
			continue
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf("Expected %s, got %s", string(c.inputVal), string(actual))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 100
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond*10)

	_, ok1 := cache.Get(keyOne)
	if ok1 {
		t.Errorf("Expected key %s to be reaped", keyOne)
	}

	keyTwo := "key2"
	cache.Add(keyTwo, []byte("val2"))

	_, ok2 := cache.Get(keyTwo)
	if !ok2 {
		t.Errorf("Expected key %s to be in cache", keyTwo)
	}

}
