package main

import "fmt"

type data struct {
	value string
	timestamp int
}

type TimeMap struct {
	container map[string][]data
}

func Constructor() TimeMap {
	return TimeMap{
		container: make(map[string][]data),
	}
}

func (this *TimeMap) Set(key, value string, timestamp int) {
	this.container[key] = append(this.container[key], data{
		value: value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.container[key]) == 0 || this.container[key][0].timestamp > timestamp {
		return ""
	}

	values := this.container[key]
	left, right := 0, len(values) - 1
	for left <= right {
		middle := left + (right - left) / 2
		if values[middle].timestamp > timestamp {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return values[right].value
}

func main() {
	cache := Constructor()
	cache.Set("hello", "Taylor", 1)
	cache.Set("hello", "Parker", 2)
	cache.Set("hello", "Marta", 3)
	returned := cache.Get("hello", 1)
	returned_again := cache.Get("hello", 2)
	returned_third := cache.Get("hello", 3)
	fmt.Println(returned, returned_again, returned_third)
}