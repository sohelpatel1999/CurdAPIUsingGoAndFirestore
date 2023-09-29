package routes

import (
	"gowithcurd/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/item/:id", controller.GetItem)
	router.GET("/item/", controller.GetAllItem)
	router.POST("/item/", controller.CreateItem)
	router.PUT("/item/:id", controller.UpdateItem)
	router.DELETE("/item/:id", controller.DeleteItem)
	return router
}
