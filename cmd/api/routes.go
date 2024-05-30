package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/repositories"
	"github.com/ilhamgepe/prakerja-s7/internal/services"
	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// init depedency
	pr := repositories.NewProductsRepo(db)
	ps := services.NewProductsService(pr)

	// HANDLER GET PRODUCT
	r.GET("/products", func(ctx *gin.Context) {
		products, err := ps.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	})

	// HANDLER GET PRODUCT BY ID
	r.GET("/products/:id", func(ctx *gin.Context) {
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
		product, err := ps.GetProduct(iduint64)
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
	})

	r.POST("/products", func(ctx *gin.Context) {
		var product models.Products
		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Status:  "bad request",
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		err := ps.AddProduct(&product)
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
	})

	r.PUT("/products/:id", func(ctx *gin.Context) {
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
		err = ps.UpdateProduct(iduint64, &product)
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
	})

	r.DELETE("/products/:id", func(ctx *gin.Context) {
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
		err = ps.DeleteProduct(iduint64)
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
	})
}
