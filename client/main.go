package main

import (
    "context"
    "fmt"
    "log"
    pb "github.com/lucilu1320/Train-Services/proto" //Adjust the path to match your project structure
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //Replace with server port as mentioned in the server/main.go code
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewTicketServiceClient(conn)

    // Purchase a ticket
    ticketRequest := &pb.TicketRequest{
        Email:      "sahil.doe@example.com",
        FirstName:  "sahil",
        LastName:   "Doe",
        From:       "City A",
        To:         "City B",
        Section:    "A",
    }
    receipt, err := client.PurchaseTicket(context.Background(), ticketRequest)
    if err != nil {
        log.Fatalf("failed to purchase ticket: %v", err)
    }
    fmt.Println("Purchased ticket:", receipt)

    //Get receipt details
    receiptDetails, err := client.GetReceiptDetails(context.Background(), &pb.User{Email: "jane.doe@example.com"})
    if err != nil {
        log.Fatalf("failed to get receipt details: %v", err)
    }
    fmt.Println("Receipt details:", receiptDetails)

    //View users and seats in a section
    userSeatList, err := client.ViewUsersAndSeats(context.Background(), &pb.Section{Section: "A"})
    if err != nil {
        log.Fatalf("failed to view users and seats: %v", err)
    }
    fmt.Println("Users and seats in section A:", userSeatList)

    //Modify a user's seat to a different available seat
    seatModification := &pb.UserSeatModification{
        Email:    "sahil.doe@example.com",
        NewSeat:  "B6",
    }
    _, err = client.ModifyUserSeat(context.Background(), seatModification)
    if err != nil {
        log.Fatalf("failed to modify user seat: %v", err)
    }
    fmt.Println("User seat modified")

    //Remove a user 
    removeResponse, err := client.RemoveUser(context.Background(), &pb.User{Email: "jane.doe@example.com"})
    if err != nil {
        log.Fatalf("failed to remove user: %v", err)
    }
    fmt.Println("User removed:", removeResponse.String())
}
