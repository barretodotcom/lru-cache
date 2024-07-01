package cache

import (
	"testing"
)

func TestRemoveLastInsertedElement(t *testing.T) {
	c := New(3)

	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Get("c")
	c.Set("d", 4)

	expectedValues := map[string]any{"a": 3, "b": -1, "c": 3, "d": 4}

	for i := range expectedValues {
		if value := c.Get(i); value != expectedValues[i] {
			t.Errorf("expected key %s, found nil", i)
		}
		return
	}

}
