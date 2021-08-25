// the ~, bitwise NOT operator, it's the negative

package main
import (
	"fmt"
)



func main() {
	fmt.Println(37 + -37) // equal 
	//	  128 64 32 16  8  4  2  1
	// =============================
	// 37 : 0  0  1  0  0  1  0  1
	// -----------------------------
	// ~37: 1  1  0  1  1  0  1  0
}

