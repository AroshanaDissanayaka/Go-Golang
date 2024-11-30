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

}