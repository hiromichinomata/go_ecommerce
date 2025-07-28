package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hiromichinomata/go_ecommerce/controllers"
	"github.com/hiromichinomata/go_ecommerce/database"
	"github.com/hiromichinomata/go_ecommerce/middleware"
	"github.com/hiromichinomata/go_ecommerce/routes"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoute(router)
	router.Use(middleware.Authentication())

	router.Get("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.removeitem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
