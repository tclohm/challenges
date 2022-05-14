package main

import (
	"log"
	"strconv"
	. "strings"
	"fmt"
)
// serialize and deserialize a lost of strings
// send a list of strings from server to client but cannot used json or protobuf

func encode(number int) string {
	if number > 0xFFFF {
		log.Fatal("Number to larget to encode")
	}
	return strconv.Itoa(number)
}

func decode(character string) int {
	integer, err := strconv.Atoi(character)
	if err != nil {
		log.Fatal("Corrupted")
	}
	return integer
} 

func isNextInteger(str string, index int) bool {
	numbers := "0123456789"
	for _, num := range numbers {
		if string(num) == string(str[index]) {
			return true
		}
	}
	return false
}

func serialize(list []string) string {
	var result = []string{}
	for _, word := range list {
		encoded_string := string( encode( len( word ) ) )
		result = append( result, encoded_string )
		result = append( result, word )
	}
	return Join(result, "")
} 

func deserialize(str string) []string {
	var result = []string{}

	if len(str) < 2 {
		return result
	}

	i := 0

	for i < len(str) {
		var integer string = string(str[i])

		for {
			if isNextInteger(str, i+1) {
				i++
				integer = integer + string(str[i])
			} else {
				break
			}
		}

		word_length := decode(integer)
		if word_length + i >= len(str) {
			log.Fatal("error occurred")
		}

		i++

		word := str[i:i+word_length]
		result = append( result , string(word) )
		i += word_length
	}
	return result
}

func main() {
	// works with "cat", "dog", "pig", "lamb", "frog", "zebra"
	// doesn't work if with "hippopotamus" because that has 12 letters. I would have to create sliding window
	original := []string{"cat", "dog", "pig", "lamb", "frog", "zebra", "hippopotamus"}
	fmt.Println("original:", original)
	serialized := serialize(original)
	fmt.Println("serialized:", serialized)
	deserialized := deserialize(serialized)
	fmt.Println("deserialized:", deserialized)
}