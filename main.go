package main
import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}