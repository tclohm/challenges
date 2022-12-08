package main

import "fmt"

func typedOut(s, t string) bool {
	var result, result2 = "", ""
	var i = len(s) - 1

	for i >= 0 {
		if s[i] == '#' {
			i--
		} else {
			result += string(s[i])
		}
		i--
	}

	i = len(t) - 1

	for i >= 0 {
		if t[i] == '#' {
			i--
		} else {
			result2 += string(t[i])
		}
		i--
	}

	return result == result2
}

func pointersTypedOut(s, t string) bool {
	var p1, p2 = len(s)-1, len(t)-1

	for p1 >= 0 || p2 >= 0 {
		if s[p1] == '#' || t[p2] == '#' {
			if s[p1] == '#' {
				var backcount = 2
				for backcount > 0 {
					p1--
					backcount--
					if s[p1] == '#' {
						backcount = 2
					}
				}
			}

			if t[p2] == '#' {
				var backcount = 2
				for backcount > 0 {
					p2--
					backcount--
					if t[p2] == '#' {
						backcount = 2
					}
				}
			}
		} else {
			if s[p1] != t[p2] { 
				return false 
			} else {
				p1--
				p2--
			}
		}
	}
	return true
}

func main() {
	a1 := "ab#z"
	a2 := "az#z"

	b1 := "az"
	b2 := "az"

	c1 := "ab#z"
	c2 := "az#z"

	fmt.Println(typedOut(a1, a2))
	fmt.Println(typedOut(b1, b2))
	fmt.Println(typedOut(c1, c2))

	fmt.Println(pointersTypedOut(a1, a2))
	fmt.Println(pointersTypedOut(b1, b2))
	fmt.Println(pointersTypedOut(c1, c2))
}