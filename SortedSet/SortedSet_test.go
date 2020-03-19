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
	zset := New(10)
	zset.Update(new(testObj), rand.Float64())
	for i := 0; i < 20; i++ {
		e := new(testObj)
		e.id = rand.Int()
		zset.Update(e, rand.Float64())
	}
	zset.Remove(new(testObj))
	fmt.Println(zset.String())

	var sorted = zset.SortedAll()
	for _, e := range sorted {
		w, _ := zset.GetWeight(e)
		if w < 0.5 {
			zset.Remove(e)
		}
	}

	zset.DeltaUpdateAll(-10)
	fmt.Println(zset.String())

	t.Log(zset.Count())
	zset.Update(testObj{id: 100}, 10.5)
	t.Log(zset.Count())
	fmt.Print(zset.String() + "\n")
	zset.DeltaUpdate(testObj{id: 100}, 0.1)
	t.Log(zset.Count())
	fmt.Print(zset.String() + "\n")
}
