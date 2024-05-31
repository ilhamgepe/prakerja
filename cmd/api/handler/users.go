package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/services"
)

type UsersHandler struct {
	UserService services.UserService
}

func NewUsersHandler(us services.UserService) *UsersHandler {
	return &UsersHandler{
		UserService: us,
	}
}

func (uh *UsersHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.UserService.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.JSON(http.StatusOK, users)
}
