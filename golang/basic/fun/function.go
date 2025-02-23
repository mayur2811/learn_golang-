package main

import "fmt"

func main(){
     greet()
     greetUser("mayur")
     add(5,6)
     multi(6,60)
     result := sub(10,5)
     fmt.Println("Result:", result)
     x,y := swap(3,7)
     fmt.Println("Swapped values:", x, y)

     q,r := divide(10,3)
     fmt.Println("Quotient:", q, "Remainder:", r)

     sum(1,2,35,6,8)

     resu := applyOperation(4,5,multiply)
     fmt.Println(resu)

     defe();
}

func greet(){
     fmt.Println("hello go!")
}

func greetUser(name string){
     fmt.Println("hello,",name)
}

func add(a int ,b int){
     fmt.Println(" sum:",a+b)
}

func multi(a,b int ){
     fmt.Println("multi ",a*b)
}

//with return 
func sub(a,b int ) int {
     return a-b  
}

func swap(a,b int) (int,int){
     return b,a 
}

//named return 
func divide(a,b int) (quotient int , remainder int){
     quotient = a/b 
     remainder = a%b
     return
}

//Variadic Functions (Multiple Arguments)
//A variadic function takes an unknown number of arguments:
func sum(numbers ...int){
     total :=0
     for _,num:= range numbers{
          total += num
     }
     fmt.Println("sum:",total)
}


//function as parameter 
func applyOperation(a,b int,operation func(int,int)int)int{
     return operation(a,b)

}

func multiply(x,y int) int{
     return x*y
}

//Anonymous Functions (Lambdas)
//You can define a function without a name inside another function:
func ann(){
     //annoymous
     add := func(a,b int)int{
     return a+b
} 

fmt.Println(add(3,7))
}

//defer postpones execution of a function until the surrounding function returns.
func defe(){
     fmt.Println("Start")
    // Defer executes in LIFO (Last In, First Out) order.
     defer fmt.Println("Deferred Call 1")
     defer fmt.Println("Deferred Call 2")
 
     fmt.Println("End")
}