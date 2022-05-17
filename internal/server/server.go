package server

import (
	"gorm.io/gorm"
	"net/http"
)

type ServerResources struct {
	DB     *gorm.DB
	Writer http.ResponseWriter
}

type Server struct {
	ServerResources
}

func NewServer(res ServerResources) *Server {
	return &Server{
		ServerResources: res,
	}
}
