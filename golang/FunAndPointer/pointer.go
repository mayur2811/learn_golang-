package main

import "fmt"

func main() {
    var num int = 10
	var ptr *int = &num

	fmt.Println("value:",num)
	fmt.Println("address:",ptr) // memory address

	//Dereferencing a Pointer
	 //*ptr accesses the value stored at the memory address.
	 fmt.Println("Value using pointer:",*ptr)

	 //Modifying Values Using Pointers
     value := 10
	 updateValue(&value)//passing address
	 fmt.Println("updated value:",value)//  anserr Output: Updated value: 50
     

	 //Pointers with Arrays and Slices
	 numbers := []int{1, 2, 3}
      modifyArray(&numbers)
     fmt.Println(numbers) // Output: [100 2 3]


	 //Assigning Functions to Variables
     var f func(string) = greet
	 f("alice")

	 //Passing Functions as Arguments
	 result := operation(5, 3, add)
     fmt.Println("Result:", result) // Output: Result: 8

	 //returning function 
	 double := mul(2)
	 fmt.Println(double(5)) //output :10

	}

func updateValue(num *int){
	*num = 50 //changin value at memory address
}

//Pointers with Arrays and Slices
func modifyArray(arr *[]int){
	(*arr)[0] = 100
}

func greet(name string) {
    fmt.Println("Hello,", name)
}

//Passing Functions as Arguments
func operation(a,b int ,op func(int,int)int)int{
	return op(a,b)
}

func add(x, y int) int {
    return x + y
}


//Returning Functions
func mul(factor int) func(int)int{
	return func(num int) int{
		return num*factor
	}
}