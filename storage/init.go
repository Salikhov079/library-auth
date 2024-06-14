package postgres

import (
	pb "github.com/Salikhov079/library-auth/genprotos"
)

type StorageRoot interface {
	User() User
}

type User interface {
	Register(user *pb.UserReq) (*pb.Void, error)
	Update(user *pb.UserRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.UserRes, error)
	GetAll(_ *pb.UserReq) (*pb.AllUsers, error)
	Login(user *pb.UserReq) (*pb.UserRes, error)
}
