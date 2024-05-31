package services

import (
	"context"

	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/repositories"
)

type ProductsService interface {
	GetProducts(ctx context.Context) (*[]models.Products, error)
	GetProduct(ctx context.Context, id uint64) (*models.Products, error)
	AddProduct(ctx context.Context, p *models.Products) error
	UpdateProduct(ctx context.Context, id uint64, p *models.Products) error
	DeleteProduct(ctx context.Context, id uint64) error
}

type productsService struct {
	pr repositories.ProductsRepo
}

func NewProductsService(pr repositories.ProductsRepo) ProductsService {
	return &productsService{
		pr: pr,
	}
}

func (ps *productsService) GetProducts(ctx context.Context) (*[]models.Products, error) {
	return ps.pr.GetProducts(ctx)
}

func (ps *productsService) GetProduct(ctx context.Context, id uint64) (*models.Products, error) {
	return ps.pr.GetProduct(ctx, id)
}

func (ps *productsService) AddProduct(ctx context.Context, p *models.Products) error {
	return ps.pr.AddProduct(ctx, p)
}

func (ps *productsService) UpdateProduct(ctx context.Context, id uint64, p *models.Products) error {
	return ps.pr.UpdateProduct(ctx, id, p)
}

func (ps *productsService) DeleteProduct(ctx context.Context, id uint64) error {
	return ps.pr.DeleteProduct(ctx, id)
}
