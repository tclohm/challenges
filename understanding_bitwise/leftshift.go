// ints or uint stored as 32 binary digit
// & , And operator, compares binary digit of two integers and return new integer
// wherever both numbers had a 1 and a 0 anywhere else


package main
import (
	"fmt"
)


// woah! right most bit 1 will make it odd

func main() {
	fmt.Println(37 << 3) // shift 37 to the left by 3 places
	//	  128 64 32 16  8  4  2  1
	// =============================
	// 37 : 0  0  1  0  0  1  0  1
	// -----------------------------
	// 3  : 1  0  0  1  0  1  0  0

	// ALSO : (2**3) * 37
}

