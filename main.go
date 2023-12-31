package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.StaticFile("/hello", "./pages/helloworld.html")
  r.Static("/content", "./content")
  r.LoadHTMLGlob("templates/*")
  r.GET("/templates/index", func(c *gin.Context) {
	  c.HTML(http.StatusOK, "action.tmpl", gin.H{
		  "title": "Welcome to htmx",
	  })
  })  
  r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
