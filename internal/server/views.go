package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tigerhallProject/internal/services"
)

func (s *Server) KnockKnock(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"success": true,
	})
}

type TigerPayload struct {
	Name      string  `json:"name" binding:"required"`
	Dob       string  `json:"dob" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	LastSeen  uint    `json:"last_seen" binding:"required"`
}

func (s *Server) AddNewTigerView(c *gin.Context) {
	body := TigerPayload{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	status, tigerObj := services.AddNewTiger(s.DB, body.Name, body.Dob, body.Latitude, body.Longitude, body.LastSeen)
	c.JSON(http.StatusOK, gin.H{"status": status, "tiger": tigerObj})
}
