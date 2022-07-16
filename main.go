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
		user := booking.User{}

		_, nameErr := booking.GetUserName(&user.Name)
		if nameErr != nil {
			fmt.Println(nameErr)
			continue
		}

		_, countErr := booking.GetTicketCount(&user.Tickets)
		if countErr != nil {
			fmt.Println(countErr)
			continue
		}

		_, emailErr := booking.GetUserEmail(&user.Email)
		if emailErr != nil {
			fmt.Println(emailErr)
			continue
		}

		booking.BookTicket(user, &conf)

		// Breaks the infinite loop
		if utils.IsZero(int(conf.Tickets)) {
			break
		}
	}

	fmt.Print("Ticket Holders: ")
	for _, v := range conf.GetConfTickets() {
		fmt.Printf("%v: %v tickets\t", v.Name, v.Tickets)
	}
	fmt.Println()
}
