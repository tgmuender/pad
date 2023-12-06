package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"xgmdr.com/pad/internal/api"

	"google.golang.org/grpc"
	pb "xgmdr.com/pad/proto"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPetServiceServer(s, &api.PetApi{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
