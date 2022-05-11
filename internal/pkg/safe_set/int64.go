package safe_set

import "sync"

type Int64 struct {
	ids  map[int64]bool
	lock *sync.RWMutex
}

func NewInt64() *Int64 {
	return &Int64{
		ids:  make(map[int64]bool),
		lock: new(sync.RWMutex),
	}
}

func (m *Int64) Has(id int64) bool {
	m.lock.RLock()
	ok := m.ids[id]
	m.lock.RUnlock()
	return ok
}

func (m *Int64) Add(id int64) {
	m.lock.Lock()
	m.ids[id] = true
	m.lock.Unlock()
}

func (m *Int64) Remove(id int64) {
	m.lock.Lock()
	m.ids[id] = false
	m.lock.Unlock()
}
