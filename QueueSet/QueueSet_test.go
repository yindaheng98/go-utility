package QueueSet

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type TestElement struct {
	ID    string
	Count uint
}

func (t TestElement) GetID() string {
	return t.ID
}

var IDs = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var count = uint(1)

func TestQueueSet(t *testing.T) {
	qs := New()
	qs.Push(TestElement{ID: "A", Count: 1})
	fmt.Println(qs.Pop())
	for i := 0; i < 50; i++ {
		ID := string(IDs[rand.Uint32()%uint32(len(IDs))])
		go qs.Push(TestElement{ID: ID, Count: count})
		count++
	}
	for i := 0; i < 50; i++ {
		go qs.Cancel(string(IDs[rand.Uint32()%uint32(len(IDs))]))
	}
	for i := 0; i < 50; i++ {
		go func() { fmt.Println(qs.Pop(), qs.GetQueueElements()) }()
		count++
	}
	//go qs.Cancel("A")
	time.Sleep(5e8)
	time.Sleep(5e8)
	func() { fmt.Println(qs.Pop()) }()
	time.Sleep(5e8)
}
