package SortedSet

import (
	"fmt"
	"math/rand"
	"testing"
)

type testObj struct {
	id int
}

func (o testObj) GetName() string {
	return fmt.Sprintf("I'm %02d", o.id)
}

func TestSortedSet(t *testing.T) {
	zset := New(30)
	for i := 0; i < 20; i++ {
		e := new(testObj)
		e.id = rand.Int()
		zset.Update(e, rand.Float64())
	}
	var sorted = zset.Sorted(16)
	for _, e := range sorted {
		w, _ := zset.GetWeight(e)
		fmt.Printf("\n%s: %.6f", e.GetName(), w)
	}
	zset.DeltaUpdateAll(10)
	sorted = zset.SortedAll()
	for _, e := range sorted {
		w, _ := zset.GetWeight(e)
		fmt.Printf("\n%s: %.6f", e.GetName(), w)
	}

	zset.Update(testObj{id: 100}, 10.5)
	fmt.Print(zset.String())
	zset.DeltaUpdate(testObj{id: 100}, 0.1)
	fmt.Print(zset.String())
}
