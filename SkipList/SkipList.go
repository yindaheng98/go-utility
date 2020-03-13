package SkipList

import (
	"math"
	"time"
)

//This is the struct of skip list.
type SkipList struct {
	root      *Node      //根节点指针
	n         uint64     //节点总数
	level     uint64     //预估的索引层数
	randLevel *RandLevel //索引层数生成器
}

//Returns a pointer to a SkipList.
//
//"listSize" is the initial size of the SkipList.
//"indexLevel" is the max level of the index.
func NewWithLevel(listSize, indexLevel uint64) *SkipList {
	C := uint64(math.Ceil(math.Pow(float64(listSize), 1.0/float64(indexLevel))))
	return &SkipList{nil, 0, indexLevel + 1,
		NewRandomLevel(C, indexLevel, time.Now().UnixNano())}
}

//Returns a pointer to a SkipList.
//
//"listSize" is the initial size of the SkipList.
//"C" is the decade factor of the index ([index in level n]=[index in level n-1]/C).
func NewWithC(listSize, C uint64) *SkipList {
	indexLevel := uint64(math.Ceil(math.Log(float64(listSize)) / math.Log(float64(C))))
	return &SkipList{nil, 0, indexLevel + 1,
		NewRandomLevel(C, indexLevel, time.Now().UnixNano())}
}

//Returns the number of nodes in SkipList.
func (sl *SkipList) Count() uint64 {
	return sl.n
}

//Find the pointer to max nodes whose value < data
func (sl *SkipList) Find(data float64) *Node {
	result := sl.find(data)
	if result == nil || len(result) < 1 {
		return nil
	}
	return result[0]
}

//找到各层index中大小小于data的最大节点的指针
func (sl *SkipList) find(data float64) []*Node {
	if sl.root == nil { //如果链表为空
		return nil //则直接返回空
	}

	//链表不为空才能开始初始化
	level := len(sl.root.next)     //根节点索引层数即时最大索引层数
	result := make([]*Node, level) //初始化结果index表
	if data < sl.root.data {       //如果链表中没有这样的节点就直接返回
		return result
	}

	//有这样的节点才能开始查找
	p := sl.root     //初始化当前指针
	pLevel := level  //初始化当前指针所在层数
	for pLevel > 0 { //循环直到pLevel到了第0层
		pLevel -= 1                           //index向下走一层
		next := p.next[pLevel]                //初始化该层的下一个节点指针
		for next != nil && next.data < data { //如果后面有节点并且其值比data小
			p = next //就往后走一步
			next = p.next[pLevel]
		} //走到头了就退出，此时的p即第pLevel层要找的节点指针
		result[pLevel] = p //记录这个指针
	}
	return result
}

//Insert a value.
//Returns the pointer to the node where the value is inserted.
func (sl *SkipList) Insert(data float64) *Node {
	sl.n++
	pres := sl.find(data)      //查找插入点
	presN := uint64(len(pres)) //插入节点的数量

	if pres == nil { //查找返回了空，说明链表为空
		sl.root = newNode(data, sl.level) //那就直接给root赋值
		return sl.root
	}

	//链表不为空才能开始正常赋值
	level := sl.randLevel.Rand() + 1
	result := newNode(data, level) //要返回的值（永远返回新指针）
	insert := result               //要执行插入操作的值

	//返回的第一个指针就为空
	if pres[0] == nil { //说明要在根节点前插
		insert.prev = sl.root.prev
		insert.next = sl.root.next //首先复制根节点的前后指针
		sl.root.prev = make([]*Node, level)
		sl.root.next = make([]*Node, level)     //然后重建根节点的前后指针
		insert = sl.root                        //“偷梁换柱”：把根节点值提出来作为要插入的值
		sl.root = result                        //然后将根节点值替换为新值
		for i := uint64(0); i < sl.level; i++ { //然后更新新的根节点的后置节点的前置指针
			if sl.root.next[i] != nil {
				sl.root.next[i].prev[i] = sl.root
			}
		}
		for i := uint64(0); i < presN; i++ { //然后更新前置节点表
			pres[i] = sl.root
		}
	}

	//最后执行插入操作
	for i := uint64(0); i < level; i++ {
		insert.prev[i] = pres[i]
		insert.next[i] = pres[i].next[i]
		pres[i].next[i] = insert
		if insert.next[i] != nil {
			insert.next[i].prev[i] = insert
		}
	}
	return result
}

//Returns list of the minimum n nodes sorted by their value in ascending order.
func (sl *SkipList) Traversal(n uint64) []*Node {
	if sl.n < n {
		n = sl.n
	}
	result := make([]*Node, n)
	node := sl.root
	for i := uint64(0); i < n && node != nil; i++ {
		result[i] = node
		node = node.next[0]
	}
	return result
}

//Returns list of the nodes sorted by their value in ascending order.
func (sl *SkipList) TraversalAll() []*Node {
	return sl.Traversal(sl.n)
}

//Delete a node.
func (sl *SkipList) Delete(node *Node) {
	sl.n--
	defer func(toDestory *Node) {
		toDestory.prev = make([]*Node, len(toDestory.prev))
		toDestory.next = make([]*Node, len(toDestory.next))
	}(node)

	length := len(node.prev)
	for i := 0; i < length; i++ {
		if node.prev[i] != nil {
			node.prev[i].next[i] = node.next[i]
		}
		if node.next[i] != nil {
			node.next[i].prev[i] = node.prev[i]
		}
	}

	if node == sl.root { //如果要删除的是个根节点
		if sl.root.next[0] == nil { //而表中只有这一个节点
			sl.root = nil //就直接删除根节点
			return
		}
		//但如果表中不止有这一个节点，则还需要进行进一步处理
		prev := sl.root.prev
		next := sl.root.next
		node = node.next[0] //以根节点的下一个节点为新根节点
		for i, n := range node.next {
			next[i] = n //首先融合根节点和新根节点的next表构造一个新的next表
		}
		node.prev = prev
		node.next = next //给新根节点的前后节点表赋值
		for i := range node.next {
			if node.next[i] != nil {
				node.next[i].prev[i] = node //然后重建根节点的后继节点的前继记录
			}
		}
		sl.root = node //最后更换根节点的记录
	}
}

//Increase the value of the node by delta, and return where the new value is inserted.
func (sl *SkipList) Delta(node *Node, delta float64) *Node {
	sl.Delete(node)
	return sl.Insert(node.data + delta)
}

//Increase the value of all nodes by delta.
func (sl *SkipList) DeltaAll(delta float64) {
	node := sl.root
	for i := uint64(0); i < sl.n && node != nil; i++ {
		node.data += delta
		node = node.next[0]
	}
}
