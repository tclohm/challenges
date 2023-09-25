package main

import "fmt"

type IntHeap []int
func (this IntHeap)  Len()          	int  		{ return len(this) }
func (this IntHeap)  Less(i, j int) 	bool 		{ return this[i] < this[j] }
func (this IntHeap)  Swap(i, j int) 	 	 		{ this[i], this[j] = this[j], this[i] }
func (this *IntHeap) Push(x interface{}) 	        { *this = append(*this, x.(int)) }
func (this *IntHeap) Pop() 				interface{} {
	hp := *this
    n := len(hp)
    el := hp[n - 1]
	*this = hp[0 : n - 1]
	return el
}

type MedianFinder struct {
    small IntHeap
    large IntHeap
}


func Constructor() MedianFinder {
    return MedianFinder{small: IntHeap{}, large: IntHeap{}}
}


func (this *MedianFinder) AddNum(num int)  {
    if len(this.large) > 0 && num > this.large[0] {
        this.large.Push(num)
    } else {
        this.small.Push(-1 * num)
    }

    if this.small.Len() > this.large.Len() + 1 {
        val := -1 * this.small.Pop().(int)
        this.large.Push(val)
    }

    if this.large.Len() > this.small.Len() + 1 {
        val := -1 * this.large.Pop().(int)
        this.large.Push(-1 * val)
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() > this.large.Len() {
        return float64(-1 * this.small[0])
    } else if this.large.Len() > this.small.Len() {
        return float64(this.large[0])
    }
    return float64(-1 * this.small[0] + this.large[0]) / 2
}

func main() {
    mf := Constructor()
    mf.AddNum(1)
    mf.AddNum(2)
    fmt.Println(mf.FindMedian())
    mf.AddNum(3)
    fmt.Println(mf.FindMedian())
}