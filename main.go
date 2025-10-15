package main

import (
	"fmt"
	"time"
)

var conferrenceName = "Go Conferrence"

const conferrenceTickets int = 50

var remainingTickets = 50

// var booking [50]string its for array
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			printFirstNames()

			if remainingTickets == 0 {
				fmt.Println("our conferrence is booked out. Come back next year ")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("your firstName or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you enetered doesn't contain @ sign")
			}
			if !isValidUserTickets {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}
}

func greetUser() {
	fmt.Printf("welcome to % v booking application \n", conferrenceName)
	fmt.Printf("We have a total of %v and %v remaining tikets available \n", conferrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of bookigs are : %v \n", firstNames)
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	//create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation letter in %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remining for %v \n ", remainingTickets, conferrenceName)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("####################################")
}
