package QueueSet

import (
	"fmt"
	"sync"
)

//Element stands for the type of elements in QueueSet.
type Element interface {
	GetID() string //GetID should returns unique ID of the element.
}

//QueueSet is the major struct of this package.
type QueueSet struct {
	queue chan uint64       //元素队列，此队列中的值表示元素在list中的位置
	list  []Element         //记录QueueSet中的所有元素。如果某个元素中途取消排队，对应位置Nil
	listp uint64            //list的队尾
	size  uint64            //QueueSet的大小
	loc   map[string]uint64 //记录ID对应的元素在list中的位置
	n     uint64            //记录元素个数s

	mu      *sync.RWMutex
	queueMu *sync.Mutex //同时只能有一个pop操作在运行
}

//New returns a pointer to a QueueSet.
func New(size uint64) *QueueSet {
	return &QueueSet{
		queue:   make(chan uint64, size-1),
		list:    make([]Element, size),
		listp:   0,
		size:    size,
		loc:     make(map[string]uint64, size),
		mu:      new(sync.RWMutex),
		queueMu: new(sync.Mutex),
	}
}

//Push an element.
func (qs *QueueSet) Push(e Element) {
	qs.mu.Lock() //Push是原子操作
	defer qs.mu.Unlock()
	fmt.Println("Pushing:", e)
	fmt.Println("Before push:", qs.list, qs.n, qs.listp)
	qs.queue <- qs.listp                              //入队列和之后的list修改不可被打断
	if lastloc, exists := qs.loc[e.GetID()]; exists { //如果已存在一个
		qs.list[lastloc] = nil //就清除前一个
		qs.n--
	}
	qs.loc[e.GetID()] = qs.listp        //记录位置
	qs.list[qs.listp] = e               //放入列表
	qs.listp = (qs.listp + 1) % qs.size //队尾后移一位
	qs.n++
	fmt.Println("After push:", qs.list, qs.n, qs.listp)
}

//Pop an element.
func (qs *QueueSet) Pop() Element {
	for {
		qs.popMu.Lock()
		loc := <-qs.queue //先出队列,对Push原子操作的结果无影响
		fmt.Println("Popping:", loc)
		qs.mu.Lock() //等待Push原子操作完成
		if e := qs.list[loc]; e != nil {
			fmt.Println("Before pop:", qs.list, qs.n)
			delete(qs.loc, e.GetID())
			qs.list[loc] = nil
			qs.n--
			qs.mu.Unlock()
			qs.popMu.Unlock()
			fmt.Println("After pop:", qs.list, qs.n)
			return e
		}
		qs.mu.Unlock()
		qs.popMu.Unlock()
	}
}

//Delete an element.
func (qs *QueueSet) Cancel(id string) {
	qs.mu.Lock() //Cancel是原子操作
	defer qs.mu.Unlock()
	fmt.Println("Before cancel:", qs.list, qs.n)
	if loc, exists := qs.loc[id]; exists {
		qs.list[loc] = nil
		delete(qs.loc, id)
		qs.n--
	}
	fmt.Println("After cancel:", qs.list, qs.n)
}

//Count returns how much elements are there in the QueueSet.
func (qs *QueueSet) Count() uint64 {
	return qs.n
}

//Exists returns the existence of an element.
func (qs *QueueSet) Exists(id string) bool {
	qs.mu.Lock()
	defer qs.mu.Unlock()
	_, exists := qs.loc[id]
	return exists
}

//GetQueueElements returns all the elements in the QueueSet.
func (qs *QueueSet) GetQueueElements() []Element {
	qs.mu.RLock()
	defer qs.mu.RUnlock()
	if qs.n <= 0 {
		return []Element{}
	}
	j := qs.n
	elements := make([]Element, qs.n)
	for i := (qs.listp - 1) % qs.size; j > 0; i = (i - 1) % qs.size {
		if qs.list[i] != nil {
			elements[j-1] = qs.list[i]
			j--
		}
	}
	return elements
}
