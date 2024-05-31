package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/prakerja-s7/helper"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
)

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		auth := strings.Split(authorization, " ")
		if len(auth) != 2 {
			ctx.AbortWithStatusJSON(401, models.APIResponse{
				Status: "unauthorized",
			})
			return
		}

		tokenReq := auth[1]

		// validate token
		token, err := helper.VerifyToken(tokenReq)
		if err != nil {
			ctx.AbortWithStatusJSON(401, models.APIResponse{
				Status: "unauthorized",
			})
			return
		}

		email, err := token.Claims.GetSubject()
		if err != nil {
			ctx.AbortWithStatusJSON(401, models.APIResponse{
				Status: "unauthorized",
			})
			return
		}

		// set email to context
		ctx.Set("email", email)

		ctx.Next()
	}
}
