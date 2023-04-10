package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meziaris/devstore/internal/pkg/config"
	"github.com/meziaris/devstore/internal/pkg/db"
)

var cfg config.Config
var DBConn *sqlx.DB

func init() {
	// read config
	loadConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = loadConfig

	// connect database
	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic(err)
	}

	DBConn = db
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
