package handler

import pb "github.com/Salikhov079/library-auth/service"

type Handler struct {
	User *pb.UserService
}

func NewHandler(us *pb.UserService) *Handler {
	return &Handler{us}
}
