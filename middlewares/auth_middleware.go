package middlewares

import (
	"net/http"
	"strings"

	"go-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    h := c.GetHeader("Authorization")
    parts := strings.Split(h, "Bearer ")
    if len(parts)!=2 { c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"}); return }
    tok, err := utils.ParseToken(parts[1])
    if err!=nil || !tok.Valid {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"}); return
    }
    cl := tok.Claims.(jwt.MapClaims)
    uid := uint(cl["user_id"].(float64))
    c.Set("userID", uid)
    c.Next()
  }
}
