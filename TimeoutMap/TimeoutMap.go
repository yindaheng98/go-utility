package TimeoutMap

import (
	"github.com/yindaheng98/go-utility/TimeoutMap/TimeoutValue"
	"sync"
	"time"
)

type TimeoutMap struct {
	elements map[string]*TimeoutValue.TimeoutValue //存储数据
	mu       *sync.RWMutex                         //数据读写锁
	mumu     *sync.Mutex                           //同一时刻只能有一个线程在执行获取锁的操作
}

//输入超时时间和删除缓存的数量新建发送器列表
func New() *TimeoutMap {
	return &TimeoutMap{make(map[string]*TimeoutValue.TimeoutValue),
		new(sync.RWMutex), new(sync.Mutex)}
}

//通过id进行更新，仅更新时间
func (m *TimeoutMap) UpdateID(id string) {
	m.mumu.Lock()
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.mumu.Unlock()
	value, exists := m.elements[id] //查询此id是否存在
	if exists {                     //如果存在
		value.Update(nil, value.GetTimeout()) //则更新
	}
}

//通过一个Element进行更新，更新存储的信息
func (m *TimeoutMap) UpdateInfo(el Element, timeout time.Duration) {
	m.mumu.Lock()
	m.mu.RLock()
	m.mumu.Unlock()
	id := el.GetID()                //先获取发送器信息中的id
	value, exists := m.elements[id] //查询此id是否存在
	if exists {                     //如果存在
		value.Update(el, timeout) //则更新
		m.mu.RUnlock()
	} else { //否则新建
		m.mumu.Lock()
		m.mu.RUnlock()
		m.mu.Lock()
		m.mumu.Unlock()
		value := TimeoutValue.New(el, timeout)
		m.elements[id] = value
		m.mu.Unlock()
		go func() {
			el.TimeoutHandler()
			value.Run()
			m.mumu.Lock()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.mumu.Unlock()
			m.delete(id) //退出时删除
		}()
	}
}

func (m *TimeoutMap) Delete(id string) {
	m.mumu.Lock()
	m.mu.RLock()
	m.mumu.Unlock()
	value, ok := m.elements[id] //先查找
	if ok {                     //如果找得到
		m.mumu.Lock()
		m.mu.RUnlock()
		m.mu.Lock()
		m.mumu.Unlock()
		value.Stop() //则使其停止
		m.delete(id) //并删除
		m.mu.Unlock()
	} else {
		m.mu.RUnlock()
	}
}

func (m *TimeoutMap) delete(id string) {
	delete(m.elements, id)
}

func (m *TimeoutMap) getElement(id string) (Element, bool) {
	var el Element = nil
	value, ok := m.elements[id] //先查找
	if ok {                     //如果找得到
		el = value.GetElement().(Element) //则返回结果
	}
	return el, ok
}

//获取某个id对应的信息
func (m *TimeoutMap) GetElement(id string) (Element, bool) {
	m.mumu.Lock()
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.mumu.Unlock()
	return m.getElement(id)
}

//获取所有的信息
func (m *TimeoutMap) GetAll() []Element {
	m.mumu.Lock()
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.mumu.Unlock()
	var els []Element
	for id := range m.elements {
		el, ok := m.getElement(id)
		if ok {
			els = append(els, el)
		}
	}
	return els
}

func (m *TimeoutMap) Count() int {
	m.mumu.Lock()
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.mumu.Unlock()
	return len(m.elements)
}
