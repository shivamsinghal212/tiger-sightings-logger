package main

import "github.com/gin-gonic/gin"
import "tigerhallProject/internal"

// AddKnockKnock Test API
func AddKnockKnock(engine *gin.Engine) {
	engine.GET("/knockknock", server.KnockKnock)
}
