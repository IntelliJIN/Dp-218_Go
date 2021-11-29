package repository

import (
	"Dp218Go/model"
	"context"
)

type UserSpecification string

type UserRepo interface {
	Create(context.Context, model.User) (*model.User, error)
	GetByID(context.Context, int) (*model.User, error)
	GetByEmail(context.Context, string) (*model.User, error)
	Update(context.Context, model.User) (*model.User, error)
	Delete(context.Context, int) error
	// GetN(context.Context, int) ([]*model.User, error)
	// SearchUser(context.Context, UserSpecification) ([]*User, error)
}
