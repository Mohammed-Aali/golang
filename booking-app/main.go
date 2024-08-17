package main

import "fmt"

func main() {
	// only works with variables
	conferenceName := "Go Congerence"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("hello to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings : %v \n", bookings)
	}

	// no syntatic sugar
	// var booking = [50]string{}
	// syntatic sugar

}
