package postgres

import (
	pb "github.com/Salikhov079/library-auth/genprotos"
)

type StorageRoot interface {
	User() User
}

type User interface {
	Register(user *pb.User) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.User, error)
	GetAll(_ *pb.User) (*pb.AllUsers, error)
	Update(user *pb.User) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	Login(user *pb.User ) (*pb.User, error)
}
