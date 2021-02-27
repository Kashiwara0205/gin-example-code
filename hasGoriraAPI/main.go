package main

import (
	"github.com/gin-gonic/gin"
	"hasGoriraAPI/app/service"
	_"encoding/json"
)

func main() {
	r := gin.Default()
	r.GET("/zoos", func(c *gin.Context) {
		var zooService service.ZooService
		zoos := zooService.GetZoos()
		c.JSON(200, zoos)
	})
	r.Run(":5050")
}