package main

import (
	"fmt"
	"strconv"
)

func restore(ipAddress string) []string {
	var list []string

	var backtrack func (currentIP string, index int, dots int)
	backtrack = func (currentIP string, index int, dots int) {
		if dots == 4 && index == len(ipAddress) {
			list = append(list, currentIP[:len(currentIP) - 1])
			return
		} 

		if dots > 4 { return }
		minimum := min(index + 3, len(ipAddress))

		for i := index ; i < minimum ; i++ {

			var integer, err = strconv.Atoi(ipAddress[index:i+1])
			
			if err != nil { return }

			if integer < 256 && (index == i || ipAddress[index] != '0') {
				
				backtrack(currentIP + string(ipAddress[index:i + 1]) + ".", i + 1, dots + 1)

			}
		}
	}
	backtrack("", 0, 0)

	return list
}

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func main() {
	fmt.Println(restore("25525511135"))
	fmt.Println(restore("0000"))
	fmt.Println(restore("101023"))
}