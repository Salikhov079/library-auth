package main

import (
	"log"
	"net"

	"github.com/Salikhov079/library-auth/config"
	pb "github.com/Salikhov079/library-auth/genprotos"
	"github.com/Salikhov079/library-auth/service"
	"github.com/Salikhov079/library-auth/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cnf := config.Load()
	db, err := postgres.NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	liss, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
