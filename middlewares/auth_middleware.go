package middlewares

import (
	"fmt"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			// 以降の処理は停止するが、今の処理はそのままなのでreturnする
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("AuthMiddleware エラー1")
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("AuthMiddleware エラー2")
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetByToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("AuthMiddleware エラー3")
			return
		}
		ctx.Set("user", user)

		fmt.Println("処理終了 next")
		ctx.Next()
	}
}
