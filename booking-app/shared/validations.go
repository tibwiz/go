package shared

import (
	"fmt"
	"strings"
)

const TotalNumberOfTickets = 50

func ValidateInput(userName string, email string, numberOfTickets int, availableTickets int) (bool, bool, bool) {
	var isUserNameValid = len(userName) > 2
	if !isUserNameValid {
		fmt.Printf("Invalid user name %v, user name should be more than 2 chars.\n", userName)
	}

	var isEmailValid = strings.Contains(email, "@")
	if !isEmailValid {
		fmt.Printf("Entered email %v does not have @ character.\n", email)
	}
	var isNumTicketsValid = numberOfTickets > 0 && numberOfTickets <= availableTickets
	if !isNumTicketsValid {
		fmt.Printf("Number of tickets should be greater than 0 and less than %v\n", availableTickets)
	}
	return isUserNameValid, isEmailValid, isNumTicketsValid
}

func GreetUser(conferenceName string, availableTickets int) {
	fmt.Printf("Welcome to %v ticket booking app\n", conferenceName)
	fmt.Printf("Total number of tickets are:%v and available number of tickets are:%v\n", TotalNumberOfTickets, availableTickets)

}
