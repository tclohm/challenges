// ints or uint stored as 32 binary digit
// | , Or operator, Either int has a 1, will return it


package main
import (
	"fmt"
)


// woah! right most bit 1 will make it odd

func main() {
	fmt.Println(37 | 23) // equal 55
	//	  128 64 32 16  8  4  2  1
	// =============================
	// 37 : 0  0  1  0  0  1  0  1
	// 23 : 0  0  0  1  0  1  1  1
	// -----------------------------
	// 55 : 0  0  1  1  0  1  1  1
}

