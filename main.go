package main

import (
	"log"
	"os"

	"github.com/KasiditR/ecommerce-go-mongo/controllers"
	"github.com/KasiditR/ecommerce-go-mongo/database"
	"github.com/KasiditR/ecommerce-go-mongo/middleware"
	"github.com/KasiditR/ecommerce-go-mongo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use((gin.Logger()))
	routes.UserRoutes(router)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.Authentication())
	{
		authRoutes.GET("/addtocart", app.AddToCart())
		authRoutes.GET("/removeitem", app.RemoveItem())
		authRoutes.GET("/cartcheckout", app.BuyFromCart())
		authRoutes.GET("/instantbuy", app.InstantBuy())
	}

	log.Fatal(router.Run(":" + port))
}
