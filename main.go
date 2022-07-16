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

		_, nameErr := booking.GetUserName(&name)
		if nameErr != nil {
			fmt.Println(nameErr)
			continue
		}

		_, countErr := booking.GetTicketCount(&count)
		if countErr != nil {
			fmt.Println(countErr)
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
