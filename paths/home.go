package paths

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	Html404Page := []byte(`
		<!doctype html>
		<html>
		    <head>
		        <title>Patch Solutions</title>
		        <script src="https://unpkg.com/htmx.org@1.9.10"></script>
		    </head>
		    <body>
		        <h1>500 Page Unavailable</h1>
		        <button hx-get="/ping" hx-swap="outerHTML">Test Ping</button>
		    </body>
		</html>
		`)

	filePath := "html/home.html"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", filePath, err)
		c.Data(500, "text/html; charset=utf-8", Html404Page)
		return
	}

	c.Data(200, "text/html; charset=utf-8", fileContent)
}
