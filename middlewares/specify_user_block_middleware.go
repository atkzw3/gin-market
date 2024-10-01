package middlewares

import (
	"fmt"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func BlockSpecifyUserBlockMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {

			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("BlockSpecifyUserBlockMiddleware エラー1")
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("BlockSpecifyUserBlockMiddleware エラー2")
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetByToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("BlockSpecifyUserBlockMiddleware エラー3")
			return
		}

		if user.ID == 1 {
			ctx.AbortWithStatus(http.StatusForbidden)
			fmt.Println("BlockSpecifyUserBlockMiddleware StatusForbidden")
			return
		}

		fmt.Println("BlockSpecifyUserBlockMiddleware処理終了 next")
		ctx.Next()
	}
}
