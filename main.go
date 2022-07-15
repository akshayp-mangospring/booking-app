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

	for {
		var name string
		var count int

		fmt.Println("Enter Name:")
		fmt.Scanf("%s", &name)
		fmt.Println("Enter Count:")
		fmt.Scanf("%d", &count)

		bookTicket(name, count, &remainingTickets)

		if isZero(int(remainingTickets)) {
			break
		}
	}
}

func isZero(count int) bool {
	return count == 0
}

func printNoTicketsLeft() {
	fmt.Println("Conference is sold out! No more tickets left.")
}

func getMarketingLine(s string) string {
	/*
		Sprintf is actually used to store a formatted string into a var
		It doesn't actually print a string, it returns a string which can be printed later
		Thus it can effectively be used as a return type for functions
	*/

	return fmt.Sprintf("The only stop to get your %s tickets", s)
}

func bookTicket(name string, count int, leftTickets *uint8) {
	// Type conversion from `uint8` to `int`
	if isZero(int(*leftTickets)) {
		printNoTicketsLeft()
		return
	}

	if count > int(*leftTickets) {
		fmt.Printf("%s booked the last %d tickets out of %d tickets requested. 0 tickets left.\n", name, *leftTickets, count)
		*leftTickets = 0
		printNoTicketsLeft()
		return
	}

	*leftTickets = *leftTickets - uint8(count)
	fmt.Printf("%s booked %d tickets. %d tickets left\n", name, count, *leftTickets)
}
