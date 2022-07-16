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
		return *name, nil
	}
	return *name, errors.New("Invalid Name. Try again")
}

func GetTicketCount(count *uint8) (uint8, error) {
	fmt.Println("Enter Count:")
	fmt.Scanf("%d", count)

	if utils.IsInRange(0, 255, int(*count)) {
		return *count, nil
	}
	return *count, errors.New("Invalid Count. Try again")
}

func GetUserEmail(email *string) (string, error) {
	fmt.Println("Enter Email:")
	fmt.Scanf("%s", email)

	if utils.IsEmailValid(*email) {
		return *email, nil
	}
	return *email, errors.New("Invalid Email. Try again")
}
