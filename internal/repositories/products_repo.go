package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"gorm.io/gorm"
)

type ProductsRepo interface {
	GetProducts(ctx context.Context) (*[]models.Products, error)

	GetProduct(ctx context.Context, id uint64) (*models.Products, error)

	AddProduct(ctx context.Context, product *models.Products) error

	UpdateProduct(ctx context.Context, id uint64, product *models.Products) error

	DeleteProduct(ctx context.Context, id uint64) error
}

type productsRepo struct {
	*gorm.DB
}

func NewProductsRepo(db *gorm.DB) ProductsRepo {
	return &productsRepo{
		DB: db,
	}
}

func (pr *productsRepo) GetProducts(ctx context.Context) (products *[]models.Products, err error) {
	err = pr.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productsRepo) GetProduct(ctx context.Context, id uint64) (product *models.Products, err error) {
	err = pr.DB.First(&product, id).Error
	return
}

func (pr *productsRepo) AddProduct(ctx context.Context, p *models.Products) (err error) {
	email := ctx.Value("email")
	var user models.Users
	err = pr.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}

	p.UserID = user.ID
	return pr.DB.Create(&p).Error
}

func (pr *productsRepo) UpdateProduct(ctx context.Context, id uint64, p *models.Products) error {
	email := ctx.Value("email")
	var user models.Users
	err := pr.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}

	p.UserID = user.ID
	var products *models.Products
	err = pr.DB.First(&products, id).Error
	if err != nil {
		return err
	}
	if products.UserID != user.ID {
		return errors.New("you have no permission to update this product")
	}

	products.Name = p.Name
	products.Price = p.Price
	if err := pr.DB.Save(products).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productsRepo) DeleteProduct(ctx context.Context, id uint64) error {
	email := ctx.Value("email")
	fmt.Println(email)
	var user models.Users
	err := pr.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}

	var products *models.Products
	err = pr.DB.First(&products, id).Error
	if err != nil {
		return err
	}
	if products.UserID != user.ID {
		return errors.New("you have no permission to update this product")
	}
	return pr.DB.Delete(&models.Products{}, id).Error
}
