package booking

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
