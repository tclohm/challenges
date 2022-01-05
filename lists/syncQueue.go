package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass 	int
	waitTicket 	int
	playPass 	bool
	playTicket 	bool
	queuePass 	chan int
	queueTicket chan int
	message 	chan int
}

func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)


	go func() {
		var message int
		for {
			select {
			case message = <- queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}

				if queue.waitPass > 0 && 
				   queue.waitTicket > 0 && 
				   !queue.playPass && 
				   !queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket--
					queue.waitPass--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

func (Queue *Queue) StartTicketIssue() {
	Queue.message <- messageTicketStart
	<-Queue.queueTicket
}

func (Queue *Queue) EndTicketIssue() {
	Queue.message <- messageTicketEnd
}

func ticketIssue(Queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		Queue.EndTicketIssue()
	}
}

func (Queue *Queue) StartPass() {
	Queue.message <- messagePassStart
	<- Queue.queuePass
}

func (Queue *Queue) EndPass() {
	Queue.message <- messagePassEnd
}

func passenger(Queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println(" Passenger starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		Queue.EndPass()
	}
}

func main() {
	var Queue *Queue = &Queue{}
	Queue.New()
	fmt.Println(Queue)
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i, " passenger in the Queue")
		go passenger(Queue)
	}

	for j := 0 ; j < 5 ; j++ {
		go ticketIssue(Queue)
	}

	select{}
}