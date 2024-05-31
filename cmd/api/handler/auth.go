package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/helper"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"github.com/ilhamgepe/prakerja-s7/internal/services"
)

type AuthsHandler struct {
	UserService services.UserService
}

func NewAuthsHandler(us services.UserService) *AuthsHandler {
	return &AuthsHandler{
		UserService: us,
	}
}

func (ah *AuthsHandler) Register(ctx *gin.Context) {
	var reqBody models.UserRegisterReq
	if err := ctx.Bind(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	// hasing password
	hashedPassword, err := helper.HashPassword(reqBody.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: "oops! something went wrong",
		})
		return
	}

	reqBody.Password = hashedPassword

	if err := ah.UserService.Register(ctx, &reqBody); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, models.APIResponse{
				Status:  "conflict",
				Message: "email already exists",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: "oops! something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "user created",
	})
}

func (ah *AuthsHandler) Login(ctx *gin.Context) {
	var reqBody models.UserLoginReq
	if err := ctx.Bind(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	// get user by email
	user, err := ah.UserService.GetUserByEmail(ctx, reqBody.Email)
	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, models.APIResponse{
				Status:  "not found",
				Message: "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: "oops! something went wrong",
		})
		return
	}

	// compare password
	if !helper.ComparePassword(user.Password, reqBody.Password) {
		ctx.JSON(http.StatusUnauthorized, models.APIResponse{
			Status:  "unauthorized",
			Message: "invalid credentials",
		})
		return
	}

	// generate token
	token, err := helper.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:  "internal server error",
			Message: "oops! something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "user logged in",
		Data: map[string]any{
			"token": token,
			"user":  user,
		},
	})
}
