package booking

import (
	"booking-app/utils"
	"errors"
	"fmt"
)

func GetUserName(name *string) (string, error) {
	fmt.Println("Enter Name:")
	fmt.Scanf("%s", name)

	if utils.IsNameValid(*name) {
		return *name, errors.New("Invalid Name. Try again")
	}
	return *name, nil
}

func GetTicketCount(count *uint8) (uint8, error) {
	fmt.Println("Enter Count:")
	fmt.Scanf("%d", count)

	if utils.IsOutOfRange(0, 255, int(*count)) {
		return *count, errors.New("Invalid Count. Try again")
	}
	return *count, nil
}
