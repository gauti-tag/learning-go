package main

import "fmt"

func main() {
	conferenceName := "Go conference" // := used to declare and initialize
	const conferenceTickets int = 50  // Positive and negative numbers
	var remainingTickets uint = 50    // POsitive numbers

	/*fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your ticket here to attend")*/

	fmt.Printf("ConferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	//var userName string
	var firstName string
	var lastName string
	var email string
	var userTickets uint // positive numbers

	var bookings [50]string // Declare an array of string type with 50 reserved spaces

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) // & is called pointer used to locate the memory address

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("the whole array %v\n", bookings) // %v used to print the value
	fmt.Printf("the first element %v\n", bookings[0])
	fmt.Printf("the array type %T \n", bookings) // %T used to check the type
	fmt.Printf("the array length %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
