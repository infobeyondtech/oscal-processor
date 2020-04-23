package main

import (
	"http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/infobeyondtech/oscal-processor/context"
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
		dir := context.UploadDir
		id := uuid.New().String()
		dst := dir + "/" + id
		dst = context.ExpandPath(dst)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK,
			gin.H{
				"uuid":     id,
				"filename": file.Filename,
			})
	})
	// Download
	router.GET("/file/:uuid", func(c *gin.Context) {
		id := c.Param("uuid")
		dir := context.DownloadDir
		src := dir + "/" + id
		src = context.ExpandPath(src)
		c.File(src)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
