package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kurojs/grabvn-golang-bootcamp/exercise/user-manament/proto"
	"github.com/kurojs/grabvn-golang-bootcamp/exercise/user-manament/server/db"
)

type gRPCServer struct {
	Db *db.MySQLClient
}

// StartGRPCServer ...
func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50501")
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

	proto.RegisterUserServiceServer(srv, &gRPCServer{Db: db.GetInstance()})
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (server *gRPCServer) CreateUser(ctx context.Context, r *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	userID := int32(time.Now().Unix())
	user := db.User{
		ID:       userID,
		Username: r.User.Username,
	}

	err := server.Db.Save(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to save user info")
	}

	return &proto.CreateUserResponse{UserID: userID}, nil
}

func (server *gRPCServer) ListUser(ctx context.Context, r *proto.ListUserRequest) (*proto.ListUserResponse, error) {
	var users []db.User

	err := server.Db.Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.NotFound, "User list is empty")
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get user list")
	}

	var userList []*proto.User
	for _, user := range users {
		userList = append(userList, &proto.User{ID: user.ID, Username: user.Username})
	}

	return &proto.ListUserResponse{Users: userList}, nil
}

func (server *gRPCServer) GetUserDetail(ctx context.Context, r *proto.GetUserDetailRequest) (*proto.GetUserDetailResponse, error) {
	var user db.User

	err := server.Db.Where(&db.User{ID: r.UserID}).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get user info")
	}

	return &proto.GetUserDetailResponse{
		User: &proto.User{
			ID:       user.ID,
			Username: user.Username,
		},
	}, nil
}

func (server *gRPCServer) UpdateUserDetail(ctx context.Context, r *proto.UpdateUserDetailRequest) (*proto.UpdateUserDetailResponse, error) {
	user := db.User{
		ID:       r.User.ID,
		Username: r.User.Username,
	}

	err := server.Db.Save(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to save user info")
	}

	return &proto.UpdateUserDetailResponse{UserID: user.ID}, nil
}

func (server *gRPCServer) DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := server.Db.Delete(&db.User{ID: r.UserID}).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete user")
	}

	return &proto.DeleteUserResponse{UserID: r.UserID}, nil
}
