package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/library-auth/api"
	"github.com/Salikhov079/library-auth/api/handler"
	"github.com/Salikhov079/library-auth/config"
	pb "github.com/Salikhov079/library-auth/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cnf := config.Load()
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
	}

	connU, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
	}

	borrower := pb.NewBorrowerServiceClient(conn)

	user := pb.NewUserServiceClient(connU)

	h := handler.NewHandler(user, borrower)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(cnf.HTTPPort)
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
