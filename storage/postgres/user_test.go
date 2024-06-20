package postgres

import (
	"log"
	"testing"

	"github.com/Salikhov079/library-auth/config"
	pb "github.com/Salikhov079/library-auth/genprotos"
	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/assert"
)

var cnf = config.Load()

func TestRegister(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.UserReq{
		UserName: faker.Username(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}
	result, err := stg.User().Register(user)

	assert.NoError(t, err, "Error during user registration: %v", err)
	assert.Nil(t, result, "Result should not be nil")
}

func TestGetById(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "cb723577-0729-4887-813a-37c41ff76036"

	user, err := stg.User().GetById(&Id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAll(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	users, err := stg.User().GetAll(&pb.UserReq{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUpdateUser(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.UserRes{
		Id:       "cb723577-0729-4887-813a-37c41ff76036",
		UserName: faker.Username(),
		Email:    faker.Email(),
	}

	result, err := stg.User().Update(user)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteUser(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById
	Id.Id = "cb723577-0729-4887-813a-37c41ff76036"
	result, err := stg.User().Delete(&Id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestLogin(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.UserReq{
		UserName: "najmiddin",
	}

	result, err := stg.User().Login(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetAllId(t *testing.T) {
	stg, err := NewPostgresStorage(&cnf)
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	users, err := stg.User().GetAllId(&pb.Void{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}
