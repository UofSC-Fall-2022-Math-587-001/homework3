package main 

import (
	"fmt"
	"os" 
	"bufio" 
	"strings"
	"strconv"
)

func gcd(a, b int) [3]int {
	// Implement the following algorithm from Exercise 1.12 of the text 
	// compute g = gcd(a,b) and u,v such that au + bv = g
	// 1. Set u = 1, g = a, x = 0, and y = b
	// 2. If y = 0, set v = (g-au)/b and return (g,u,v) 
	// 3. Divide g by y with remainder, g = qy+t, with 0 <= t < y 
	// 4. Set s = u - qx
	// 5. Set u = x and g = y
	// 6. Set x = s and y = t
	// 7. Go to Step (2)
	var g, u, v int 
	return [3]int{g,u,v} 
}
// Additional questions:
// - What happens when b = 0? Can you fix this? 
// - Modify the program to return u > 0 and as small as possible. 
// Hint: if (u,v) is a solution so is (u+b/g,v-a/g). 

// This function builds in the user interaction in homework 1. We 
// stripped out of main into its own function since we call it 
// more than once. We return both an integer and any errors 
// along the way. 
func getInt(s string) (int, error) {
	fmt.Printf("Enter the %s integer: ",s)
	reader := bufio.NewReader(os.Stdin) 
	input, err := reader.ReadString('\n') 
	if err != nil {
		return 0, err 
	}
	input = strings.TrimSuffix(input, "\n")
	// we use the standard library's function to try to 
	// iterpret a string as an integer 
	out, err2 := strconv.ParseInt(input, 10, 0) 
	if err2 != nil {
		return 0, err2 
	}
	// According to the documentation, out should be a int 
	// but it seems to be an int64. We cast it to int and 
	// cross our fingers. 
	return int(out), nil 
}

func main() {
	int1, err1 := getInt("first")
	if err1 != nil {
		fmt.Println("Something went wrong. Please try again. This might help: ", err1)
		return 
	}

	int2, err2 := getInt("second")
	if err2 != nil {
		fmt.Println("Something went wrong. Please try again. This might help: ", err2)
		return
	}

	g := gcd(int1,int2)[0]
	u := gcd(int1,int2)[1]
	v := gcd(int1,int2)[2]

	fmt.Printf("The gcd of %d and %d is %d\n", int1, int2, g)
	fmt.Printf("We have %d * %d + %d * %d = %d\n", int1, u, int2, v, g)
} 
