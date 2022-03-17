package endpoint

import "context"

type userUseCase interface {
	Add(ctx context.Context, userID string, age int) error
}
