package SkipList

import (
	"math/rand"
	"testing"
)

func TestSkipList(t *testing.T) {
	skiplist := NewWithLevel(30, 5)
	t.Log(skiplist.Find(100))
	for i := 0; i < 20; i++ {
		t.Log(skiplist.Insert(rand.Float64() * 100))
	}
	for i := 0; i < 20; i++ {
		np := skiplist.Insert(rand.Float64() * -100)
		if np.data <= -50.0 {
			skiplist.Delete(np)
		}
	}
	t.Log(skiplist.root)
	sorted := skiplist.Traversal(30)
	t.Log(len(sorted))
	t.Log(sorted)
	for _, node := range sorted {
		t.Log(node)
	}
	skiplist.DeltaAll(100)
	sorted = skiplist.TraversalAll()
	t.Log(sorted)
	for _, node := range sorted {
		t.Log(node)
	}
	node := skiplist.Insert(150)
	t.Log(skiplist.Delta(node, 10))
	sorted = skiplist.TraversalAll()
	t.Log(sorted)
	for _, node := range sorted {
		t.Log(node)
	}
}
