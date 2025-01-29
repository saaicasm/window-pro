package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Building struct {
	City     string
	Site     string
	Building string
	Windows  []Window
}

type Window struct {
	Height int
	Width  int
}

var buildings []Building

func IndexHandler(c *gin.Context) {
	log.Println("Rendering index page")
	// c.HTML(http.StatusOK, "index.tmpl", nil)
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title": "Windows Tracking Log",
	})
}

func CreateBuildingHandler(c *gin.Context) {
	city := c.PostForm("city")
	site := c.PostForm("site")
	building := c.PostForm("building")
	numWindows, _ := strconv.Atoi(c.PostForm("num_windows"))

	log.Printf("Creating building: City=%s, Site=%s, Building=%s, Windows=%d\n", city, site, building, numWindows)

	c.HTML(http.StatusOK, "windows.tmpl", gin.H{
		"city":       city,
		"site":       site,
		"building":   building,
		"numWindows": numWindows,
	})
}

func SaveWindowsHandler(c *gin.Context) {
	city := c.PostForm("city")
	site := c.PostForm("site")
	building := c.PostForm("building")
	numWindows, _ := strconv.Atoi(c.PostForm("num_windows"))

	log.Printf("Saving windows for: City=%s, Site=%s, Building=%s, Windows=%d\n", city, site, building, numWindows)

	var windows []Window
	for i := 0; i < numWindows; i++ {
		height, _ := strconv.Atoi(c.PostForm("heights[]"))
		width, _ := strconv.Atoi(c.PostForm("widths[]"))
		windows = append(windows, Window{Height: height, Width: width})
	}

	buildings = append(buildings, Building{
		City:     city,
		Site:     site,
		Building: building,
		Windows:  windows,
	})

	c.Redirect(http.StatusSeeOther, "/summary")
}

func SummaryHandler(c *gin.Context) {
	log.Printf("Rendering summary page with %d buildings\n", len(buildings))
	c.HTML(http.StatusOK, "summary.tmpl", gin.H{
		"buildings": buildings,
	})
}
