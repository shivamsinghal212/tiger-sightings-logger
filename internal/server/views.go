package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"tigerhallProject/internal/Utilities"
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

type TigerSightingPayload struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	LastSeen  uint    `json:"last_seen" binding:"required"`
	FileName  string  `json:"file_name"`
}

func (s *Server) AddNewTigerSightingView(c *gin.Context) {
	tigerId := c.Param("tigerId")
	tigerIdInt, err := strconv.Atoi(tigerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var resizedFileName string
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		resizedFileName = Utilities.ResizeImageUtil(file.Filename)
		e := os.Remove(file.Filename)
		if e != nil {
			log.Fatal(e)
		}
		break
	}
	var latitude float64
	if latitude, err = strconv.ParseFloat(c.Request.PostFormValue("latitude"), 64); err == nil {
	}
	var longitude float64
	if latitude, err = strconv.ParseFloat(c.Request.PostFormValue("longitude"), 64); err == nil {
	}
	last_seen_int, err := strconv.Atoi(c.Request.PostFormValue("last_seen"))
	body := TigerSightingPayload{Latitude: latitude, Longitude: longitude, LastSeen: uint(last_seen_int),
		FileName: resizedFileName}
	status := services.AddNewTigerSighting(s.DB, uint(tigerIdInt), body.Latitude, body.Longitude, body.LastSeen, body.FileName)
	c.JSON(http.StatusOK, gin.H{"status": status})

}

func (s *Server) GetAllTigers(c *gin.Context) {
	tigers := services.GetAllTigers(s.DB, c)
	c.JSON(http.StatusOK, gin.H{"data": tigers})

}

func (s *Server) GetAllTigerSightings(c *gin.Context) {
	tigerId := c.Param("tigerId")
	tigerIdInt, err := strconv.Atoi(tigerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tigers := services.GetAllSightings(s.DB, c, uint(tigerIdInt))
	c.JSON(http.StatusOK, gin.H{"data": tigers})
}
