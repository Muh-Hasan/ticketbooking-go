package booking

import "strings"

func IsValidName(name string) bool {
	return len(name) >= 2
}

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func IsValidTicketNumber(requested uint) bool {
	return requested > 0 && requested <= 5
}

func HasEnoughTickets(requested, remaining uint) bool {
	return requested <= remaining
}
