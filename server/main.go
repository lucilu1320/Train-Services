package main

import (
    "context"
    "errors"
    "log"
    "net"
    "sync"
    pb "github.com/lucilu1320/Train-Services/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedTicketServiceServer
    tickets         map[string]*pb.Receipt
    availableSeats  map[string][]string
    mutex           sync.Mutex
}

func newServer() *server {
    return &server{
        tickets: make(map[string]*pb.Receipt),
        availableSeats: map[string][]string{
            "A": []string{"A1", "A2", "A3"}, //Example seats for each section
            "B": []string{"B1", "B2", "B3"},
        },
    }
}

func (s *server) allocateSeat(section string) (string, error) {
    if seats, ok := s.availableSeats[section]; ok && len(seats) > 0 {
        allocatedSeat := seats[0]
        s.availableSeats[section] = seats[1:]
        return allocatedSeat, nil
    }
    return "", errors.New("no seats available in the requested section")
}


//Purchase Ticket handles the ticket purchase
func (s *server) PurchaseTicket(ctx context.Context, req *pb.TicketRequest) (*pb.Receipt, error) {
s.mutex.Lock()
defer s.mutex.Unlock()

if _, ok := s.tickets[req.Email]; ok {
return nil, errors.New("a ticket already exists for this email")
}

seat, err := s.allocateSeat(req.Section)
if err != nil {
return nil, err
}

receipt := &pb.Receipt{
From: req.From,
To: req.To,
User: req.FirstName + " " + req.LastName,
PricePaid: 20.00,
Seat: seat,
Section: req.Section,
}

s.tickets[req.Email] = receipt
return receipt, nil
}

//Get ReceiptDetails returns the details of a receipt for a given user
func (s *server) GetReceiptDetails(ctx context.Context, req *pb.User) (*pb.Receipt, error) {
s.mutex.Lock()
defer s.mutex.Unlock()

if receipt, ok := s.tickets[req.Email]; ok {
return receipt, nil
}
return nil, errors.New("no receipt found for the provided email")
}

//View Users And Seats lists users and their seats in a specific section
func (s *server) ViewUsersAndSeats(ctx context.Context, req *pb.Section) (*pb.UserSeatList, error) {
s.mutex.Lock()
defer s.mutex.Unlock()

userSeats := &pb.UserSeatList{}
for _, receipt := range s.tickets {
if receipt.Section == req.Section {
userSeat := &pb.UserSeat{
User: receipt.User,
Seat: receipt.Seat,
}
userSeats.UserSeat = append(userSeats.UserSeat, userSeat)
}
}
return userSeats, nil
}

//Remove User removes a user's ticket from the system
func (s *server) RemoveUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {
s.mutex.Lock()
defer s.mutex.Unlock()
if receipt, ok := s.tickets[req.Email]; ok {
//Add the seat back to available seats
s.availableSeats[receipt.Section] = append(s.availableSeats[receipt.Section], receipt.Seat)
delete(s.tickets, req.Email)
return &pb.Empty{}, nil
}
return nil, errors.New("user not found")
}

//Modify User Seat changes a user's seat
func (s *server) ModifyUserSeat(ctx context.Context, req *pb.UserSeatModification) (*pb.Empty, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    receipt, ok := s.tickets[req.Email] //Correct field name from .proto file
    if !ok {
        return nil, errors.New("user not found")
    }

    //Check if the new seat is available
    for i, seat := range s.availableSeats[receipt.Section] {
        if seat == req.NewSeat { 
            //Update the seat
            receipt.Seat = req.NewSeat
            s.availableSeats[receipt.Section] = append(s.availableSeats[receipt.Section][:i], s.availableSeats[receipt.Section][i+1:]...)
            return &pb.Empty{}, nil
        }
    }
    return nil, errors.New("requested seat not available")
}




func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTicketServiceServer(s, newServer())
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
