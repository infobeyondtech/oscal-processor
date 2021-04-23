package main

import (
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
    //"encoding/json"

	"github.com/infobeyondtech/oscal-processor/context"
	"github.com/infobeyondtech/oscal-processor/models/profile"
	request_models "github.com/infobeyondtech/oscal-processor/models/requests"
	"github.com/infobeyondtech/oscal-processor/models/profile_navigator"
	"github.com/infobeyondtech/oscal-processor/models/control"
	"github.com/infobeyondtech/oscal-processor/models/param_value"
)

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
		var json request_models.CreatProfileRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a file id
		fid, err := profile.CreateProfile(json.Controls, json.Baseline, json.Catalogs, json.Title, json.OrgUuid, json.OrgName, json.OrgEmail)
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
	// Get Control
	r.POST("/control/:controlid", func(c *gin.Context) {
		id := c.Param("controlid")
        ctrl := control.GetControl(id)
        c.JSON(http.StatusOK, ctrl)
	})
	// Get ParamValue
    r.POST("/getparam/:uuid/:paramid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		paramid := c.Param("paramid")
        pv := param_value.GetParamValue(uuid, paramid)
        c.JSON(http.StatusOK, pv)
	})
	// Set ParamValue
    r.POST("/setparam/:uuid/:paramid/:value", func(c *gin.Context) {
		uuid := c.Param("uuid")
		paramid := c.Param("paramid")
        value := c.Param("value")
        pv := param_value.SetParamValue(uuid, paramid, value)
        c.JSON(http.StatusOK, pv)
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
	//r.POST("/profile/set-title", func(c *gin.Context) {
	//	var json request_models.SetTitleVersionRequest
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}

	//	// load from file
	//	iid := json.UUID
	//	dir := context.DownloadDir
	//	inputSrc := dir + "/" + iid
	//	inputSrc = context.ExpandPath(inputSrc)
	//	p := &profile.Profile{}
	//	profile.LoadFromFile(p, inputSrc)

	//	// Set title and version
	//	er := profile.SetTitleVersion(p, json.Version, json.OscalVersion, json.Title)
	//	if er != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
	//		return
	//	}
	//})
	r.POST("/profile/add-address", func(c *gin.Context) {
		var json request_models.AddAddressRequest
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

		// add address
		er := profile.AddAddress(p, json.Addresses, json.City, json.State, json.PostalCode)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}
	})
	r.POST("/profile/add-role-party", func(c *gin.Context) {
		var json request_models.AddRolePartyRequest
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

		// add role party
		er := profile.AddRoleParty(p, json.RoleID, json.Title, json.PartyID, json.OrgName, json.Email)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}
	})
	//r.RunTLS("gamma.infobeyondtech.com:9888", "cert.cert", "cert.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
