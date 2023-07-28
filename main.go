package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Package level variables : available for the functions and main func
const confrenceTickets uint = 50

// Package variable declaration
var conferenceName = "Go Confrence"
var remainingTickets uint = 50

// Creating empty list of maps/ slice of maps: var bookings = make([]map[string]string, 0)
var bookings = make([]userData, 0) // of type struct

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOFTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// defining a slice array not assigning any fixed no inside []

	greetUsers()

	// Call user input function
	firstName, lastName, email, userTickets := getUserInput()

	// Call validation function
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		// Call booking logic function
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Call function print first names
		firstNames := printFirstNames() // returned value is saved in value to print
		fmt.Printf("All the bookings with first names are here: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Tickets sold out. Check back next year. ")
			//break
		}
	} else {
		if !isValidName {
			fmt.Printf("First name or last name you entered is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Invalid email entered\n")
		}
		if !isValidTicketNumber {
			fmt.Printf("Number of tickets you entered is invalid! \n")
		}
	}
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func printFirstNames() []string {
	// Just print array of 1st name of person who booked ticket
	// Using Slicing inside the slicing
	firstNames := []string{} // fistNames is the local variavle for printFirstName function
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// I/P from user for input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create an empty map for a user (use make for it ) :var userData = make(map[string]string)
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOFTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n\n ", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(7 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("\n Sending ticket: %v \n to email address %v\n\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
