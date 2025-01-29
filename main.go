package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/saaicasm/windows/handlers"
)

func main() {
	r := gin.Default()

	// Add custom template functions
	r.SetFuncMap(template.FuncMap{
		"seq": func(n int) []int {
			result := make([]int, n)
			for i := range result {
				result[i] = i + 1
			}
			return result
		},
	})

	// Load templates
	r.LoadHTMLGlob("templates/*.html")

	// Serve static files
	r.Static("/static", "./static")

	// Routes
	r.GET("/", handlers.IndexHandler)
	r.POST("/building", handlers.CreateBuildingHandler)
	r.POST("/windows", handlers.SaveWindowsHandler)
	r.GET("/summary", handlers.SummaryHandler)

	// Start the server
	r.Run(":8080")
}
