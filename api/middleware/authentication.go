package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gitlab.udevs.io/eld/eld_go_api_gateway/api/models"
	"gitlab.udevs.io/eld/eld_go_api_gateway/config"
	"gitlab.udevs.io/eld/eld_go_api_gateway/pkg/jwt"
)

func Authentication(handlerOptions *v1.HandlerV1Options, cfg *config.Config) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		isOpen := config.OpenApis[ctx.FullPath()]
		if isOpen {
			ctx.Next()
			return
		}

		token := ctx.GetHeader("Authorization")
		if token != "" {
			if len(strings.Split(token, " ")) > 1 {
				token = strings.Split(token, " ")[1]
			}

			tokenInfo, err := jwt.JWTExtract(token, cfg.JWTSigningKey)
			if err != nil {
				ctx.JSON(401, models.Response{
					Code:    config.ErrorCodeUnauthorized,
					Message: err.Error(),
				})
				ctx.Abort()
				return
			}

			for key, value := range tokenInfo {
				ctx.Request.Header.Set(key, value)
				key1 := key
				value1 := value
				ctx.Set(key1, value1)
			}
		}

		ctx.Next()
	}
}
