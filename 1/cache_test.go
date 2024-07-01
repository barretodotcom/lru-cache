package cache

import (
	"testing"

	"github.com/barretodotcom/cache-lru/utils"
)

func TestSet(t *testing.T) {
	providers := []utils.DataProvider{{Expected: 2, Key: "2"}, {Expected: -1, Key: "1"}}
	c := New(1)

	c.Set("1", 1)
	c.Set("2", 2)

	for _, p := range providers {
		if p.Expected != c.Get(p.Key) {
			t.Error("test failed, key: ", p.Key)
		}
	}
}
