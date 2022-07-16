package main

import (
	"booking-app/booking"
	"booking-app/utils"
	"fmt"
)

func main() {
	conf := booking.Init("GopherConf 2022", 50)

	fmt.Println(booking.GetMarketingLine(conf.Name))
	fmt.Printf("We have %d tickets left. Book quickly to get your spot!\n", conf.Tickets)

	// Inifinte loop until all tickets are booked
	for {
		var name string
		var count uint8

		fmt.Println("Enter Name:")
		fmt.Scanf("%s", &name)

		if utils.IsNameValid(name) {
			fmt.Println("Invalid name")
			continue
		}

		fmt.Println("Enter Count:")
		fmt.Scanf("%d", &count)

		if utils.IsOutOfRange(0, 255, int(count)) {
			fmt.Println("Invalid Count")
			continue
		}

		booking.BookTicket(name, count, &conf)

		// Breaks the infinite loop
		if utils.IsZero(int(conf.Tickets)) {
			break
		}
	}

	fmt.Print("Ticket Bookers: ")
	for _, v := range conf.GetTicketBookers() {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}
