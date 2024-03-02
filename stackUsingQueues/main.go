package main

import "fmt"

func main() {
	x := 10
	obj := Constructor()
	obj.Push(x)
	param_2 := obj.Pop()
	fmt.Println("param2 ", param_2)
	param_3 := obj.Top()
	fmt.Println("param_3 ", param_3)
	param_4 := obj.Empty()
	fmt.Println("param_4 ", param_4)

}

type MyStack struct {
	Numbers []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.Numbers = append(this.Numbers, x)
}

func (this *MyStack) Pop() int {
	numb := this.Numbers[len(this.Numbers)-1]
	this.Numbers = this.Numbers[:len(this.Numbers)-1]
	return numb
}

func (this *MyStack) Top() int {
	if len(this.Numbers) > 0 {
		return this.Numbers[len(this.Numbers)-1]
	}
	return 0
}

func (this *MyStack) Empty() bool {
	return len(this.Numbers) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
