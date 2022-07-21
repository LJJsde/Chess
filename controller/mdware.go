package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"Chess/util"
	"Chess/module"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")

		if token == "" {
			module.ResponseWithJson(500, "", nil, c)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				module.ResponseWithJson(501, "", nil, c)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				module.ResponseWithJson(502, "", nil, c)
				c.Abort()
				return
			} else {
				c.Set("ID", claims.ID)
				//c.Set("email",claims.Email)
				c.Next()
			}
		}
	}
}
