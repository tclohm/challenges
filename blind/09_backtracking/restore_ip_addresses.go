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

// other way
func restoreIpAddresses(s string) []string { 
    res := []string{}

    for a := 1; a <= 3; a++ {
        for b := 1; b <= 3; b++ {
            for c := 1; c <= 3; c++ {

                d := len(s) - (a + b + c)
                if d > 0 && d <= 3 && a+b+c+d == len(s) {
                    A, _ := strconv.Atoi(s[:a])
                    B, _ := strconv.Atoi(s[a:a+b])
                    C, _ := strconv.Atoi(s[a+b:a+b+c])
                    D, _ := strconv.Atoi(s[a+b+c:])

                    if A <= 255 && B <= 255 && C <= 255 && D <= 255 {
                        addr := fmt.Sprintf("%d.%d.%d.%d", A, B, C, D)
                        
                        // To check if there was any leading zeros, 
                        // If leading zeros was there, then len will be reduced
                        // +3, for 3 '.'
                        if len(addr) == len(s) + 3 {
                            res = append(res, addr)
                        }
                    } 
                }
            }
        }
    }

    return res
}