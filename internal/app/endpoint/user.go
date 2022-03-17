package endpoint

import (
	"context"
	"log"

	"github.com/kridtinc/go-tdd/proto/userpb"
)

// UserEndpoint endpoint containing user operation
type UserEndpoint struct {
	userUseCase userUseCase
	userpb.UnimplementedUserServer
}

// NewUserEndpoint new endpoint instance
func NewUserEndpoint(userUseCase userUseCase) *UserEndpoint {
	return &UserEndpoint{
		userUseCase: userUseCase,
	}
}

// Add endpoint for add user to system
func (u *UserEndpoint) Add(ctx context.Context, req *userpb.AddRequest) (*userpb.AddResponse, error) {
	log.Println("incoming request user ", req.GetUserId(), req.GetAge())

	err := u.userUseCase.Add(ctx, req.GetUserId(), int(req.GetAge()))
	if err != nil {
		return nil, err
	}

	return &userpb.AddResponse{
		StatusCode: 200,
	}, nil
}

func (u *UserEndpoint) mustEmbedUnimplementedUserServer() {
	log.Fatal("unimplemented")
}
