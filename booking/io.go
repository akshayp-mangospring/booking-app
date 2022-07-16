package booking

import (
	"booking-app/utils"
	"fmt"
)

func CheckAndAcceptUserName(name *string) bool {
	fmt.Println("Enter Name:")
	fmt.Scanf("%s", name)

	if utils.IsNameValid(*name) {
		fmt.Println("Invalid Name. Try again")
		return true
	}
	return false
}

func CheckAndAcceptTicketCount(count *uint8) bool {
	fmt.Println("Enter Count:")
	fmt.Scanf("%d", count)

	if utils.IsOutOfRange(0, 255, int(*count)) {
		fmt.Println("Invalid Count. Try again")
		return true
	}
	return false
}
