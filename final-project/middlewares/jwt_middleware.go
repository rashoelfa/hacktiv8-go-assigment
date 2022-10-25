package middlewares

import (
	"github.com/gin-gonic/gin"
	"mygram/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := utils.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("user", verifyToken)
		c.Next()
	}
}
