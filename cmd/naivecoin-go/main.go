package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/naivecoin-go/handler"
)

func main() {
	initHTTPServer()
}

func initHTTPServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/blocks", handler.GetBlocks)
	r.POST("/blocks", handler.MineBlock)

	r.GET("/peers", handler.GetPeers)
	r.POST("/peers", handler.AddPeer)

	r.Run()
}
