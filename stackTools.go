package go_utils

import (
	"fmt"
	"sync"
)

// 创建一个泛型的先进先出（FIFO）栈
type FIFOStack[T any] struct {
	data    []T
	lk      *sync.Mutex
	PushHd  bool
	PopTail bool
}

func NewFIFOStack[T any]() *FIFOStack[T] {
	return &FIFOStack[T]{lk: &sync.Mutex{}}
}
func (s *FIFOStack[T]) SetPush2Head() {
	s.PushHd = true
}
func (s *FIFOStack[T]) SetPopTail() {
	s.PopTail = true
}

// 入栈操作
func (s *FIFOStack[T]) Push(value ...T) {
	s.lk.Lock()
	defer s.lk.Unlock()
	if s.PushHd {
		s.data = append(value, s.data...)
	} else {
		s.data = append(s.data, value...)
	}
}
func (s *FIFOStack[T]) Len() int {
	return len(s.data)
}

// 出栈操作
func (s *FIFOStack[T]) Pop() (T, error) {
	s.lk.Lock()
	defer s.lk.Unlock()
	var r T
	if len(s.data) == 0 {
		return r, fmt.Errorf("栈为空")
	}
	var value T
	if s.PopTail {
		n4 := len(s.data) - 1
		value = s.data[n4]
		s.data = s.data[:n4]
	} else {
		value = s.data[0]
		s.data = s.data[1:]
	}
	return value, nil
}
