package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

const (
	BucketSize = 5
	Increment  = 2
)

func NewHMap[T1 comparable, T2 any](size int) *HMap[T1, T2] {
	return &HMap[T1, T2]{
		size:     size,
		capacity: size * BucketSize,
		buckets:  make([]*bucket[T1, T2], size),
	}
}

type HMap[T1 comparable, T2 any] struct {
	size     int
	capacity int
	buckets  []*bucket[T1, T2]
	olds     []*bucket[T1, T2]
	head     *element[T1, T2]
	tail     *element[T1, T2]
}

func (t *HMap[T1, T2]) Print() {
	b1, c1 := t.getCountElement(t.buckets, t.size)
	fmt.Println(b1, c1)
	b2, c2 := t.getCountElement(t.olds, t.size/Increment)
	fmt.Println(b2, c2)
}

func (t *HMap[T1, T2]) getCountElement(buckets []*bucket[T1, T2], size int) ([]int, int) {
	var a1 []int
	count := 0
	if buckets != nil {
		a1 = make([]int, size)
		for i, b := range buckets {
			if b == nil {
				continue
			}
			e := b.top
			for {
				if e == nil {
					break
				}
				if !e.invalid {
					a1[i]++
					count++
				}
				e = e.next
			}
		}
	}
	return a1, count
}

func (t *HMap[T1, T2]) Insert(key T1, value T2) {
	if t.capacity == 0 {
		t.olds = t.buckets
		t.size = t.size * Increment
		t.capacity = t.size * BucketSize
		t.buckets = make([]*bucket[T1, T2], t.size)
	}
	t.insert(key, value)

	if t.head.size != t.size {
		t.insert(t.head.key, t.head.value)

		var null T1
		t.head.key = null
		t.head.invalid = true
		t.head = t.head._next
	}
}

func (t *HMap[T1, T2]) insert(key T1, value T2) {
	ix := Ix(key, t.size)
	if t.buckets[ix] == nil {
		t.buckets[ix] = &bucket[T1, T2]{}
	}
	if el, ok := t.buckets[ix].Insert(key, value, t.size); ok {
		if t.head == nil {
			t.head = el
			t.tail = el
		} else {
			t.tail._next = el
			t.tail = el
		}
		t.capacity--
	}
}

func (t *HMap[T1, T2]) Lookup(key T1) (value T2, ok bool) {
	if value, ok = t.lookup(t.buckets, key); !ok {
		value, ok = t.lookup(t.olds, key)
	}
	return
}

func (t *HMap[T1, T2]) lookup(buckets []*bucket[T1, T2], key T1) (T2, bool) {
	ix := Ix(key, len(buckets))
	if buckets[ix] != nil {
		el := buckets[ix].top
		for {
			if el != nil {
				if el.key == key {
					return el.value, true
				}
				el = el.next
			} else {
				break
			}
		}
	}
	var null T2
	return null, false
}

type bucket[T1 comparable, T2 any] struct {
	top *element[T1, T2]
}

func (t *bucket[T1, T2]) Insert2(key T1, value T2, size int, e *element[T1, T2]) (*element[T1, T2], bool) {
	if e == nil {
		t.top = &element[T1, T2]{
			key: key, value: value, next: t.top, size: size,
		}
		return t.top, true
	} else if e.key == key {
		e.value = value
		return e, false
	} else {
		return t.Insert2(key, value, size, e.next)
	}
}

func (t *bucket[T1, T2]) Insert(key T1, value T2, size int) (*element[T1, T2], bool) {
	e := t.top
	for {
		if e == nil {
			t.top = &element[T1, T2]{
				key: key, value: value, next: t.top, size: size,
			}
			return t.top, true
		}
		if e.key == key {
			e.value = value
			return e, false
		}
		e = e.next
	}
}

type element[T1 any, T2 any] struct {
	key     T1
	value   T2
	next    *element[T1, T2]
	_next   *element[T1, T2]
	size    int
	invalid bool
}

func Ix[T any](key T, size int) int {
	return hash(key) % size
}

func hash(x any) int {
	switch v := x.(type) {
	case int:
		return _hash([]byte(strconv.Itoa(v)))
	case string:
		return _hash([]byte(v))
	default:
		panic("hash function of this type not implemented")
	}
}

func _hash(x []byte) int {
	h := fnv.New32a()
	_, _ = h.Write(x)
	return int(h.Sum32())
}
