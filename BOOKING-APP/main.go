package main

import ("fmt" )

func main(){
	const conferenceName string = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint =50

	fmt.Printf("hello welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still availble\n",conferenceTickets,remainingTickets)
	fmt.Println("gets your conference ticket from here!")
    //book 

	var bookings [50]string 

	//declare username and ticket 
	var userName string
	var userTickets uint
	var userEmail string 

	//gets user name and number ticket emailid
	
	fmt.Print("enter your name : ")
    fmt.Scan(&userName)

    fmt.Print("enter number of tickets : ")
	fmt.Scan(&userTickets)
	
	fmt.Print("enter emailId : ")
	fmt.Scan(&userEmail)
    
	//ticket booked
	fmt.Printf("Thank you %v for bookeing %v tickets. You will receice a confirmation email at %v \n",userName,userTickets,userEmail)
    
	// remaining tickets
	remainingTickets = remainingTickets - userTickets
	bookings[0] = userName

	fmt.Println("booked person ",bookings[0])

   fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName )



}