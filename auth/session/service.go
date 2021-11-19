package session

import (
	"github.com/ITA-Dnipro/Dp-218_Go/repository"
	"github.com/gorilla/sessions"
)

const sessionName = "login"

type AuthService struct {
	DB        repository.UserRepo
	sessStore sessions.Store
}

func NewAuthService(db repository.UserRepo, store sessions.Store) *AuthService {
	return &AuthService{
		DB:        db,
		sessStore: store,
	}
}

// func (sv *AuthService) SSignUp(ctx context.Context, user model.User) error {
// 	if _, err := sv.DB.GetByEmail(ctx, user.Email); err != nil {
// 		return ErrSignIn
// 	}

// 	err := user.HashPassword()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = sv.DB.Create(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	// nuser.Sanitize()

// 	u, err := sv.DB.GetByEmail(ctx, user.Email)
// 	fmt.Println("signed up user: ", u)
// }
