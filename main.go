package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // initialize a list of maps


type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){

	greetUsers()

	for{

		firstName, lastName, email, userTickets := getUserInput()
		validName, validEmail, validTicketNum := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if validEmail && validName && validTicketNum {

			bookTicket( firstName, lastName, userTickets, email)
			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the users: %v\n", firstNames)



			if remainingTickets == 0 {
				fmt.Println("Sorry the conference is booked")
				break
			}
		} else {
			// fmt.Printf("we only have %v remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
			if !validName {
				fmt.Println("User Name input is incorrect")
			}
			if !validEmail {
				fmt.Println("User Email input is incorrect")
			}
			if !validTicketNum {
				fmt.Println("User Ticket Number input is incorrect")
			}
		}

		
	}
}

func greetUsers(){
	fmt.Printf("Hello welcome to %v \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available \n", conferenceTickets, remainingTickets)

}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}



func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName,email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	//creating map for user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings : %v \n", bookings)

	fmt.Printf("User %v %v booked %v tickets, you will receive an email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets reamaining for the %v \n", remainingTickets, conferenceName)

}
// function to print ticket details
func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("######################")
}