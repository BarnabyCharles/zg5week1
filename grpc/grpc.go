package grpc

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func Register(port int, register func(s *grpc2.Server), cert, key string) error {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	s := grpc2.NewServer(grpc2.Creds(creds))
	// 反射查询
	reflection.Register(s)
	register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return err
}
