package main

import (
    "context"
    "errors"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/lucilu1320/Train-Services" // Adjust the path to match your project structure
)

type server struct {
    tickets map[string]*pb.Receipt
    availableSeats map[string][]string // Map to track available seats in each section
}

func (s *server) PurchaseTicket(ctx context.Context, req *pb.TicketRequest) (*pb.Receipt, error) {
    if _, ok := s.tickets[req.Email]; ok {
        return nil, errors.New("ticket already exists for this email")
    }

    seat := allocateSeat(s.availableSeats, req.Section)
    if seat == "" {
        return nil, errors.New("no seats available in the requested section")
    }

    receipt := &pb.Receipt{
        From: req.From,
        To: req.To,
        User: req.FirstName + " " + req.LastName,
        PricePaid: 20.00, // Example price
        Seat: seat,
        Section: req.Section,
    }

    s.tickets[req.Email] = receipt
    s.availableSeats[req.Section] = removeSeat(s.availableSeats[req.Section], seat)

    return receipt, nil
}

// ... (other methods remain the same)

// Allocate a seat in the specified section
func allocateSeat(availableSeats map[string][]string, section string) string {
    if availableSeats[section] == nil || len(availableSeats[section]) == 0 {
        return "" // No seats available
    }
    return availableSeats[section][0] // Allocate the first available seat
}

// Remove a seat from the list of available seats
func removeSeat(seats []string, seatToRemove string) []string {
    for i, seat := range seats {
        if seat == seatToRemove {
            return append(seats[:i], seats[i+1:]...)
        }
    }
    return seats // Seat not found
}

func main() {
    // Initialize available seats
    availableSeats := map[string][]string{
        "A": []string{"A1", "A2", "A3", /* ... */}, // Adjust seat lists as needed
        "B": []string{"B1", "B2", "B3", /* ... */},
    }

    lis, err := net.Listen("tcp", ":50051") // Adjust port if needed
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTicketServiceServer(s, &server{tickets: make(map[string]*pb.Receipt), availableSeats: availableSeats})
    log.Println("Server listening on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
