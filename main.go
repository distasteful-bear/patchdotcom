package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ContactForm struct {
	FirstName string `json:"first-name" binding:"required"`
	LastName  string `json:"last-name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Company   string `json:"company"`
	Service   string `json:"service"`
	Message   string `json:"message" binding:"required"`
}

func main() {
	r := gin.Default()

	r.Static("src/static", "./src/static")

	r.StaticFile("/", "src/home.html")
	r.StaticFile("/home", "src/home.html")
	r.StaticFile("/services", "src/services.html")
	r.StaticFile("/contact", "src/contact.html")

	r.POST("/contact", func(c *gin.Context) {
		var form ContactForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Contact form submission: %s %s <%s> - %s", form.FirstName, form.LastName, form.Email, form.Service)

		// TODO: handle the submission (e.g. send email, store in DB)

		c.JSON(http.StatusOK, gin.H{"message": "Thank you for reaching out! We'll be in touch soon."})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
