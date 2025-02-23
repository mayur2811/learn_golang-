package main

import "fmt"

func main(){
	//declar array 
	arr := []int{10,20,30} // intilize 
  // var numbers [5]int
	
  //intilize with size 
  //  var arr = [5]int{1,2,5,6}

  //array with inferred size
  arrr := [...]int{10,20,30} //compiler determine size
  
  //acces ad modification
  arrr[1] = 50
 
  fmt.Println(arrr)

	fmt.Println(arr)

	// loop  over array 
	for index , value := range arrr{
		fmt.Println(index,value)
	}
	
}