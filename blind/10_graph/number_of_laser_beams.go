package main

import "fmt"

type Bank struct {
	layout []string
}

func Constructor(bankLayout []string) Bank {
	return Bank{
		layout: bankLayout,
	}
}

func (b Bank) countLasers() int {
	return len(b.layout)
}

func main() {
	layout := []string{"011001", "000000", "010100", "001000"}
	bank := Constructor(layout)
	fmt.Println(bank)
	fmt.Println(bank.countLasers())
}