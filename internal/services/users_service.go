package services

import (
	"context"

	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/repositories"
)

type UserService interface {
	GetUsers(ctx context.Context) (*[]models.Users, error)
	GetUser(ctx context.Context, id uint64) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	AddUser(ctx context.Context, user *models.Users) error
	UpdateUser(ctx context.Context, id uint64, user *models.Users) error
	DeleteUser(ctx context.Context, id uint64) error
	Register(ctx context.Context, user *models.UserRegisterReq) error
	// Login(user *models.UserLoginReq) (*models.Users, error)
}

type userService struct {
	ur repositories.UserRepo
}

func NewUsersService(ur repositories.UserRepo) UserService {
	return &userService{
		ur: ur,
	}
}

// func (us *userService) Login(user *models.UserLoginReq) (*models.Users, error) {
// 	return us.ur.Login(user)
// }

func (us *userService) Register(ctx context.Context, user *models.UserRegisterReq) error {
	return us.ur.Register(ctx, user)
}

func (us *userService) GetUsers(ctx context.Context) (*[]models.Users, error) {
	return us.ur.GetUsers(ctx)
}

func (us *userService) GetUser(ctx context.Context, id uint64) (*models.Users, error) {
	return us.ur.GetUser(ctx, id)
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return us.ur.GetUserByEmail(ctx, email)
}

func (us *userService) AddUser(ctx context.Context, user *models.Users) error {
	return us.ur.AddUser(ctx, user)
}

func (us *userService) UpdateUser(ctx context.Context, id uint64, user *models.Users) error {
	return us.ur.UpdateUser(ctx, id, user)
}

func (us *userService) DeleteUser(ctx context.Context, id uint64) error {
	return us.ur.DeleteUser(ctx, id)
}
