package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/infobeyondtech/oscal-processor/context"
	"github.com/infobeyondtech/oscal-processor/models/profile"
	"github.com/infobeyondtech/oscal-processor/models/profile_navigator"
)

type CreatProfileDTO struct {
	Baseline string   `json:"baseline" binding:"required"`
	Controls []string `json:"controls" binding:"required"`
	Catalogs []string `json:"catalogs" binding:"required"`
}

type SetTitleVersionDTO struct {
	UUID         string `json:"uuid" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Version      string `json:"version" binding:"required"`
	OscalVersion string `json:"oscalversion" binding:"required"`
}

func main() {
	r := gin.Default()
	// Cors support
	r.Use(cors.Default())
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
	// Get Profile Navigator
	r.POST("/profile/navigator/:uuid", func(c *gin.Context) {

		// Load the profile
		iid := c.Param("uuid")
		dir := context.DownloadDir
		inputSrc := dir + "/" + iid
		inputSrc = context.ExpandPath(inputSrc)
		p := &profile.Profile{}
		profile.LoadFromFile(p, inputSrc)

        // Get the profile's navigator
        pn := &profile_navigator.ProfileNavigator{}
        profile_navigator.CreateProfileNavigator(pn, p)

        c.JSON(http.StatusOK, pn.Groups)

	})
	// Resolve profile
	r.POST("/profile/resolve/:uuid", func(c *gin.Context) {
		rules := context.OSCALRepo +
			"/src/utils/util/resolver-pipeline/oscal-profile-RESOLVE.xsl"
		rules = context.ExpandPath(rules)
		jarPath := context.JarLibDir + "/saxon-he-10.0.jar"
		jarPath = context.ExpandPath(jarPath)

		// input path
		iid := c.Param("uuid")
		dir := context.DownloadDir
		inputSrc := dir + "/" + iid
		inputSrc = context.ExpandPath(inputSrc)

		// output path
		output := context.TempDir
		oid := uuid.New().String()
		output = output + "/" + oid + ".xml"
		output = context.ExpandPath(output)

		err := profile.ResolveProfile(jarPath, rules, inputSrc, output)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// Modify Metadata
	r.POST("/profile/set-title", func(c *gin.Context) {
		var json SetTitleVersionDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// load from file
		iid := json.UUID
		dir := context.DownloadDir
		inputSrc := dir + "/" + iid
		inputSrc = context.ExpandPath(inputSrc)
		p := &profile.Profile{}
		profile.LoadFromFile(p, inputSrc)

		// Set title and version
		er := profile.SetTitleVersion(p, json.Version, json.OscalVersion, json.Title)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}
	})
	//r.RunTLS("gamma.infobeyondtech.com:9888", "cert.cert", "cert.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
