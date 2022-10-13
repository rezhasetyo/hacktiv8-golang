package routers

import (
	ctrl "assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router gin.Engine) *gin.Engine {
	router.POST("/orders", ctrl.PostOrder)
	router.GET("/orders", ctrl.GetOrders)
	router.PUT("/orders/:orderId", ctrl.PutOrder)
	router.DELETE("/orders/:orderId", ctrl.DeleteOrder)

	return &router
}
