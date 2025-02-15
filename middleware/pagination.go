package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
		offset := (page - 1) * limit
		c.Set("offset", offset)
		c.Set("limit", limit)
		c.Next()
	}
}
