package mapdb

import (
	"Dp218Go/model"
	"context"
	"errors"
)

type UserRepoMap struct {
	db        map[string]*model.User
	valid     map[int]bool
	idcounter int
}

func NewUserRepoMap(db map[string]*model.User) *UserRepoMap {
	return &UserRepoMap{
		db:        make(map[string]*model.User),
		valid:     make(map[int]bool),
		idcounter: 0,
	}
}

func (r *UserRepoMap) Create(ctx context.Context, u model.User) (*model.User, error) {

	if _, ok := r.db[u.Email]; ok {
		return nil, errors.New("user exists")

	}
	r.idcounter++
	user := &u
	user.ID = r.idcounter
	r.db[u.Email] = user
	r.valid[u.ID] = false //  can be omited
	return user, nil
}

func (r *UserRepoMap) GetByID(ctx context.Context, id int) (*model.User, error) {

	return nil, errors.New("not implemented")

}
func (r *UserRepoMap) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := r.db[email]
	if user != nil {
		return user, nil
	}
	return nil, errors.New("user not exists")
}

func (r *UserRepoMap) Update(ctx context.Context, u model.User) (*model.User, error) {
	return nil, errors.New("not implemented")
}
func (r *UserRepoMap) Delete(ctx context.Context, id int) error {
	return errors.New("not implemented")
}

func (r *UserRepoMap) GetN(ctx context.Context, n int) ([]*model.User, int, error) {
	return nil, 0, errors.New("not implemented")
}
