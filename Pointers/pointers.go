package main

import "fmt"

//pointer is use store address in variable

func main(){

	var num int = 10;

	fmt.Println(num);


	 var ptr *int 

	ptr = &num

	fmt.Println(ptr);


	//Dereference ptr to get the value it points to


	fmt.Println(*ptr)


	*ptr = 200;

    //assign value to variable using Pointer
	
	fmt.Println(num)

}