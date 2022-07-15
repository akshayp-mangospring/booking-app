package main

import "fmt"

func main() {
	confName := "GopherConf 2022"

	// Constants in Go are untyped, hence we can declare them without even declaring a type
	const ticketsCount uint8 = 50
	remainingTickets := ticketsCount

	fmt.Println("Welcome to our conference booking application!")
	fmt.Println(getMarketingLine(confName))
	fmt.Printf("We have %d tickets left. Book quickly to get your spot!\n", remainingTickets)

	bookTicket("Akshay", 5, &remainingTickets)
	bookTicket("Sourabh", 5, &remainingTickets)
	bookTicket("Nivedita", 5, &remainingTickets)
	bookTicket("Akshay", 5, &remainingTickets)
	bookTicket("Sourabh", 5, &remainingTickets)
	bookTicket("Nivedita", 5, &remainingTickets)
	bookTicket("Akshay", 5, &remainingTickets)
	bookTicket("Sourabh", 5, &remainingTickets)
	bookTicket("Nivedita", 5, &remainingTickets)
	bookTicket("Akshay", 5, &remainingTickets)
	bookTicket("Sourabh", 5, &remainingTickets)
	bookTicket("Nivedita", 5, &remainingTickets)
}

func getMarketingLine(s string) string {
	/*
		Sprintf is actually used to store a formatted string into a var
		It doesn't actually print a string, it returns a string which can be printed later
		Thus it can effectively be used as a return type for functions
	*/

	return fmt.Sprintf("The only stop to get your %s tickets", s)
}

func bookTicket(name string, count uint8, leftTickets *uint8) {
	if *leftTickets == 0 {
		fmt.Println("Conference is sold out! No more tickets left.")
		return
	}

	*leftTickets = *leftTickets - count
	fmt.Printf("%s booked %d tickets. %d tickets left\n", name, count, *leftTickets)
}
