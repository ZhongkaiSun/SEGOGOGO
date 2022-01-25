package middleware

import (
	"backend/common"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.Id
		DB := common.GetDB()
		var customer model.Customer
		DB.First(&customer, userId)

		// 用户
		if customer.Username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 用户存在 将user 的信息写入上下文
		c.Set("user", customer)

		c.Next()
	}
}
