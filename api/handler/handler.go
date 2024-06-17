package handler

import pb "github.com/Salikhov079/library-auth/genprotos"

type Handler struct {
	User pb.UserServiceClient
	Borrower pb.BorrowerServiceClient
}

func NewHandler(user pb.UserServiceClient, borrower pb.BorrowerServiceClient) *Handler {
	return &Handler{user, borrower}
}
