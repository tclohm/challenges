// logical order of items inside -- Dijkstra shortest path
// message queues or task queues
// binary heap
// min heap, smallest key is always at the front
// max heap, largest key always at front

package main

type BinaryHeap struct {
	items []int
}

func (bh *BinaryHeap) PercolateUp() {
	i := len(bh.items)

	for i / 2 > 0 {
		if bh.items[i] < bh.items[i / 2] {
			bh.items[i / 2], bh.items[i] = bh.items[i], bh.items[i / 2]
		}
		i = i / 2
	}
}

func (bh *BinaryHeap) PercolateDown(i int) {
	for i * 2 < len(bh.items) {
		mc := bh.MinChild(i)
		if bh.items[i] > bh.items[mc] {
			bh.items[i], bh.items[mc] = bh.items[mc], bh.items[i]
		}
		i = mc
	}
}

func (bh *BinaryHeap) MinChild(i int) int {
	if i * 2 + 1 > len(bh.items) {
		return i * 2
	}

	if bh.items[i * 2] < bh.items[i * 2 + 1] {
		return i * 2
	}

	return i * 2 + 1
}

func (bh *BinaryHeap) DeleteMin() int {
	return_value := bh.items[1]
	bh.items[1] = bh.items[len(bh.items)]
	bh.items = bh.items[:len(bh.items) - 1]
	bh.PercolateDown(1)
	return return_value
}

func (bh *BinaryHeap) Insert(k int) {
	bh.items = append(bh.items, k)
	bh.PercolateUp()
}

func (bh *BinaryHeap) BuildHeap(list []int) {
	i := len(list) / 2
	start := []int{0}
	bh.items = append(start, list...)
	for i > 0 {
		bh.PercolateDown(i)
		i -= 1
	}
}

func main() {

}