package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cases := []struct {
		key   string
		input []byte
	}{
		{
			key:   "one",
			input: []byte{72, 101, 108, 108, 111},
		},
		{
			key:   "two",
			input: []byte{72, 101, 5, 1, 111},
		},
		{
			key:   "three",
			input: []byte{1, 77},
		},
	}

	cache := NewCache(500 * time.Millisecond)

	for _, c := range cases {
		cache.Add(c.key, c.input)
	}

	for _, c := range cases {
		actual, ok := cache.Get(c.key)
		expected := c.input
		if !ok {
			t.Errorf("Cached entry does not exists!")
		}

		if len(actual) != len(expected) {
			t.Errorf("Cached data inconsistent: input %v, expected %v", actual, expected)
			return
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Cached data inconsistent: input %v, expected %v", actual, expected)
			}
		}
	}

	time.Sleep(1 * time.Second)

	for _, c := range cases {
		_, exists := cache.Get(c.key)
		if exists {
			t.Errorf("Cache should be empty at this point!")
		}
	}
}
