package Routers

import (
	controllers "belajar-gin/Controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	// inisialisasi router
	router := gin.Default()

	// router API
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)	

	//router API menampilkan semua cars
	router.GET("/cars", controllers.GetAll)
	return router
}