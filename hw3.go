package main 

import (
	"fmt"
	"os" 
	"bufio" 
	"strings"
	"strconv"
)

func fastPower(N, g, A int) int {
	// Implement the fast-powering algorithm from Exercise 1.25 of the text 
	// Input: Positive integers N, g, and A
	// 1. Set a = g and b = 1.
	// 2. Loop while A > 0. 
	// 	3. If A = 1 (mod 2), set b = ba (mod N). 
	//	4. Set a = a^2 (mod N) and A = largest integer <= A/2 
	//	5. If A > 0, continue with loop at Step 2. 
	// 6. Return the number b, which equals g^A mod N. 
	var b int 
	return b 
}

func primalityTest(N int) bool {
	// This function takes in an integer and returns true if it is prime 
	// false if not. How fast can you make it go?
	var b bool 
	return b 
}

func primRootsModN(N int) []int {
	// Recall that g in Z/NZ is primitive root if every nonzero element of 
	// can be written as g^k for some integer k. This function should 
	// take in N and return all primitive roots in Z/NZ. To simplify, 
	// you can assume that we always feed in prime N. How fast can you make 
	// go?
	var l []int 
	return l
}

func getInt(s string) (int, error) {
	fmt.Println("Enter your integer: ")
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

	roots := primRootsModN(int1)
	fmt.Printf("The primitive roots modulo %d are %d\n",int1,roots)

} 
