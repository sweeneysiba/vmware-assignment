package main

import (
	"fmt"
	"log"
	"os"
	"vmware/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		// c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		// c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// if c.Request.Method == "OPTIONS" {
		// 	fmt.Println("OPTIONS")
		// 	c.AbortWithStatus(200)
		// } else {
		// 	c.Next()
		// }
	}
}

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// can implement a auth token
		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	{
		assignment := new(controllers.AssignmentController)
		v1.POST("/assignemt/get", assignment.Get)
	}

	port := os.Getenv("PORT")
	fmt.Println("PORT", port)
	r.Run(":" + port)
}
