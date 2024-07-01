package cache

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/barretodotcom/cache-lru/utils"
)

func TestCache(t *testing.T) {
	dataProvider := []utils.DataProvider{{Expected: &Item{key: "2", value: 10, time: time.Now()}, Key: "2"}, {Expected: nil, Key: "1"}}
	c := New(1)
	c.Set("1", 10)
	c.Set("2", 10)

	for _, d := range dataProvider {
		if !reflect.DeepEqual(d.Expected, c.Get(d.Key)) {
			fmt.Println(c.Get(d.Key), d.Expected)
			t.Error("test failed, key: ", d.Key)
		}
	}
}
