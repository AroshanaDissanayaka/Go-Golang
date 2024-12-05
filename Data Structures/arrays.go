package main

import "fmt"

func main() {

  
  //define array Golang

	var array = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array[1]);
   

  //assign Elements to arry

	var colors [3]string 

	colors[0] = "red";

	fmt.Println(len(colors));


    //


	var numbers = []int {1,2,3,4,5};

	fmt.Println(numbers)


	//2 Dimentional array (Matrix , 2D Array)



     var twoDArray = [3][3]int {{2,4,6},{8,10,5},{11,14,17}}

	 fmt.Println(twoDArray)

       
	 fmt.Println(twoDArray[0][2])




}