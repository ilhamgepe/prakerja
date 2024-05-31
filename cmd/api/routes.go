package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/cmd/api/handler"
	"github.com/ilhamgepe/prakerja-s7/cmd/api/middleware"
	"github.com/ilhamgepe/prakerja-s7/internal/repositories"
	"github.com/ilhamgepe/prakerja-s7/internal/services"
	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// init depedency
	pr := repositories.NewProductsRepo(db)
	ps := services.NewProductsService(pr)

	ur := repositories.NewUserRepo(db)
	us := services.NewUsersService(ur)

	// init handler
	productHandler := handler.NewProductsHandler(ps)
	userHandler := handler.NewUsersHandler(us)
	authHandler := handler.NewAuthsHandler(us)

	products := r.Group("/products")
	{
		// middleware
		products.Use(middleware.WithAuth())
		products.GET("/", productHandler.GetProducts)

		products.GET("/:id", productHandler.GetProduct)

		products.POST("/", productHandler.AddProduct)

		products.PUT("/:id", productHandler.UpdateProduct)

		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	users := r.Group("/users")
	{
		// middleware
		products.Use(middleware.WithAuth())
		users.GET("/", userHandler.GetUsers)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}
}
