package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/handl3r/GoForum/api/auth"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	errList := make(map[string]string)
	return func(context *gin.Context) {
		err := auth.TokenValid(context.Request)
		if err != nil {
			errList["unauthorized"] = "unauthorized"
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTION, GET, PUT, PATCH, DELETE")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	}

}
