package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-straming-sample/grpc-sever/gen/go/hello"
	"log"
	"net"
	"time"
)

type MessengerServer struct {
	pb.UnimplementedMessengerServer
}

func (s *MessengerServer) SayHello(request *pb.HelloRequest, stream pb.Messenger_SayHelloServer) error {
	for i := 0; i < int(request.Count); i++ {
		if err := stream.Send(&pb.HelloReply{Message: "Hello " + request.Name}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &MessengerServer{})
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
