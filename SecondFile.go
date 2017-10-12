package main 

import "fmt"

func main(){
	const pi float64 = 3.14159265;

	var myName string = "Derek Banas";

	fmt.Println(myName + " is a robot");

	var isOver50 bool = true;

	fmt.Printf("%.3f \n", pi); // this is how you print out a float
	// %.3 only prints the first 3 digits

	fmt.Printf("%T \n", pi); // this prints out the datatype

	//to print a boolean type,
	fmt.Printf("This prints out: %t \n", isOver50);
	// lowercase t

	fmt.Printf("%d \n",100); 
	// %d is used to print out integer data types.
}