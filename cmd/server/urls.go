package main

import (
	"github.com/gin-gonic/gin"
	"tigerhallProject/internal/server"
)

// AddKnockKnock Test API
func AddKnockKnock(engine *gin.Engine, serv *server.Server) {
	engine.GET("/knockknock", serv.KnockKnock)

	engine.POST("/api/tiger", serv.AddNewTigerView)
	engine.GET("/api/tigers", serv.GetAllTigers)

	engine.POST("/api/tiger-sighting/:tigerId", serv.AddNewTigerSightingView)
	engine.GET("/api/tiger-sighting/:tigerId", serv.GetAllTigerSightings)

}
