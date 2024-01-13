package main

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/grpc"

    pb "github.com/lucilu1320/Train-Services/" // Adjust the path to match your project structure
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Replace with server address
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewTicketServiceClient(conn)

    // Purchase a ticket
    ticketRequest := &pb.TicketRequest{
        Email:      "john.doe@example.com",
        FirstName:  "John",
        LastName:   "Doe",
        From:       "City A",
        To:         "City B",
        Section:    "A", // Specify the desired section
    }
    receipt, err := client.PurchaseTicket(context.Background(), ticketRequest)
    if err != nil {
        log.Fatalf("failed to purchase ticket: %v", err)
    }
    fmt.Println("Purchased ticket:", receipt)

    // Get receipt details
    receiptDetails, err := client.GetReceiptDetails(context.Background(), &pb.User{Email: "john.doe@example.com"})
    if err != nil {
        log.Fatalf("failed to get receipt details: %v", err)
    }
    fmt.Println("Receipt details:", receiptDetails)

    // View users and seats in a section
    userSeatList, err := client.ViewUsersAndSeats(context.Background(), &pb.Section{Section: "A"})
    if err != nil {
        log.Fatalf("failed to view users and seats: %v", err)
    }
    fmt.Println("Users and seats in section A:", userSeatList)

    // Remove a user
    removeResponse, err := client.RemoveUser(context.Background(), &pb.User{Email: "john.doe@example.com"})
    if err != nil {
        log.Fatalf("failed to remove user: %v", err)
    }
    fmt.Println("User removed:", removeResponse.String())

    // Modify a user's seat
    seatModification := &pb.UserSeatModification{
        User:      "john.doe@example.com",
        NewSeat:   "B2",
    }
    _, err = client.ModifyUserSeat(context.Background(), seatModification)
    if err != nil {
        log.Fatalf("failed to modify user seat: %v", err)
    }
    fmt.Println("User seat modified")
}
