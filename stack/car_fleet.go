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

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})

	var stack []float32
	for i := len(cars) - 1 ; i >= 0 ; i-- {
		timeToDest := float32(target - cars[i].pos) / float32(cars[i].spd)
		if len(stack) != 0 && stack[len(stack) - 1] >= timeToDest {
			continue
		}
		stack = append(stack, timeToDest)
	}

	return len(stack)
}