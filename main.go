package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

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

	r.LoadHTMLGlob("src/**/*")

	r.Static("src/styles", "./src/styles")

	r.GET("/", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "home", gin.H{
	        "message": "success",
	    })
	})
	r.GET("/home", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "home", gin.H{
	        "message": "success",
	    })
	})
	r.GET("/contact", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "contact", gin.H{
	        "message": "success",
	    })
	})
	r.GET("/services", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "services", gin.H{
	        "message": "success",
	    })
	})

	r.POST("/contact", func(c *gin.Context) {
		var form ContactForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Contact form submission: %s %s <%s> - %s", form.FirstName, form.LastName, form.Email, form.Service)

		from := mail.NewEmail("Patch Solutions Contact Form", "james.m@patch-solutions.com")
		to := mail.NewEmail("Information", "james.m@patch-solutions.com")

		subject := "New Contact Form Submission"

		plainTextContent := fmt.Sprintf("A new contact form submission!%0A From:%0A%s %s%0A <%s> - %s.%0ACompany:%0A %s%0A Message:%0A%s", form.FirstName, form.LastName, form.Email, form.Service, form.Company, form.Message)
		htmlContent := fmt.Sprintf("A new contact form submission!%0A From:%0A%s %s%0A <%s> - %s.%0ACompany:%0A %s%0A Message:%0A%s", form.FirstName, form.LastName, form.Email, form.Service, form.Company, form.Message)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

		response, err := client.Send(message)

		if err != nil || response.StatusCode != 202 {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred while sending your message. Please contact us directly with the listed email address."})
		} else {
			fmt.Println(response.StatusCode)
			c.JSON(http.StatusOK, gin.H{"message": "Thank you for reaching out! We'll be in touch soon."})
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
