package collections

import "container/heap"

type PriorityQueue[T any] interface {
	Len() int
	Add(item T)
	Poll() T
	Empty() bool
	Contains(item T) bool
	Update()
}

func NewPriorityQueue[T comparable](cmp func(a, b T) bool) PriorityQueue[T] {
	queue := &priorityQueue[T]{cmp, make([]T, 0), make(map[T]bool)}
	heap.Init(queue)
	return queue
}

func (p *priorityQueue[T]) Update() {
	heap.Init(p)
}

func (p *priorityQueue[T]) Add(item T) {
	heap.Push(p, item)
	p.exists[item] = true
}

func (p *priorityQueue[T]) Poll() T {
	e := heap.Pop(p).(T)
	delete(p.exists, e)
	return e
}

func (p *priorityQueue[T]) Empty() bool {
	return p.Len() == 0
}

func (p *priorityQueue[T]) Contains(item T) bool {
	return p.exists[item]
}

type priorityQueue[T comparable] struct {
	cmp    func(a, b T) bool
	elems  []T
	exists map[T]bool
}

func (p *priorityQueue[T]) Len() int {
	return len(p.elems)
}

func (p *priorityQueue[T]) Less(i, j int) bool {
	return p.cmp(p.elems[i], p.elems[j])
}

func (p *priorityQueue[T]) Swap(i, j int) {
	tmp := p.elems[i]
	p.elems[i] = p.elems[j]
	p.elems[j] = tmp
}

func (p *priorityQueue[T]) Push(x any) {
	p.elems = append(p.elems, x.(T))
}

func (p *priorityQueue[T]) Pop() any {
	e := p.elems[len(p.elems)-1]
	p.elems = p.elems[0 : len(p.elems)-1]
	return e
}
