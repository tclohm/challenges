package main


import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	sc.Scan()

	const differentCharacters = 14

	for i := range sc.Text() {
		set := make(map[byte]bool)
		
		for j := 0 ; j < differentCharacters ; j++ {
			set[sc.Text()[i+j]] = true
		}

		if len(set) == differentCharacters {
			fmt.Println(i + differentCharacters)
			break
		}
	}
}