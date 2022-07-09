package main

import (
	"fmt"
	"sort"
)

type carInfo struct {
	pos int
	spd int
}

func carFleet(target int, position, speed []int) int {
	if len(position) == 1 {
		return 1
	}

	var cars []carInfo

	for i := 0 ; i < len(position) ; i++ {
		cars = append(cars, carInfo{position[i], speed[i]})
	}

	// sort by positioning
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})

	var stack []float32
	for i := len(cars) - 1 ; i >= 0 ; i-- {
		timeToDest := float32(target - cars[i].pos) / float32(cars[i].spd)
		// does new car time bigger or equal to car in stack, if so don't add
		if len(stack) != 0 && stack[len(stack) - 1] >= timeToDest {
			continue
		}
		stack = append(stack, timeToDest)
	}
	// return the stack length
	return len(stack)
}

func main() {
	p1 := []int{10,8,0,5,3}
	s1 := []int{2,4,1,1,3}
	fmt.Println(carFleet(12, p1, s1))
	p2 := []int{3}
	s2 := []int{3}
	fmt.Println(carFleet(10, p2, s2))
	p3 := []int{0,2,4}
	s3 := []int{2,4,1}
	fmt.Println(carFleet(100, p3, s3))
}