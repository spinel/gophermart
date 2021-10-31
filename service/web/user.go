package web

import (
	"context"

	"github.com/google/uuid"
	"github.com/spinel/gophermart/model"
	"github.com/spinel/gophermart/store"
	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

// UserWebService ...
type UserWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewUserWebService is a user service
func NewUserWebService(ctx context.Context, store *store.Store) *UserWebService {
	return &UserWebService{
		ctx:   ctx,
		store: store,
	}
}

// Create user service
func (svc UserWebService) Create(ctx context.Context, userRegisterForm model.UserRegisterForm) (*model.User, error) {
	passwordHash, err := hashPassword(userRegisterForm.Password)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Login error")
	}

	user := &model.User{
		UUID:     uuid.New(),
		Login:    userRegisterForm.Login,
		Password: passwordHash,
	}

	user, err = svc.store.User.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Create error")
	}

	return user, nil
}

// Login user signin
func (svc UserWebService) Login(ctx context.Context, userRegisterForm model.UserRegisterForm) (*model.User, error) {
	user, err := svc.store.User.GetByLogin(ctx, userRegisterForm.Login)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Login error")
	}

	if ok := checkPasswordHash(user.Password, userRegisterForm.Password); ok {
		return user, nil
	}

	return nil, errors.Wrap(err, "svc.User.Permission denied")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}
