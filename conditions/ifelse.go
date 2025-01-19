package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	bookings := []string{}

	fmt.Printf("Welcome to  %v booking application\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length : %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Remaining tickets are: %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Oh!! Sorry ;( , Our conferecne is booked out, Come back next year!. ")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name u entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email is not valid!")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets are invalid")
			}
		}

	}

}
