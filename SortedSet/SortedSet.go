package SortedSet

import (
	"fmt"
	"github.com/yindaheng98/go-utility/SkipList"
)

//一个用跳表和hashmap实现的有序集合
type SortedSet struct {
	skiplist     *SkipList.SkipList
	whosStringIs map[string]*SkipList.Node  //序列化Element->*node的map
	whosNodeIs   map[*SkipList.Node]Element //*node->*Element的map
}

func New(size uint64) *SortedSet {
	return &SortedSet{SkipList.NewWithC(size, 2),
		make(map[string]*SkipList.Node),
		make(map[*SkipList.Node]Element)}
}

//向集合中更新一个元素
func (set *SortedSet) Update(obj Element, weight float64) {
	str := obj.GetName()
	set.remove(str)
	nodep := set.skiplist.Insert(weight)
	set.whosStringIs[str] = nodep
	set.whosNodeIs[nodep] = obj
}

//改变集合中一个元素的值
func (set *SortedSet) DeltaUpdate(obj Element, delta float64) {
	str := obj.GetName()
	if oldnodep, ok := set.whosStringIs[str]; ok {
		newnodep := set.skiplist.Delta(oldnodep, delta)
		set.whosStringIs[str] = newnodep
		delete(set.whosNodeIs, oldnodep)
		set.whosNodeIs[newnodep] = obj
	}
}

//集合中所有元素的值加一个delta
func (set *SortedSet) DeltaUpdateAll(delta float64) {
	set.skiplist.DeltaAll(delta)
}

//从集合中删除一个元素
func (set *SortedSet) Remove(obj Element) {
	set.remove(obj.GetName())
}

func (set *SortedSet) remove(str string) {
	if nodep, ok := set.whosStringIs[str]; ok {
		set.skiplist.Delete(nodep)
		delete(set.whosStringIs, str)
		delete(set.whosNodeIs, nodep)
	}
}

func (set *SortedSet) GetWeight(obj Element) (float64, bool) {
	nodep, ok := set.whosStringIs[obj.GetName()]
	if ok {
		return nodep.GetData(), true
	}
	return 0, false
}

func (set *SortedSet) Sorted(n uint64) []Element {
	return set.nodepsToElements(set.skiplist.Traversal(n))
}

func (set *SortedSet) SortedAll() []Element {
	return set.nodepsToElements(set.skiplist.TraversalAll())

}

func (set *SortedSet) nodepsToElements(nodeps []*SkipList.Node) []Element {
	length := len(nodeps)
	result := make([]Element, length)
	for i := 0; i < length; i++ {
		result[i] = set.whosNodeIs[nodeps[i]]
	}
	return result
}

func (set *SortedSet) String() string {
	s := ""
	els := set.SortedAll()
	for _, el := range els {
		w, _ := set.GetWeight(el)
		s += fmt.Sprintf("\t%s: %f,\n", el.GetName(), w)
	}
	return fmt.Sprintf("SortedSet{\n%s}", s)
}
