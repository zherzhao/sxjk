package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsV1() gin.HandlerFunc {
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, UPDATE, DELETE ,GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func CorsV2() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, UPDATE, DELETE ,GET, OPTIONS")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.Header("Access-Control-Allow-Origin", "*")
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
