package middlewares

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckToken(ctx *gin.Context) {
	expectedToken := os.Getenv("SECRET")
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.AbortWithError(401, errors.New("unauthorized"))
		return
	}

	authArr := strings.Split(authorization, " ")
	if len(authArr) < 2 {
		ctx.AbortWithError(401, errors.New("unauthorized"))
		return
	}
	token := authArr[1]

	if token != expectedToken {
		ctx.AbortWithError(401, errors.New("unauthorized"))
		return
	}
	ctx.Next()
}
