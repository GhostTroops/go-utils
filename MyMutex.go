package go_utils

import "sync"

var lock1 sync.Mutex

type MyMutex struct {
	*sync.Mutex
	Name string
	m    *map[string]*MyMutex
}

var (
	lockMap = map[string]*MyMutex{}
)

func GetLock(s string) *MyMutex {
	lock1.Lock()
	defer lock1.Unlock()
	if o, ok := lockMap[s]; ok {
		return o
	}
	return NewMyMutex(&lockMap, s)
}

func NewMyMutex(m1 *map[string]*MyMutex, Name string) *MyMutex {
	r := MyMutex{m: m1, Name: Name}
	r.Mutex = &sync.Mutex{}
	(*m1)[Name] = &r
	return &r
}

func (r *MyMutex) Lock() *MyMutex {
	r.Mutex.Lock()
	return r
}
func (r *MyMutex) Unlock() {
	lock1.Lock()
	defer lock1.Unlock()
	defer r.Mutex.Unlock()
	delete(*r.m, r.Name)
	r.m = nil
}
