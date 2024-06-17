package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/library-auth/genprotos"
	s "github.com/Salikhov079/library-auth/storage"
)

type UserService struct {
	stg s.StorageRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.StorageRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) RegisterUser(ctx context.Context, user *pb.UserReq) (*pb.Void, error) {
	pb, err := c.stg.User().Register(user)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *UserService) GetAllUsers(ctx context.Context, pb *pb.UserReq) (*pb.AllUsers, error) {
	Users, err := c.stg.User().GetAll(pb)
	if err != nil {
		log.Print(err)
	}

	return Users, err
}

func (c *UserService) GetByIdUser(ctx context.Context, id *pb.ById) (*pb.UserRes, error) {
	prod, err := c.stg.User().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *UserService) UpdateUser(ctx context.Context, User *pb.UserRes) (*pb.Void, error) {
	pb, err := c.stg.User().Update(User)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) DeleteUser(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.User().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) LoginUser(ctx context.Context, username *pb.UserReq) (*pb.UserRes, error) {
	prod, err := c.stg.User().Login(username)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *UserService) GetAllIdUsers(ctx context.Context, v *pb.Void) (*pb.AllUsers, error) {
	res, err := c.stg.User().GetAllId(v)
	if err != nil {
		log.Print(err)
	}

	return res, err
}
