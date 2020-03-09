package TimeoutMap

import (
	"github.com/yindaheng98/go-utility/TimeoutMap/TimeoutValue"
	"sync"
	"time"
)

type TimeoutMap struct {
	elements map[string]*TimeoutValue.TimeoutValue //存储数据
	mu       *sync.RWMutex                         //数据读写锁
}

//输入超时时间和删除缓存的数量新建发送器列表
func New() *TimeoutMap {
	return &TimeoutMap{make(map[string]*TimeoutValue.TimeoutValue), new(sync.RWMutex)}
}

//通过id进行更新，仅更新时间
func (m *TimeoutMap) UpdateID(id string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.elements[id] //查询此id是否存在
	if exists {                     //如果存在
		value.Update(nil, value.GetTimeout()) //则更新
	}
}

//通过一个Element进行更新，更新存储的信息
func (m *TimeoutMap) UpdateInfo(el Element, timeout time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := el.GetID()                //先获取发送器信息中的id
	value, exists := m.elements[id] //查询此id是否存在
	if exists {                     //如果存在
		value.Update(el, timeout) //则更新
	} else { //否则新建
		value := TimeoutValue.New(el, timeout)
		m.elements[id] = value
		go func() {
			value.GetElement().(Element).NewAddedHandler()
			value.Run()
			m.mu.Lock()
			defer m.mu.Unlock()
			delete(m.elements, id) //退出时删除
			value.GetElement().(Element).DeletedHandler()
		}()
	}
}

func (m *TimeoutMap) Delete(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.elements[id] //先查找
	if ok {                     //如果找得到
		value.Stop() //则使其停止即删除
	}
}

func (m *TimeoutMap) DeleteAll() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, value := range m.elements {
		value.Stop() //则使其停止即删除
	}
}

//获取某个id对应的信息
func (m *TimeoutMap) GetElement(id string) (Element, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var el Element = nil
	value, ok := m.elements[id] //先查找
	if ok {                     //如果找得到
		el = value.GetElement().(Element) //则返回结果
	}
	return el, ok
}

//获取所有的信息
func (m *TimeoutMap) GetAll() []Element {
	m.mu.RLock()
	defer m.mu.RUnlock()
	els := make([]Element, len(m.elements))
	i := 0
	for _, value := range m.elements {
		els[i] = value.GetElement().(Element)
		i += 1
	}
	return els
}

func (m *TimeoutMap) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.elements)
}
