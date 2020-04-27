package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/infobeyondtech/oscal-processor/context"
	"github.com/infobeyondtech/oscal-processor/models/profile"
)

type CreatProfileDTO struct {
	Baseline string   `json:"baseline" binding:"required"`
	Controls []string `json:"controls" binding:"required"`
	Catalogs []string `json:"catalogs" binding:"required"`
}

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
	r.GET("/download/:uuid", func(c *gin.Context) {
		id := c.Param("uuid")
		// Search the file in three directories
		dir := context.DownloadDir
		src := dir + "/" + id
		src = context.ExpandPath(src)

		c.File(src)
	})
	// Create profile
	r.POST("/profile/create", func(c *gin.Context) {
		var json CreatProfileDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a file
		fid, err := profile.CreateProfile(json.Controls, json.Baseline, json.Catalogs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK,
			gin.H{
				"id":  fid,
				"ext": "xml",
			})

	})

	r.Run("gamma.infobeyondtech.com:9888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
