package main

import (
	"fmt"
	"strings"
)

var conferencename = "Go Conference"

const conferenceTicket = 5

var remainingTickets uint = conferenceTicket
var booking = []string{}

func main() {
	fmt.Println("Welcome to", conferencename, " booking application")
	greatUser()
	for remainingTickets > 0 && len(booking) < 50 {

		username1, username2, email, userTicket := userinput()
		isvalid_name, isvalid_email, isvalid_ticket := validateUserInput(username1, username2, email, userTicket)
		if !isvalid_name || !isvalid_ticket || !isvalid_email {
			if !isvalid_name {
				fmt.Println("Your username is too short")
			}
			if !isvalid_email {
				fmt.Println("Your email address does not contain @ sign")
			}
			if !isvalid_ticket {
				fmt.Println("Number of tickets you entered is invalid")
			}
			break
		}
		if userTicket > remainingTickets {
			fmt.Printf("We have only %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTicket)
			break
		}
		bookTicket(userTicket, username1, username2, email)
		firstname := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstname)
	}
}

func greatUser() {
	fmt.Println("We have total of", conferenceTicket, "Tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstnames := []string{}
	for _, bookin := range booking {
		var names = strings.Fields(bookin)
		firstnames = append(firstnames, names[0])
	}
	fmt.Println(firstnames)
	return firstnames
}
func userinput() (string, string, string, uint) {
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
	return username1, username2, email, userTicket
}
func validateUserInput(username1 string, username2 string, email string, userTicket uint) (bool, bool, bool) {
	isvalid_email := strings.Contains(email, "@")
	isvalid_name := len(username1) >= 2 && len(username2) >= 2
	isvalid_ticket := userTicket > 0
	return isvalid_name, isvalid_email, isvalid_ticket
}

func bookTicket(userTicket uint, username1 string, username2 string, email string) {
	remainingTickets = remainingTickets - userTicket
	//
	booking = append(booking, username1+" "+username2)
	//
	fmt.Printf("Thank you %v for booking %v tickets.You will receive a confirmation email at %v\n", username1, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencename)
}
