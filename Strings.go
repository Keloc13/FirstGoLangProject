package main 

import ("fmt" 
"strings"
"sort"
"os"
"log"
"io/ioutil")

func main(){
	samString := "Hello world";

	fmt.Println(strings.Contains(samString, "lo"))
	fmt.Println(strings.Index(samString, "lo"))
	fmt.Println(strings.Count(samString, "l"))
	fmt.Println("This is a test for pushing");
	testString1();
	createFile();
}

func testString1(){

	csvString := "1,2,3,4,5,6";

	fmt.Println(strings.Split(csvString,","));

	listOfLetters := []string{"c","a","b"};

	sort.Strings(listOfLetters);

	fmt.Println("Letters: ", listOfLetters);
}

func createFile(){

	if _, er := os.Stat("sample.txt");
	os.IsNotExist(er){
		file, err := os.Create("samp.txt");

		if err != nil{
		log.Fatal(err);
	}

	file.WriteString("This is some random text");
	file.Close();
	}
	
	stream, err := ioutil.ReadFile("samp.txt");

	if err != nil{
		log.Fatal(err);
	}

	readString := string(stream);
	fmt.Println(readString);
}
