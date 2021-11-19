package repository

import (
	"context"

	"github.com/ITA-Dnipro/Dp-218_Go/model"
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
