package usecase

import (
	"context"
	"log"

	"github.com/kridtinc/go-tdd/internal/app/entity"
)

// UserList list of user
var UserList = []entity.User{}

// UserUseCase instance for user business logic
type UserUseCase struct {
}

// NewUserUseCase initiate user usecase instace
func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

// Add function for insert user
func (u *UserUseCase) Add(ctx context.Context, userID string, age int) error {
	UserList = append(UserList, entity.User{
		ID:  userID,
		Age: age,
	})

	log.Printf("now user list is %v\n", UserList)

	return nil
}
