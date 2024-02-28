package grpc

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Register(port, register func(s *grpc2.Server)) error {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc2.NewServer()
	reflection.Register(s)
	register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return err
}
