package main

import (
	"fmt"
	"sync"
	"ticketbooking/internal/booking"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Ticket Booking CLI")
	fmt.Println("-------------------")
	fmt.Println("Event Name: Conference")
	fmt.Println("Event Date: TBD")
	fmt.Println("Total Tickets:", booking.TotalTickets)
	fmt.Println("Remaining Tickets:", booking.RemainingTickets)

	var choice string = "y"

	for booking.RemainingTickets > 0 && (choice == "y" || choice == "Y") {

		fmt.Println("-------------------")
		fmt.Println("To book your tickets, please provide your details.")

		var firstName, lastName, email string
		var tickets uint

		for {
			fmt.Print("First Name: ")
			fmt.Scan(&firstName)
			if booking.IsValidName(firstName) {
				break
			}
			fmt.Println("❌ First name must be at least 2 characters. Try again.")
		}

		for {
			fmt.Print("Last Name: ")
			fmt.Scan(&lastName)
			if booking.IsValidName(lastName) {
				break
			}
			fmt.Println("❌ Last name must be at least 2 characters. Try again.")
		}

		for {
			fmt.Print("Email: ")
			fmt.Scan(&email)
			if booking.IsValidEmail(email) {
				break
			}
			fmt.Println("❌ Email must be valid. Try again.")
		}

		for {
			fmt.Println("INFO: Max only 5 tickets can be purchased.")
			fmt.Printf("How many tickets do you want? ")
			fmt.Scan(&tickets)

			if !booking.IsValidTicketNumber(tickets) {
				fmt.Println("❌ You can only buy a maximum of 5 tickets.")
				continue
			}

			if tickets > booking.RemainingTickets {
				fmt.Printf("❌ Only %d tickets are left, not %d.\n",
					booking.RemainingTickets, tickets)

				var accept string
				fmt.Printf("Do you want to buy %d tickets instead? (y/n): ",
					booking.RemainingTickets)
				fmt.Scan(&accept)

				if accept == "y" || accept == "Y" {
					tickets = booking.RemainingTickets
					break
				}

				continue
			}

			if booking.HasEnoughTickets(tickets, booking.RemainingTickets) {
				break
			}

			fmt.Println("❌ Invalid ticket count. Try again.")
		}

		booking.BookingTicket(firstName, lastName, email, tickets)
		fmt.Printf("Thank you %s %s for booking %d tickets. A confirmation email will be sent to %s.\n",
			firstName, lastName, tickets, email)

		wg.Add(1)
		go booking.SendTicket(tickets, firstName, lastName, email, &wg)
		fmt.Println("Remaining Tickets:", booking.RemainingTickets)

		if booking.RemainingTickets == 0 {
			fmt.Println("\n🎉 All tickets sold out!")
			break
		}

		fmt.Print("Do you want to add more bookings? (y/n): ")
		fmt.Scan(&choice)
	}

	for _, b := range booking.Bookings {
		fmt.Printf("Booking ID %d: %s %s booked %d tickets\n",
			b.ID, b.FirstName, b.LastName, b.Tickets)
	}

	wg.Wait()
}
