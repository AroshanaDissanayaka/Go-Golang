package main

import "fmt"

func main() {

	//perform arithmatic operations

	a := 10

	b := 3

	sum := a + b

	difference := a - b

	quotation := a / b

	remainder := a % b

	fmt.Println("sum" , sum);
	fmt.Println("Difference" , difference);
	fmt.Println("quotation" , quotation) ;
	fmt.Println("reminder" , remainder);
	

    c:= 5;

   // Compound assignment operators

	c= c + 2;

    fmt.Println(c); // c= c + 2

	c += 10 // c= c+10

	    
    
    c++ // c= c+1;

	c--// c = c -1'


	
	// Bitwise AND
	fmt.Println("Bitwise AND:", 5 & 3) // Output: 1

	// Bitwise OR
	fmt.Println("Bitwise OR:", 5 | 3) // Output: 7

	// Bitwise XOR
	fmt.Println("Bitwise XOR:", 5 ^ 3) // Output: 6

	// Bitwise AND NOT
	fmt.Println("Bitwise AND NOT:", 5 &^ 3) // Output: 4

	// Bitwise left shift
	fmt.Println("Bitwise left shift:", 5 << 1) // Output: 10

	// Bitwise right shift
	fmt.Println("Bitwise right shift:", 5 >> 1) // Output: 2

	// Bitwise NOT
	fmt.Println("Bitwise NOT:", ^5) // Output: -6


}