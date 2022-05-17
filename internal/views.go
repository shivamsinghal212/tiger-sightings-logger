package server

import "github.com/gin-gonic/gin"

func KnockKnock(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"success": true,
	})
}