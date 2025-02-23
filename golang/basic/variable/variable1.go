package main

import("fmt")

func main(){
	fmt.Print("hi");

	//variable
	var name string  = "mayur"
	var age int  = 23
	var price float64 = 99.99 
	var isGoAwesome bool =  true

	//dynamic allocation
	id := 2222
	college := "ait"

	//consttant 
	const PI = 3.14 

	//multiple intlizetion
	var x,y,z int = 2,5,6

	var str string   // Default: ""
    var num int      // Default: 0
    var isValid bool // Default: false

    fmt.Println(str, num, isValid)

	fmt.Println(x, y, z)

	fmt.Print(id,college)
	fmt.Println("Value of PI:", PI)

	fmt.Println(name, age, price, isGoAwesome,)


	// int -> %d 
	// string -> %s
	//float -> %f
	//tyep -> %T
}
