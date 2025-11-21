package booking

import (
	"fmt"
	"sync"
	"time"
)

type Booking struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Tickets   uint
}

var TotalTickets uint = 5
var RemainingTickets uint = 5
var Bookings = make([]Booking, 0)
var bookingId uint = 1

func BookingTicket(firstName, lastName, email string, tickets uint) {
	booking := Booking{bookingId, firstName, lastName, email, tickets}

	Bookings = append(Bookings, booking)

	RemainingTickets -= tickets

	bookingId++
}

func SendTicket(tickets uint, firstName, lastName, email string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	fmt.Println("####################")
	fmt.Println("Sending Email: %d tickets for %s %s\n to email: %s", tickets, firstName, lastName, email)
	fmt.Println("####################")
	wg.Done()
}
