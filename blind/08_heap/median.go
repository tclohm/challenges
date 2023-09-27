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
    // make sure every numb small is <= every num in larger
    if this.large.Len() > 0 && num > this.large[0] {
        this.large.Push(num)
    } else {
        this.small.Push(-1 * num)
    }
    // uneven size?
    if this.small.Len() > this.large.Len() + 1 {
        val := -1 * this.small.Pop().(int)
        this.large.Push(val)
    }
    // uneven size? -- 2 or greater
    if this.large.Len() > this.small.Len() + 1 {
        val := this.large.Pop().(int)
        this.small.Push(-1 * val)
    }
}

// if small heap larger than large heap
// return first element from small and make it positive
// if large heap is larger than small heap
// return from large heap
// if both heaps are equal
// take the from the first of both heaps and add them together and divide by 2
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