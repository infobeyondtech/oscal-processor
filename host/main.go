package main

import (
	"fmt"
	"http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	// Ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Upload
	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		// log.Println(file.Filename)

		// Generate a file
		dir := "/home/tom/oscal_processing_space"
		id := uuid.New().String()
		dst := dir + "/output/" + id + ".xml"

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// Download
	router.GET("file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
