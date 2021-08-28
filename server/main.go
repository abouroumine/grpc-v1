package main

import (
	pb "abouroumine.com/grpc-v1/usermg"
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

type UserMgServer struct {
	pb.UnimplementedUserMgServer
}

func (s *UserMgServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received New User: %v", in.GetName())
	var userId = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Userid: userId}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserMgServer(s, &UserMgServer{})
	log.Printf("Server Listening at: %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed To Listen: %v\n", err)
	}
}
