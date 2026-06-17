package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestNewHMapInsert(t *testing.T) {
	m := NewHMap[int, int](2)
	for i := 1; i <= 10000000; i++ {
		m.Insert(i, i)
		m.Print()
	}
	return
	for i := 1; i <= 10000000; i++ {
		_, _ = m.Lookup(i)
	}
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10000000; i++ {
		m[i] = i
	}
	return
	for i := 0; i < 10000000; i++ {
		_, _ = m[i]
	}
}

type Element[T1 any] struct {
	key   T1
	value T1
}

func _type(a any, t reflect.Type) {
	fmt.Println(reflect.TypeOf(a), t)
}

func TestMemo2(t *testing.T) {
	fmt.Println(unsafe.Sizeof(Element[int]{}))

	key := 1010101010
	value := 2020202020

	e := Element[int]{
		key:   key,
		value: value,
	}
	fmt.Println(e)

	fmt.Println(len(make([]Element[int], 100)))

	b := *(*[unsafe.Sizeof(e)]byte)(unsafe.Pointer(&e))
	fmt.Println(b)

	e2 := (*Element[int])(unsafe.Pointer(&b))
	fmt.Println(e2)

	fmt.Println(e2.key, e2.value)

	buff := make([]byte, 32)

	copy(buff, b[:])

	fmt.Println(buff)
}

func TestMem(t *testing.T) {
	/*b := element{
		key:   10,
		value: 10,
	}
	p := unsafe.Pointer(&b)
	m := *(*element)(p)

	fmt.Println(m)*/
}
