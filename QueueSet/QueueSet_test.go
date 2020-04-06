package QueueSet

import (
	"fmt"
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

func TestQueueSet(t *testing.T) {
	qs := New(5)
	qs.Push(TestElement{ID: "A", Count: 1})
	go qs.Push(TestElement{ID: "B", Count: 1})
	go qs.Push(TestElement{ID: "A", Count: 2})
	go qs.Push(TestElement{ID: "A", Count: 3})
	go qs.Push(TestElement{ID: "B", Count: 2})
	go qs.Push(TestElement{ID: "C", Count: 1})
	go qs.Push(TestElement{ID: "C", Count: 2})
	fmt.Println(qs.Pop())
	go func() { fmt.Println(qs.Pop()) }()
	go qs.Cancel("A")
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	go func() { fmt.Println(qs.Pop()) }()
	time.Sleep(1e8)
	//func() { fmt.Println(qs.Pop()) }()
	time.Sleep(1e8)
}
