package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend-onboard-task/helpers"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "토큰 정보가 없습니다."})
            c.Abort()
            return
        }

        userID, userEmail, err := helpers.GetUserIDEndEmailFromToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("userID", userID)
        c.Set("userEmail", userEmail)

        c.Next()
    }
}