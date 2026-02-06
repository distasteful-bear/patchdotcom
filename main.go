package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("src/static", "./src/static")

	r.StaticFile("/", "src/home.html")
	r.StaticFile("/home", "src/home.html")
	r.StaticFile("/services", "src/services.html")
	r.StaticFile("/contact", "src/contact.html")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
