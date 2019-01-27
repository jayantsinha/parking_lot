package main

import "math"

type EmptySlots struct {
	set map[int]bool
}

func NewEmptySlots() *EmptySlots {
	return &EmptySlots{make(map[int]bool)}
}

func (set *EmptySlots) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found	//False if it existed already
}

func (set *EmptySlots) Contains(i int) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *EmptySlots) Remove(i int) {
	delete(set.set, i)
}

func (set *EmptySlots) Size() int {
	return len(set.set)
}

func (set *EmptySlots) GetMin() int {
	if len(set.set) == 0 {
		return -1
	}
	min := math.MaxInt32
	for k, _ := range set.set {
		if k < min {
			min = k
		}
	}
	return min
}