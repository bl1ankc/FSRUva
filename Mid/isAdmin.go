package Mid

import (
	"github.com/gin-gonic/gin"
	"main/Model"
)

// VerifyAdmin 验证管理员
func VerifyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		IsAdmin := c.MustGet("admin").(bool)

		if IsAdmin == false {
			c.JSON(401, gin.H{"message": "非管理员非法操作"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// VerifySuperAdmin 验证超级管理员
func VerifySuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(Model.User)

		if user.AdminType != 1 {
			c.JSON(401, gin.H{"msg": "非老师操作"})
			c.Abort()
			return
		}

		c.Next()
	}
}
