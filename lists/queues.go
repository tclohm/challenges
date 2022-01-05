package main

import (
	"fmt"
)

type Queue []*Order

type Order struct {
	priority int
	quantity int
	product string
	customer string
}

func (order *Order) New(priority int, quantity int, product string, customer string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customer = customer
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "CPU"
	var customerName1 string = "Jorge Sans"
	order1.New(priority1, quantity1, product1, customerName1)

	var order2 *Order = &Order{}
	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "RAM"
	var customerName2 string = "Lana Vella"
	order2.New(priority2, quantity2, product2, customerName2)

	queue.Add(order1)
	queue.Add(order2)

	for i := 0 ; i < len(queue) ; i++ {
		fmt.Println(queue[i])
	}
}