package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "GO conference"
var conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	name string
	email string
	numTickets uint
}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()

	for remainingTickets != 0 {
		userName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateInput(userName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userName, userEmail, userTickets)

			wg.Add(1)
			go sendTicket(userName, userEmail, userTickets)

			fmt.Printf("\nThe bookings for %v are as follows: %v\n", conferenceName, bookings)

			if remainingTickets == 0 {
				fmt.Println("\nOur conference tickets are sold out. Please come back next year.")
				break
			}
		} else {
			fmt.Println("\nThe booking is invalid.\n")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("\nWelcome to %v ticket booking application\n", conferenceName)
	fmt.Printf("\nWe have a total of %v tickets and %v tickets are available\n", conferenceTickets, remainingTickets)
	fmt.Println("\nBook your tickets here to attend")
}

func getUserInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint

	fmt.Println("\nEnter your name:")
	fmt.Scan(&userName)
	fmt.Println("\nEnter your email:")
	fmt.Scan(&userEmail)
	fmt.Println("\nEnter the number of tickets you want to book:")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func bookTicket(userName string, userEmail string, userTickets uint) {

	var userData = userData{
		name: userName,
		email: userEmail,
		numTickets: userTickets,
	}

	bookings = append(bookings, userData)
	remainingTickets = uint(remainingTickets) - userTickets

	fmt.Printf("\nThank you %v for booking %v tickets. You will receive a ticket booking confirmation on your email %v\n", userName, userTickets,  userEmail)
	fmt.Printf("\n%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(name string, email string, numTickets uint) {
	time.Sleep(20 * time.Second)

	fmt.Println("\n#############\n")
	var ticketDetails = fmt.Sprintf("Sending %v tickets to %v on %v", numTickets, name, email)
	fmt.Println(ticketDetails)
	fmt.Println("\n#############\n")

	wg.Done()
}


