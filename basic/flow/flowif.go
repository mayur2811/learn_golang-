package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("welcome ")

	if 10>20 {
		fmt.Println("mayur")
	}else{
		fmt.Println("zala")
	}

	x := 10 
	if x < 10 {
		fmt.Println("x is greater")
	}else if x == 10 {
		fmt.Println("x is lesser")
	}else{
		fmt.Println("xxxx ")
	}
   
	//switch 
	switch 20 {
	case 10 :
		fmt.Println("value is 10")
	case 20:
		fmt.Println("value is 20")
	default:
		fmt.Println("no value in it")
	}
	
	num := 20 
	switch {
	case num <10 : 
	fmt.Println("num is less then 10")
	case num >10 && num < 30:
		fmt.Println("num is between 10 -30")
	default:
		fmt.Println("num is  30 or greater ")
	}

	 
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func ()  {
		time.Sleep(10*time.Second)
		ch1 <- "message from ch1"
	}()

	go func ()  {
		time.Sleep(9*time.Second)
		ch2 <- "message from ch2"
	}()

  	select{
 	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <- ch2 :
		fmt.Println(msg2)
	}

	i := 0
	for i < 5 {
		fmt.Println("value of  i: ",i)
		i++
	}

	for i := 0; i < 5;i++{
		fmt.Println("i: ",i)
	}

	numbers := []int{10,20,30,40}
	for index,value := range numbers{
		fmt.Println("index :%d, value :%d\n",index,value)
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}
    for key, value := range fruits {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }

	word := "golang"
	for index,char := range word{
		fmt.Println("index :%d , character :%c\n",index,char)
	}

}