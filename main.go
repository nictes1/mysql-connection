package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

func connect() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Unable load .env file %s", err.Error())
	}

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	d, err := gorm.Open("mysql", DBURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	db = d
}

func main() {
	connect()
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
