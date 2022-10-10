package main

import (
	"log"

	"github.com/b2jant/twiss/twiss_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	config := cors.Default()
	router.Use(config)

	v1 := router.Group("/twissapi/v1")
	{
		twitter_route := v1.Group("/get_twitter_sentiment")
		{
			twitter_route.GET("/", routes.FetchTwitter)
			twitter_route.GET("/test", routes.FetchTwitterTest)
			twitter_route.GET("/:query", routes.FetchTwitterWithGCP)
		}
	}

	router.Run()
}
