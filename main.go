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

		plainTextContent := fmt.Sprintf(
			"New Contact Form Submission\n\n"+
				"Name: %s %s\n"+
				"Email: %s\n"+
				"Company: %s\n"+
				"Service: %s\n\n"+
				"Message:\n%s",
			form.FirstName, form.LastName, form.Email, form.Company, form.Service, form.Message,
		)
		htmlContent := fmt.Sprintf(`
			<div style="font-family: sans-serif; max-width: 600px; margin: 0 auto;">
				<h2 style="color: #0a0a0a; border-bottom: 2px solid #0ea5e9; padding-bottom: 8px;">New Contact Form Submission</h2>
				<table style="width: 100%%; border-collapse: collapse; margin-top: 16px;">
					<tr>
						<td style="padding: 8px 12px; font-weight: bold; color: #555; width: 100px;">Name</td>
						<td style="padding: 8px 12px;">%s %s</td>
					</tr>
					<tr style="background-color: #f9fafb;">
						<td style="padding: 8px 12px; font-weight: bold; color: #555;">Email</td>
						<td style="padding: 8px 12px;"><a href="mailto:%s">%s</a></td>
					</tr>
					<tr>
						<td style="padding: 8px 12px; font-weight: bold; color: #555;">Company</td>
						<td style="padding: 8px 12px;">%s</td>
					</tr>
					<tr style="background-color: #f9fafb;">
						<td style="padding: 8px 12px; font-weight: bold; color: #555;">Service</td>
						<td style="padding: 8px 12px;">%s</td>
					</tr>
				</table>
				<div style="margin-top: 24px;">
					<h3 style="color: #555; margin-bottom: 8px;">Message</h3>
					<p style="background-color: #f9fafb; padding: 16px; border-radius: 8px; line-height: 1.6;">%s</p>
				</div>
			</div>`,
			form.FirstName, form.LastName, form.Email, form.Email, form.Company, form.Service, form.Message,
		)
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
