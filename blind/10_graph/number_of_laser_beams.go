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

func (b Bank) count(row int) int {
	total := 0
	for col := 0 ; col < len(b.layout[row]) ; col++ {
		// byte
		if b.layout[row][col] == '1' {
			total += 1
		}
	}
	//fmt.Println("ROW:", row, "total count:", total)
	return total
}

func (b Bank) countLasers() int {

	rows := len(b.layout)

	prev := b.count(0)

	total := 0

	for row := 1 ; row < rows ; row++ {
		curr := b.count(row)
		if curr > 0 {
			total += (prev * curr)
			prev = curr
		}
		
	}

	return total
}

func main() {
	layout := []string{"011001", "000000", "010100", "001000"}
	bank := Constructor(layout)
	fmt.Println(bank.countLasers())
}