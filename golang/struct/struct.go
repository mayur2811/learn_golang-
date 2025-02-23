package main

import "fmt"

// Define a struct
type Person struct{
	Name string
	Age int
}

type Student struct {
	Name string
	Age int
}

type Car struct {
    Brand string
    Year  int
}

type Book struct {
    Title  string
    Author string
}

func main() {
    //creating a struct instance 
    person1 :=  Person{"mayur",25}

	fmt.Println(person1)

	//accesing fields
	fmt.Println(person1.Name)
	fmt.Println(person1.Age)
  
	//Struct with Named Fields
    car1 := Car{Brand: "tesla",Year:2023}
	fmt.Println(car1)
    
   //Struct with Pointers
	book1 := Book{"go","mayur"}
	changeTitle(&book1,"adbvance go ")
    fmt.Println(book1) 

	rect := Rectangle{5, 10}
    fmt.Println("Area:", rect.Area()) // Output: 50

} 
 //Struct with Pointers
func changeTitle(b *Book,newTitle string){
	b.Title = newTitle //modify original value
}


type Rectangle struct {
    Width, Height int
}
//Declaring Methods
func (r Rectangle) Area() int {
	return r.Width * r.Height
}