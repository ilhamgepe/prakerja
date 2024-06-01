package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/services"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductService services.ProductsService
}

func NewProductsHandler(ps services.ProductsService) *ProductHandler {
	return &ProductHandler{
		ProductService: ps,
	}
}

func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := ph.ProductService.GetProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: "something went wrong",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	iduint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: "id must be a number",
			Data:    nil,
		})
		return
	}
	product, err := ph.ProductService.GetProduct(ctx, iduint64)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, models.APIResponse{
				Status:  "not found",
				Message: fmt.Sprintf("product with id %d not found", iduint64),
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": models.APIResponse{
			Status:  "internal server error",
			Message: "id must be a number",
			Data:    nil,
		}})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) AddProduct(ctx *gin.Context) {
	var product models.Products
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.APIResponse{
			Status:  "unauthorized",
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}
	stdCtx := context.WithValue(ctx, "email", email)
	err := ph.ProductService.AddProduct(stdCtx, &product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.APIResponse{
		Status:  "created",
		Message: "product created",
	})
}

func (ph *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	iduint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: "id must be a number",
			Data:    nil,
		})
		return
	}

	var product models.Products
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.APIResponse{
			Status:  "unauthorized",
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}
	stdCtx := context.WithValue(ctx, "email", email)
	err = ph.ProductService.UpdateProduct(stdCtx, iduint64, &product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "product updated",
	})
}

func (ph *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	iduint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: "id must be a number",
			Data:    nil,
		})
		return
	}

	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.APIResponse{
			Status:  "unauthorized",
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}
	stdCtx := context.WithValue(ctx, "email", email)
	err = ph.ProductService.DeleteProduct(stdCtx, iduint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "product deleted",
	})
}
