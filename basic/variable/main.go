package main

import "fmt"

func main() {

	//var variableName dataType
	
	// var empName string = "mayursinh" //explicitly assign value and type
	// empName = "zala"
	// var salary int = 1000000 
	// var workHour float32 = 8.5
	// const gender string = "male"
	//var inOffice bool = true

    
	//fmt.Print("emp_Name := ",empName)

    //short declaration 
	// age := 25  // implicitly declares and assign a int 
    // empName := "rakesh"
	// salry := 100000
	// workHour := 8.5
	// gender := "male"
	// inOffice = true

	//fmt.Println("emp details are ",empName)


	//multiple variable 
	var x,y int = 10, 20
	fmt.Println("x:",x,"y:",y)

	a ,b := "hello " , 2.5
	fmt.Println(a," ",b)

	 birthYear ,joinYear := 2000 , 2015
	 fmt.Print("birthYear: ",birthYear ," " , "joinYear: ",joinYear)


	/**  data types  
	Basic Types:
	int, float64,32, string, bool
    Composite Types:
     array, slice, map, struct  **/


	 //composite 

	 //array  var arrayName [size]dataType
      var empList [2]string  // declare and array of 200 string
	  empList[0] ,empList[1]= "mayursinh", "kuldip"

	  //short 
	  //studentList := [2]string{"mayursinh","kuldip"}

	  //slice   sliceName := []dataType
      number := []int{1,2,23,5}
	  number = append(number, 56)

	  slice := number[1:3] // from index 1 to 3 

	  fmt.Println("Slice:", slice)



      //map  A map is a collection of key-value pairs, where each key is unique
	  //Keys must be of a comparable type (e.g., int, string).vlues can be of any type.
      
	  //mapName := make(map[keyType]valueType)
      
	  List := make(map[int]string)
		List[01] = "raju"
		List[02]  = "mahesh"

		fmt.Println("list is ",List[01])

	  capitals := map[string]string{
		 "india" : "new delhi" ,
		 "france" : "paris" ,
		 "germany" : "berlin",
	  }
        fmt.Println("capitals:",capitals)

		//add or update an element
       capitals["japan"] ="tokyo"
	   fmt.Println("Updated Capitals:", capitals)
	  
	   fmt.Println("india capital is ",capitals["india"])

	  	//check existence
		capital , exits := capitals["usa"]
	  	if exits {
			fmt.Println("Capital of USA:", capital)
		}
       
		//delete an element
		delete(capitals,"france")


		//struct 
		//A struct is a composite data type that groups fields under a single name. Fields can have different types.
		  
		Person1 :=  Person{
			Name : "alice",
			Age : 25 ,
			City : "rajkot",
		}

		Person2 := Person{
			Name : "mayursinh",
			Age : 23,
			City : "sure..",
		}
		
		fmt.Println("person 1:",Person1)
		
		Person2.Age = 32
		fmt.Println("person 2 age :", Person2.Age)

		//anonymous struct 
		person3 := struct {
			Name string
			Hobby string
		}{
			Name:  "Bob",
			Hobby: "Reading",
		}
		fmt.Println("Person 2:", person3)

	    


} 

/** syntax
		type StructName struct {
		field1 dataType
		field2 dataType
			} **/

	type Person struct{
		Name string
		Age int
		City string
	}