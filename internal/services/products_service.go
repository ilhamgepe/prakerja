package services

import (
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/repositories"
)

type ProductsService interface {
	GetProducts() (*[]models.Products, error)
	GetProduct(id uint64) (*models.Products, error)
	AddProduct(p *models.Products) error
	UpdateProduct(id uint64, p *models.Products) error
	DeleteProduct(id uint64) error
}

type productsService struct {
	pr repositories.ProductsRepo
}

func NewProductsService(pr repositories.ProductsRepo) ProductsService {
	return &productsService{
		pr: pr,
	}
}

func (ps *productsService) GetProducts() (*[]models.Products, error) {
	return ps.pr.GetProducts()
}

func (ps *productsService) GetProduct(id uint64) (*models.Products, error) {
	return ps.pr.GetProduct(id)
}

func (ps *productsService) AddProduct(p *models.Products) error {
	return ps.pr.AddProduct(p)
}

func (ps *productsService) UpdateProduct(id uint64, p *models.Products) error {
	return ps.pr.UpdateProduct(id, p)
}

func (ps *productsService) DeleteProduct(id uint64) error {
	return ps.pr.DeleteProduct(id)
}
