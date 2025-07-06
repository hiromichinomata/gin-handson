package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Asia/Tokyo
	time.Local = time.FixedZone("Local", 9*60*60)

	// dotenv
	err := dotenvutil.Init()
	if err != nil {
		log.Fatalf("dotenv error: %+v\n", err)
	}

	// server init
	err = configutil.ServerInit()
	if err != nil {
		log.Fatalf("server error: %+v\n", err)
	}

	// db init
	err := configutil.DBInit()
	if err != nil {
		log.Fatalf("db error: %+v\n", err)
	}
	defer configutil.DBClose()


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
