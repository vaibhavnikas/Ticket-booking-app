package main

import "strings"

func validateInput(userName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) > 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTickets := userTickets > 0
	return isValidName, isValidEmail, isValidTickets
}