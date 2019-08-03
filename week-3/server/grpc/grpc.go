package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	"github.com/kurojs/grabvn-golang-bootcamp/week-3/db"
	"github.com/kurojs/grabvn-golang-bootcamp/week-3/proto"
	"github.com/kurojs/grabvn-golang-bootcamp/week-3/returncode"
)

const port = ":50051"

// StartGRPCServer ...
func StartGRPCServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			srv.GracefulStop()
		}
	}()

	proto.RegisterPassengerFeedbackServiceServer(srv, &gRPCServer{Db: db.GetInstance()})
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type gRPCServer struct {
	Db *db.MySQLClient
}

func (server *gRPCServer) AddFeedback(ctx context.Context, req *proto.AddFeedbackRequest) (*proto.AddFeedbackResponse, error) {
	// Check duplicata
	err := server.Db.First(&db.PassengerFeedback{BookingCode: req.BookingCode}).Error
	if err != gorm.ErrRecordNotFound {
		return &proto.AddFeedbackResponse{
			ReturnCode:    returncode.DuplicateFeedback,
			ReturnMessage: returncode.Message(returncode.DuplicateFeedback),
		}, nil
	}

	feedback := &db.PassengerFeedback{
		BookingCode: req.BookingCode,
		FeedBack:    req.Feedback,
		PassengerID: req.PassengerID,
	}
	err = server.Db.Save(feedback).Error
	return &proto.AddFeedbackResponse{
		ReturnCode:    returncode.OK,
		ReturnMessage: returncode.Message(returncode.OK),
	}, nil
}

func (server *gRPCServer) GetFeedbackByPassengerID(ctx context.Context, req *proto.GetFeedbackByPassengerIDRequest) (*proto.GetFeedbackByPassengerIDResponse, error) {
	var feedbacks []db.PassengerFeedback

	err := server.Db.Where(&db.PassengerFeedback{PassengerID: req.PassengerID}).Find(&feedbacks).Error
	if err == gorm.ErrRecordNotFound {
		return &proto.GetFeedbackByPassengerIDResponse{
			ReturnCode:    returncode.NotFound,
			ReturnMessage: returncode.Message(returncode.NotFound),
		}, nil
	}

	if err != nil {
		return &proto.GetFeedbackByPassengerIDResponse{
			ReturnCode:    returncode.Unknown,
			ReturnMessage: err.Error(),
		}, nil
	}

	var result []*proto.PassengerFeedback
	for _, feedback := range feedbacks {
		result = append(result, &proto.PassengerFeedback{
			BookingCode: feedback.BookingCode,
			Feedback:    feedback.FeedBack,
			PassengerID: feedback.PassengerID,
		})
	}

	return &proto.GetFeedbackByPassengerIDResponse{
		Feedbacks:     result,
		ReturnCode:    returncode.OK,
		ReturnMessage: returncode.Message(returncode.OK),
	}, nil
}

func (server *gRPCServer) GetFeedbackByBookingCode(ctx context.Context, req *proto.GetFeedbackByBookingCodeRequest) (*proto.GetFeedbackByBookingCodeResponse, error) {
	var feedback db.PassengerFeedback

	err := server.Db.Where(&db.PassengerFeedback{BookingCode: req.BookingCode}).Find(&feedback).Error
	if err == gorm.ErrRecordNotFound {
		return &proto.GetFeedbackByBookingCodeResponse{
			ReturnCode:    returncode.NotFound,
			ReturnMessage: returncode.Message(returncode.NotFound),
		}, nil
	}

	if err != nil {
		return &proto.GetFeedbackByBookingCodeResponse{
			ReturnCode:    returncode.Unknown,
			ReturnMessage: err.Error(),
		}, nil
	}

	return &proto.GetFeedbackByBookingCodeResponse{
		Feedback: &proto.PassengerFeedback{
			BookingCode: feedback.BookingCode,
			Feedback:    feedback.FeedBack,
			PassengerID: feedback.PassengerID,
		},
		ReturnCode:    returncode.OK,
		ReturnMessage: returncode.Message(returncode.OK),
	}, nil
}

func (server *gRPCServer) DeleteFeedbackByPassengerID(ctx context.Context, req *proto.DeleteFeedbackByPassengerIDRequest) (*proto.DeleteFeedbackByPassengerIDResponse, error) {
	err := server.Db.Where(&db.PassengerFeedback{PassengerID: req.PassengerID}).Delete(&db.PassengerFeedback{}).Error
	if err == gorm.ErrRecordNotFound {
		return &proto.DeleteFeedbackByPassengerIDResponse{
			ReturnCode:    returncode.NotFound,
			ReturnMessage: returncode.Message(returncode.NotFound),
		}, nil
	}

	if err != nil {
		return &proto.DeleteFeedbackByPassengerIDResponse{
			ReturnCode:    returncode.Unknown,
			ReturnMessage: err.Error(),
		}, nil
	}

	return &proto.DeleteFeedbackByPassengerIDResponse{
		ReturnCode:    returncode.OK,
		ReturnMessage: returncode.Message(returncode.OK),
	}, nil
}
