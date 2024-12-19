package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferencename = "Go Conference"
	fmt.Println("Welcome to", conferencename, " booking application")
	const conferenceTicket = 5
	var remainingTickets uint = conferenceTicket
	booking := []string{}
	fmt.Println("We have total of", conferenceTicket, "Tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
	for {
		var username1 string
		var username2 string
		var userTicket uint
		var email string
		fmt.Println("Enter Your Username1:")
		fmt.Scan(&username1)
		fmt.Println("Enter Your Username2:")
		fmt.Scan(&username2)
		fmt.Println("Enter your Email address:")
		fmt.Scan(&email)
		fmt.Println("Enter Your Ticket:")
		fmt.Scan(&userTicket)
		remainingTickets = remainingTickets - userTicket
		//
		booking = append(booking, username1+" "+username2)
		//
		fmt.Printf("Thank you %v for booking %v tickets.You will receive a confirmation email at %v\n", username1, userTicket, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencename)
		firstnames := []string{}
		for _, booking := range booking {
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstnames)
		Isticket := remainingTickets == 0
		if Isticket {
			fmt.Println("Our conference is booked out.come back next year")
			break
		}
	}
}
