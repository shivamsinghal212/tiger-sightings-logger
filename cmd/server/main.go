package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tigerhallProject/common"
	"tigerhallProject/internal/models"
	"tigerhallProject/internal/server"
	"time"
)

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

func MustCreateServer(config common.Config) *server.Server {

	resources := server.ServerResources{
		DB: MustCreatePgConnection(config),
	}
	return server.NewServer(resources)
}

func main() {
	configFilePtr := flag.String("config-file", "./configs/local.toml", "Config file to use")
	engine := gin.New()
	config := common.NewConfig(*configFilePtr)

	serv := MustCreateServer(config)

	err := serv.DB.AutoMigrate(&models.Tiger{})
	err = serv.DB.AutoMigrate(&models.TigerSighting{})
	err = serv.DB.AutoMigrate(&models.TigerSightingImage{})
	if err != nil {
		fmt.Println("Migration failed")
	} else {
		fmt.Println("Migration Success")
	}
	AddKnockKnock(engine, serv)
	s := &http.Server{
		Addr:         ":" + strconv.FormatUint(config.Server.Port, 10),
		Handler:      engine,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
		IdleTimeout:  300 * time.Second,
	}
	fmt.Printf("Server Started %d\n", config.Server.Port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("ListenAndServe:" + err.Error())
	}

}
