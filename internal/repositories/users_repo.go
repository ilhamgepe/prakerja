package repositories

import (
	"context"

	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers(ctx context.Context) (*[]models.Users, error)
	GetUser(ctx context.Context, id uint64) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	AddUser(ctx context.Context, user *models.Users) error
	UpdateUser(ctx context.Context, id uint64, user *models.Users) error
	DeleteUser(ctx context.Context, id uint64) error
	Register(ctx context.Context, user *models.UserRegisterReq) error
	// Login(user *models.UserLoginReq) (*models.Users, error)
}

type userRepo struct {
	*gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

// func (pr *userRepo) Login(user *models.UserLoginReq) (*models.Users, error) {
// 	var u *models.Users
// 	err := pr.DB.Where("email = ?", user.Email).First(&u).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u, nil
// }

func (pr *userRepo) Register(ctx context.Context, user *models.UserRegisterReq) error {
	var u models.Users
	u.Email = user.Email
	u.Username = user.Username
	u.Password = user.Password

	return pr.DB.Create(&u).Error
}

func (pr *userRepo) GetUsers(ctx context.Context) (users *[]models.Users, err error) {
	err = pr.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (pr *userRepo) GetUser(ctx context.Context, id uint64) (user *models.Users, err error) {
	err = pr.DB.First(&user, id).Error
	return
}

func (pr *userRepo) GetUserByEmail(ctx context.Context, email string) (user *models.Users, err error) {
	err = pr.DB.Where("email = ?", email).First(&user).Error
	return
}

func (pr *userRepo) AddUser(ctx context.Context, user *models.Users) (err error) {
	return pr.DB.Create(&user).Error
}

func (pr *userRepo) UpdateUser(ctx context.Context, id uint64, user *models.Users) error {
	var u *models.Users
	err := pr.DB.First(&u, id).Error
	if err != nil {
		return err
	}
	u.Username = user.Username
	u.Email = user.Email
	if err := pr.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (pr *userRepo) DeleteUser(ctx context.Context, id uint64) error {
	var user *models.Users
	err := pr.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	return pr.DB.Delete(&user).Error
}
