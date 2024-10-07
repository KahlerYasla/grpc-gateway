package main

import (
	"log"
	"net"

	"/internal/service/auth"
	genAuth "/internal/service/auth/proto/gen"
	"/src/internal/service/chat"
	genChat "/src/internal/service/chat/proto/gen"
	"/src/pkg/mid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := NewGRPCServer()
	log.Printf("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func NewGRPCServer() *grpc.Server {
	certFile := "certs/server.crt"
	keyFile := "certs/server.key"

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.StreamInterceptor(mid.StreamAuthInterceptor()),
	)
	genChat.RegisterChatServiceServer(s, &chat.ChatService{})
	genAuth.RegisterAuthServiceServer(s, &auth.AuthService{})
	return s
}
