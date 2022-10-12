package main

import (
	"ass-02/configs"
	"ass-02/controllers"

	"github.com/gin-gonic/gin"
)


func main() {
	db:= configs.InitDb()
	inDB   := &controllers.InDB{DB : db}

	r := gin.Default()
	
	r.POST("/order", inDB.Post)
	r.PUT("/order/:orderId", inDB.Update)
	r.GET("/orders", inDB.All)
	r.DELETE("/order/:id", inDB.Delete)

	r.Run(":3000")
}