package booking

import "fmt"

// Public Exports

func GetMarketingLine(s string) string {
	/*
		Sprintf is actually used to store a formatted string into a var
		It doesn't actually print a string, it returns a string which can be printed later
		Thus it can effectively be used as a return type for functions
	*/

	return fmt.Sprintf("Welcome to our conference booking application!\nThe only stop to get your %s tickets", s)
}

// Private funcs within a package

func printNoTicketsLeft() {
	fmt.Println("Conference is sold out! No more tickets left.")
}
