package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-auth/genprotos"

	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Register(user *pb.User) (*pb.Void, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO users
			(id, user_name, password, email)
		VALUES ($1, $2, $3, $4)
	`

	_, err := p.db.Exec(query, id, user.UserName, user.Password, user.Email)
	return nil, err
}

func (p *UserStorage) GetById(id *pb.ById) (*pb.User, error) {
	query := `
		SELECT 
			user_name, email from users 
		WHERE id = $1 and deleted_at = 0
	`

	row := p.db.QueryRow(query, id.Id)
	var user pb.User

	err := row.Scan(&user.UserName, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserStorage) GetAll(user *pb.User) (*pb.AllUsers, error) {
	var arr []interface{}
	count := 1

	query := `
		SELECT 
			user_name, email from users 
		WHERE deleted_at = 0 
	`

	if len(user.UserName) > 0 {
		query += fmt.Sprintf(" and user_name = $%d ", count)
		count++
		arr = append(arr, user.UserName)
	}

	if len(user.Email) > 0 {
		query += fmt.Sprintf(" and email = $%d ", count)
		count++
		arr = append(arr, user.Email)
	}

	var users pb.AllUsers
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var user pb.User

		err = row.Scan(&user.UserName, &user.Email)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}

	return &users, nil
}

func (p *UserStorage) Update(user *pb.User) (*pb.Void, error) {
	query := `
		UPDATE users
		SET user_name = $2, email = $3
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, user.Id, user.UserName, user.Email)
	return nil, err
}

func (p *UserStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE users
		SET deleted_at = $2
		WHERE id = $1
	`

	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *UserStorage) Login(user *pb.User) (*pb.User, error) {
	query := `
		SELECT 
			email 
		FROM users
		WHERE
			user_name = $1 and deleted_at = 0
	`
	row := p.db.QueryRow(query, user.UserName)

	err := row.Scan(&user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
