package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tigerhallProject/internal/models"
	"time"

	//_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"tigerhallProject/common"
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

func MustCreatePgConnection(config common.Config) *gorm.DB {


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Database.Postgres.Host,
		config.Database.Postgres.User,
		config.Database.Postgres.Password,
		config.Database.Postgres.Db,
		5432)


	dbs, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))

	if err != nil {
		fmt.Println("failed setting up postgres connection")
	}
	return dbs
}

func MustCreateServer(config common.Config) *Server {

	resources := ServerResources{
		DB:              MustCreatePgConnection(config),
	}
	return NewServer(resources)
}


func main() {
	configFilePtr := flag.String("config-file", "./configs/local.toml", "Config file to use")
	engine := gin.New()
	config := common.NewConfig(*configFilePtr)

	server := MustCreateServer(config)

	err := server.DB.AutoMigrate(&models.Tiger{})
	err = server.DB.AutoMigrate(&models.TigerSighting{})
	err = server.DB.AutoMigrate(&models.TigerSightingImage{})
	if err != nil {
		fmt.Println("Migration failed")
	} else {
		fmt.Println("Migration Success")
	}
	AddKnockKnock(engine)
	s := &http.Server{
		Addr:           ":" + strconv.FormatUint(config.Server.Port, 10),
		Handler:        engine,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		IdleTimeout:    300 * time.Second,
	}
	fmt.Printf("Server Started %d\n", config.Server.Port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("ListenAndServe:" + err.Error())
	}

}
