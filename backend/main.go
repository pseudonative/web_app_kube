package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

const (
	host     = "terraform-20231206172344785100000010.cylldswjreil.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "subotai"
	password = "thefourwinds"
	dbname   = "tulsadoom"
)
