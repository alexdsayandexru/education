package easy

import (
	"fmt"
	"testing"
)

func Test_225(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	buffer []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.buffer = append(this.buffer, x)
}

func (this *MyStack) Pop() int {
	val := this.Top()
	this.buffer = append([]int{}, this.buffer[:len(this.buffer)-1]...)
	return val
}

func (this *MyStack) Top() int {
	return this.buffer[len(this.buffer)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.buffer) == 0
}
