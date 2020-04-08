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
	queue   []Element         //队列
	loc     map[string]uint64 //记录ID对应的元素在队列中的位置
	lastloc uint64            //队尾指针
	pushed  chan bool         //每当有元素push入队列，就向这个通道发信息

	mu *sync.RWMutex
}

//New returns a pointer to a QueueSet.
func New() *QueueSet {
	return &QueueSet{[]Element{}, map[string]uint64{}, 0, make(chan bool, 1), new(sync.RWMutex)}
}

//Push an element.
func (qs *QueueSet) Push(e Element) {
	qs.mu.Lock() //Push是原子操作
	defer qs.mu.Unlock()
	defer func() {
		select {
		case qs.pushed <- true:
		default:
		}
	}() //每当有元素push入队列，就向通道发信息
	fmt.Println("Pushing:", e)
	fmt.Println("Before push:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
	if loc, exists := qs.loc[e.GetID()]; exists { //如果已存在前一个值，需要进行特殊处理
		if loc+1 >= qs.lastloc { //如果已经是最末尾的值
			qs.queue[loc] = e //直接改这个值就好
			return
		}
		//不是最末尾的值
		qs.queue[loc] = nil  //那就将原来的值置空
		if qs.lastloc <= 1 { //如果队尾指针在最开头
			qs.lastloc = 0 //直接置0
		} else { //否则
			for ; qs.lastloc > 0 && qs.queue[qs.lastloc-1] == nil; qs.lastloc-- {
				//循环直到qs.lastloc指向qs.queue中的最后一个为空的位置
			}
		}
	}
	qlength := uint64(len(qs.queue))
	if qs.lastloc >= qlength { //队列不够长
		qs.queue = append(qs.queue, make([]Element, qs.lastloc-qlength+10)...) //就先扩展
	}
	qs.queue[qs.lastloc] = e       //入队列
	qs.loc[e.GetID()] = qs.lastloc //记录位置
	qs.lastloc++                   //队尾指针后移一位
	fmt.Println("After push:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
}

//Pop an element.
func (qs *QueueSet) Pop() Element {
	for {
		qs.mu.Lock()                              //Pop是原子操作
		for i := uint64(0); i < qs.lastloc; i++ { //从开头开始遍历
			if e := qs.queue[i]; e != nil { //找到一个非空元素
				fmt.Println("Popping:", e)
				fmt.Println("Before pop:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
				qs.queue[i] = nil
				delete(qs.loc, e.GetID())
				if len(qs.loc) <= 0 { //如果空了
					qs.queue = []Element{} //那就直接置空
					qs.lastloc = 0
				} else if i >= 10 { //如果开头的空值超过10个
					qs.queue = qs.queue[i+1:] //就切掉开头这些空值
					qs.lastloc -= i           //然后修改队尾指针
					for _, ee := range qs.queue {
						if ee != nil {
							qs.loc[ee.GetID()] -= i //并且重载后面所有元素位置记录
						}
					}
				}
				fmt.Println("After pop:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
				qs.mu.Unlock()
				return e //找到非空元素即返回
			}
		}
		qs.mu.Unlock()
		<-qs.pushed //找不到就等下一个push
	}
}

//Delete an element.
func (qs *QueueSet) Cancel(id string) {
	qs.mu.Lock() //Cancel是原子操作
	defer qs.mu.Unlock()
	fmt.Println("Before cancel:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
	if loc, exists := qs.loc[id]; exists {
		qs.queue[loc] = nil
		delete(qs.loc, id)
	}
	fmt.Println("After cancel:", qs.lastloc, qs.queue, len(qs.queue), qs.loc, len(qs.loc))
}

//Count returns how much elements are there in the QueueSet.
func (qs *QueueSet) Count() uint64 {
	qs.mu.RLock()
	defer qs.mu.RUnlock()
	return uint64(len(qs.loc))
}

//Exists returns the existence of an element.
func (qs *QueueSet) Exists(id string) bool {
	qs.mu.RLock()
	defer qs.mu.RUnlock()
	_, exists := qs.loc[id]
	return exists
}

//GetQueueElements returns all the elements in the QueueSet.
func (qs *QueueSet) GetQueueElements() []Element {
	qs.mu.RLock()
	defer qs.mu.RUnlock()
	if len(qs.loc) <= 0 {
		return []Element{}
	}

	elements := make([]Element, len(qs.loc))
	for i, j := uint64(0), uint64(0); i < qs.lastloc; i++ {
		if qs.queue[i] != nil {
			elements[j] = qs.queue[i]
			j++
		}
	}
	return elements
}
