package main

import "fmt"

func main() {
	var conferenceName = "Go Congerence"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("hello to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
