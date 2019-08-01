package main

import (
	"context"
	"log"
	"time"

	"github.com/kurojs/grabvn-golang-bootcamp/week-3/proto"
	"google.golang.org/grpc"
)

var (
	address     = "localhost:50051"
	passengerID = 1
	bookingCode = "B01"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	srv := proto.NewPassengerFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	addFeedbackReq := &proto.AddFeedbackRequest{
		BookingCode: bookingCode,
		Feedback:    "This is a feedback",
		PassengerID: int32(passengerID),
	}

	addFeedBackRes, err := srv.AddFeedback(ctx, addFeedbackReq)
	if err != nil {
		log.Printf("failed to add feedback %v", err)
	}
	log.Println("Add Feedback Respone:", addFeedBackRes)

	delFeedbackReq := &proto.DeleteFeedbackByPassengerIDRequest{
		PassengerID: int32(passengerID),
	}

	getFeedbackByBookingCodeReq := &proto.GetFeedbackByBookingCodeRequest{
		BookingCode: bookingCode,
	}
	getFeedbackByBookingCodeRes, err := srv.GetFeedbackByBookingCode(ctx, getFeedbackByBookingCodeReq)
	if err != nil {
		log.Printf("failed to get feedback by booking code %v", err)
	}
	log.Println("Get feedback by booking code:", getFeedbackByBookingCodeRes)

	getFeedbackPassengerIDCodeReq := &proto.GetFeedbackByPassengerIDRequest{
		PassengerID: int32(passengerID),
	}
	getFeedbackByPassengerIDRes, err := srv.GetFeedbackByPassengerID(ctx, getFeedbackPassengerIDCodeReq)
	if err != nil {
		log.Printf("failed to get feedback by booking code %v", err)
	}
	log.Println("Get feedback by booking code:", getFeedbackByPassengerIDRes)

	delFeedbackRes, err := srv.DeleteFeedbackByPassengerID(ctx, delFeedbackReq)
	if err != nil {
		log.Printf("failed to delete feedback by user id %v", err)
	}
	log.Println("Delete feedback:", delFeedbackRes)
}
