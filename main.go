package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/library-auth/config"
	"github.com/Salikhov079/library-auth/api"
	"github.com/Salikhov079/library-auth/api/handler"
	pb "github.com/Salikhov079/library-auth/service"
	"github.com/Salikhov079/library-auth/storage/postgres"
)

func main() {
	cnf := config.Load()
	db, err := postgres.NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	us := pb.NewUserService(db)

	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(cnf.HTTPPort)
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
