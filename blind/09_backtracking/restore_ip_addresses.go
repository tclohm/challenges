package main

import "fmt"

func restore(ipAddress string) []string {
	// four integers - seperated by a single dot
	// each integer between [0-255]
	// no leading zeros
	return []string{}
}

func main() {
	fmt.Println(restore("25525511135"))
	fmt.Println(restore("0000"))
	fmt.Println(restore("101023"))
}