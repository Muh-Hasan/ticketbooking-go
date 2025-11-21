# Ticket Booking CLI

A simple command-line ticket booking application built in Go. Created to practice Go syntax, struct usage, slices, validation, and basic concurrency using goroutines.

## Features
- Simple interactive CLI for booking tickets.
- Demonstrates Go structs, slices, and validation logic.
- Uses goroutines for basic concurrent operations.

## Build & Run
From the repository root:

```bat
go run ./cmd/app
```

Or build a binary and run:

```bat
go build -o ticketbooking ./cmd/app
ticketbooking.exe
```

## Project Structure
- `cmd/app/main.go` — CLI entrypoint.
- `internal/booking/booking.go` — core booking logic.
- `internal/booking/validator.go` — input validation helpers.
- `go.mod` — module definition.

