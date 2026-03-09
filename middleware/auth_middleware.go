package middleware

import "github.com/gin-gonic/gin"

var SECRET_KEY = []byte("supersecretkey")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context){
		authHeader := g.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnAutherized, gin.H{
				"error": "no token",
			})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		
	}
}
