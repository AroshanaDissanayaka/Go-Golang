package main;
import "fmt"

func main(){

/*fmt.Println("hello Aroshana")
fmt.Println("hello Dissanayake")


fmt.Print("aroshana");
fmt.Print("aroshana");
*/

//Golang Variables

//int

var firstNumber int = 20;
var secondNumber int =30;

//String
var userName string = "aroshana";


//float

var userScore float32 = 45.2;
var userBalance float64 = -45.2;


//boolian

var isUserFound bool = true;

//multiple variable

var firstName ,secondName string = "aroshana","Dissanayaka";


//*******easy vay to define variable in Golang


newUserName:="Aroshana Dissanayaka"

fmt.Println(newUserName)

fmt.Println(firstName,secondName);

fmt.Println(firstNumber+secondNumber);

fmt.Println(userName);

fmt.Println(userScore,userBalance);

fmt.Printf("Type:%T", userBalance);

fmt.Println(isUserFound);

//Constants

var userAge int =24;

fmt.Println(userAge);


//Calculate the area of a given circle of Radius


    var radius float32 = 5.0;

    const pi float32 = 3.14;

//equation 

var area float32 = 0.0;

area = pi * radius * radius

fmt.Println(area);



//Type Coversions


var intNum  int = 10;
var floatNum float64 = 3.5;

sum := intNum + int(floatNum)

fmt.Println(sum);

sum2 := float64(intNum) + floatNum

fmt.Println(sum2);



}