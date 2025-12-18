package paths

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	filePath := "public/home.html"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", filePath, err)
		c.JSON(500, gin.H{"error": "could not locate file"})
		return
	}

	c.Data(200, "text/html; charset=utf-8", fileContent)
}
