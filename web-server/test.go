package main

import "github.com/gin-gonic/gin"
import "net/http"
import (
	"os"
	"io"
	"log"
)




func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.tmpl")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "The title is working",
		})
		f.Write(["GET /index writes to file"])
		log.Println("GET /index")
	})
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	r.StaticFile("/static.html", "./static.html")

	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
