package routes

import (
	controller "restaurant_management/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:id", controller.GetOrderItemsByOrderId())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:id", controller.UpdateOrderItem())
}
