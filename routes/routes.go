package routes

import (
	"github.com/KasiditR/ecommerce-go-mongo/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.UserInfo())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.POST("/users/addaddress", controllers.AddAddress())
	incomingRoutes.PUT("/users/edithomeaddress", controllers.EditHomeAddress())
	incomingRoutes.PUT("/users/editworkaddress", controllers.EditWorkAddress())
	incomingRoutes.DELETE("/users/deleteaddress", controllers.DeleteAddress())
}
