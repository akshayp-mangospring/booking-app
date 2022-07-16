package booking

import (
	"booking-app/utils"
	"fmt"
)

// Usage of pointers for altering data directly at memory location
func BookTicket(u User, c *Conf) {
	// Type conversion from `uint8` to `int`
	if utils.IsZero(int(c.Tickets)) {
		printNoTicketsLeft()
		return
	}

	if u.Tickets > c.Tickets {
		fmt.Printf("%s booked the last %d tickets out of %d tickets requested. 0 tickets left.\n", u.Name, c.Tickets, u.Tickets)
		u.Tickets = c.Tickets
		c.users = append(c.users, u)
		c.Tickets = 0
		printNoTicketsLeft()
		return
	}

	c.Tickets = c.Tickets - u.Tickets
	c.users = append(c.users, u)
	fmt.Printf("%s booked %d tickets. %d tickets left\n", u.Name, u.Tickets, c.Tickets)
}
