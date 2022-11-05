package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jm61/jwt/controllers"
	"github.com/jm61/jwt/initializers"
	"github.com/jm61/jwt/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDb()
}

//var f = fmt.Printf
//var l = fmt.Println

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"Client IP": c.ClientIP(),
			"message":   `If the client is 127.0.0.1, use the X-Forwarded-For header to deduce the original client IP from the trust-worthy parts of that header.Otherwise, simply return the direct client IP`,
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
