package main

import (
	"context"
	"log"
	"net"

	"github.com/viniciusmazon/lab-go-grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("Message received:: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}

func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpc_Error := grpcServer.Serve(listener)
	if grpc_Error != nil {
		log.Fatal(err)
	}

	log.Print("Server is running")
}
