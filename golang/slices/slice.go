package main

import "fmt"

func main() {
   //Slices expand dynamically (unlike arrays which have a fixed size). They are more commonly used in Go.
   var nums []int //declar slice
   fmt.Println(nums)

	num := []int{10,20,30,40}

	fmt.Println(num) // Output: []

	//Appending Elements to a Slice
	nums = append(nums,4,5)// Add elements dynamically
	fmt.Println(nums)

    //Slicing a Slice
	part1 := nums[:3]
	part2 := nums[2:]


    fmt.Println(part1) 
    fmt.Println(part2) 

	//Using make() to Create a Slice
	number := make([]int,5) //create slice of size 5
	fmt.Println(number)  // Output: [0 0 0 0 0]

    //Copying Slices
	src := []int{1,23,3,4}
	dst := make([]int,len(src))

	copy(dst,src)//copy element
	fmt.Println(dst)//// Output: [1 23 3 4]
}
