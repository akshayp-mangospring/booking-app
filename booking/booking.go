package booking

import (
	"booking-app/utils"
	"fmt"
)

// Usage of pointers for altering data directly at memory location
func BookTicket(name string, count uint8, email string, c *Conf) {
	// Type conversion from `uint8` to `int`
	if utils.IsZero(int(c.Tickets)) {
		printNoTicketsLeft()
		return
	}

	if count > c.Tickets {
		fmt.Printf("%s booked the last %d tickets out of %d tickets requested. 0 tickets left.\n", name, c.Tickets, count)
		c.users = append(c.users, User{Name: name, Email: email, Tickets: c.Tickets})
		c.Tickets = 0
		printNoTicketsLeft()
		return
	}

	c.Tickets = c.Tickets - count
	c.users = append(c.users, User{Name: name, Email: email, Tickets: count})
	fmt.Printf("%s booked %d tickets. %d tickets left\n", name, count, c.Tickets)
}
