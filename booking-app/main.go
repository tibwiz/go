package main

import (
	"booking-app/shared"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var availableTickets = 50
var bookingData = make([]UserData, 0)

type UserData struct {
	name            string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	shared.GreetUser(conferenceName, availableTickets)
	for {
		var userName string
		var email string
		var numberOfTickets int

		fmt.Print("Enter your name")
		fmt.Scan(&userName)

		fmt.Print("Enter your email")
		fmt.Scan(&email)

		fmt.Print("How many tickets do you want? ")
		fmt.Scan(&numberOfTickets)

		isUserNameValid, isEmailValid, isNumTicketsValid := shared.ValidateInput(userName, email, numberOfTickets, availableTickets)
		if isUserNameValid && isEmailValid && isNumTicketsValid {
			bookTicket(numberOfTickets, userName, email)
			wg.Add(1)
			go sendTicket(email, numberOfTickets)

		}
		if availableTickets == 0 {
			printFinalMessage()
			break
		}

	}
	wg.Wait()
}

func bookTicket(numberOfTickets int, userName string, email string) {
	printBookingMessage(numberOfTickets)
	var userData = UserData{
		name:            userName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}
	bookingData = append(bookingData, userData)
}

func printBookingMessage(numberOfTickets int) {
	availableTickets = availableTickets - numberOfTickets
	fmt.Printf("You have successfully booked %v tickets.\n", numberOfTickets)
	fmt.Printf("Total number of tickets are:%v and available number of tickets are:%v\n", shared.TotalNumberOfTickets, availableTickets)

}
func printFinalMessage() {
	for _, data := range bookingData {
		fmt.Printf("The booked user is %v\n", data.name)
	}
	fmt.Printf("All tickets booked.The bookings are in the names of %v\n", bookingData)
}
func sendTicket(email string, numberOfTickets int) {

	time.Sleep(10 * time.Second)
	fmt.Println("######################")
	fmt.Printf("Sending ticket to %v for %v tickets", email, numberOfTickets)
	fmt.Println("######################")
	wg.Done()
}
