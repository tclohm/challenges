package main

import "fmt"

type MinStack struct {
    container []int
}


func Constructor() MinStack {
    ms := MinStack{
        container: make([]int, 0),
    }
    return ms
}


func (this *MinStack) Push(val int)  {
    this.container = append(this.container, val)
}


func (this *MinStack) Pop()  {
    end := len(this.container) - 1
    this.container = this.container[:end]
}


func (this *MinStack) Top() int {
    top := len(this.container) - 1
    return this.container[top]
}


func (this *MinStack) GetMin() int {
    min := 10000000
    for _, i := range this.container {
        if i < min {
            min = i
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
 func main() {
    obj := Constructor()
    fmt.Println(obj)
    obj.Push(8)
    obj.Push(-1)
    obj.Push(4)
    obj.Push(27)
    fmt.Println(obj)
    obj.Pop()
    fmt.Println(obj)
    fmt.Println(obj.Top())
    fmt.Println(obj.GetMin())
    obj.Pop()
    obj.Pop()
    fmt.Println(obj.Top())
 }