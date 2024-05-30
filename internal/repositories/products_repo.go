package repositories

import (
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"gorm.io/gorm"
)

type ProductsRepo interface {
	GetProducts() (*[]models.Products, error)

	GetProduct(id uint64) (*models.Products, error)

	AddProduct(product *models.Products) error

	UpdateProduct(id uint64, product *models.Products) error

	DeleteProduct(id uint64) error
}

type productsRepo struct {
	*gorm.DB
}

func NewProductsRepo(db *gorm.DB) ProductsRepo {
	return &productsRepo{
		DB: db,
	}
}

func (pr *productsRepo) GetProducts() (products *[]models.Products, err error) {
	err = pr.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productsRepo) GetProduct(id uint64) (product *models.Products, err error) {
	err = pr.DB.First(&product, id).Error
	return
}

func (pr *productsRepo) AddProduct(p *models.Products) (err error) {
	return pr.DB.Create(&p).Error
}

func (pr *productsRepo) UpdateProduct(id uint64, p *models.Products) error {
	var products *models.Products
	err := pr.DB.First(&products, id).Error
	if err != nil {
		return err
	}
	products.Name = p.Name
	products.Price = p.Price
	if err := pr.DB.Save(products).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productsRepo) DeleteProduct(id uint64) error {
	return pr.DB.Delete(&models.Products{}, id).Error
}
