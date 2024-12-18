package main

import "fmt"

func main() {
	var conferencename = "Go Conference"
	fmt.Println("Welcome to", conferencename, " booking application")
	const conferenceTicket = 50
	var remainingTickets uint = conferenceTicket
	fmt.Println("We have total of", conferenceTicket, "Tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
	for {
		var username string
		var userTicket uint
		var email string
		fmt.Println("Enter Your Username:")
		fmt.Scan(&username)
		fmt.Println("Enter your Email address:")
		fmt.Scan(&email)
		fmt.Println("Enter Your Ticket:")
		fmt.Scan(&userTicket)
		remainingTickets = remainingTickets - userTicket

		//

		//

		//

		fmt.Printf("Thank you %v for booking %v tickets.You will receive a confirmation email at %v\n", username, userTicket, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencename)
	}
}
