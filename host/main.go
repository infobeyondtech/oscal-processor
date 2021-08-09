package main

import (
    "fmt"
    "github.com/infobeyondtech/oscal-processor/models/component_value"
    "net/http"
    "strconv"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    //"encoding/json"
    sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
    "github.com/infobeyondtech/oscal-processor/context"
    "github.com/infobeyondtech/oscal-processor/models/component_user_party"
    "github.com/infobeyondtech/oscal-processor/models/control"
    data_models "github.com/infobeyondtech/oscal-processor/models/data_models"
    "github.com/infobeyondtech/oscal-processor/models/information"
    "github.com/infobeyondtech/oscal-processor/models/inventory_item_component"
    "github.com/infobeyondtech/oscal-processor/models/inventory_item_user_party"
    "github.com/infobeyondtech/oscal-processor/models/param_value"
    "github.com/infobeyondtech/oscal-processor/models/profile"
    "github.com/infobeyondtech/oscal-processor/models/profile_navigator"
    sspEngine "github.com/infobeyondtech/oscal-processor/models/ssp"
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
        var json data_models.CreatProfileRequest
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
    r.GET("/profile/navigator/:uuid", func(c *gin.Context) {

        // Load the profile
        iid := c.Param("uuid")
        dir := context.DownloadDir
        inputSrc := dir + "/" + iid
        inputSrc = inputSrc + ".xml"
        inputSrc = context.ExpandPath(inputSrc)

        p := &profile.Profile{}
        profile.LoadFromFile(p, inputSrc)

        // Get the profile's navigator
        pn := &profile_navigator.ProfileNavigator{}
        profile_navigator.CreateProfileNavigator(pn, p)

        fmt.Println(pn.Groups)

        c.JSON(http.StatusOK, pn.Groups)

    })
    // Get Control
    r.GET("/control/:controlid", func(c *gin.Context) {
        id := c.Param("controlid")
        ctrl := control.GetControl(id, false)
        c.JSON(http.StatusOK, ctrl)
    })
    r.GET("/control_enhancement/:enhid", func(c *gin.Context) {
        id := c.Param("enhid")
        ctrl := control.GetControl(id, true)
        c.JSON(http.StatusOK, ctrl)
    })
    // Get ParamValue
    r.GET("/getparam/:record_id", func(c *gin.Context) {
        recordId, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        pv := param_value.GetParamValue(recordId)
        c.JSON(http.StatusOK, pv)
    })
    // Set ParamValue
    r.POST("/setparam/:record_id/:value", func(c *gin.Context) {
        // TODO: Does this need to set the parameter in either profile or
        // the profile's implementation?
        record_id, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        value := c.Param("value")
        pv := param_value.UpdateParamValue(record_id, value)
        c.JSON(http.StatusOK, pv)
    })

    r.POST("/createparam/:project_id/:component_id/:param_id/:value", func(c *gin.Context) {
        // TODO: Does this need to set the parameter in either profile or
        // the profile's implementation?
        project_id, err := strconv.Atoi(c.Param("project_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        component_id := c.Param("component_id")
        param_id := c.Param("param_id")
        value := c.Param("value")
        record_id := param_value.CreateParamValue(project_id, component_id, param_id, value)
        c.JSON(http.StatusOK, strconv.Itoa(record_id))
    })

    r.POST("/deleteparam/:record_id", func(c *gin.Context) {
        record_id, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        param_value.DeleteParamValue(record_id)
    })

    // Get ParamValue
    r.GET("/getcomponent/:record_id", func(c *gin.Context) {
        recordId, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        cv := component_value.GetComponentValue(recordId)
        c.JSON(http.StatusOK, cv)
    })

    // Set ParamValue
    r.POST("/setcomponent/:record_id/:component_id", func(c *gin.Context) {
        // TODO: Does this need to set the parameter in either profile or
        // the profile's implementation?
        record_id, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        component_id := c.Param("component_id")
        cv := component_value.UpdateComponentValue(record_id, component_id)
        c.JSON(http.StatusOK, cv)
    })

    r.POST("/createcomponent/:project_id/:statement_id/:component_id", func(c *gin.Context) {
        // TODO: Does this need to set the parameter in either profile or
        // the profile's implementation?
        project_id, err := strconv.Atoi(c.Param("project_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        statement_id := c.Param("statement_id")
        component_id := c.Param("component_id")
        record_id := component_value.CreateComponentValue(project_id, statement_id, component_id)
        c.JSON(http.StatusOK, strconv.Itoa(record_id))
    })

    r.POST("/deletecomponent/:record_id", func(c *gin.Context) {
        record_id, err := strconv.Atoi(c.Param("record_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        component_value.DeleteComponentValue(record_id)
    })

    r.GET("/get-components/:project_id", func(c *gin.Context) {
        project_id, err := strconv.Atoi(c.Param("project_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        cvs := component_value.GetComponent(project_id)
        c.JSON(http.StatusOK, cvs)
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
    r.POST("/profile/add-party", func(c *gin.Context) {
        var json data_models.AddPartyRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file
        iid := json.UUID
        dir := context.DownloadDir
        inputSrc := dir + "/" + iid + ".xml"
        inputSrc = context.ExpandPath(inputSrc)

        p := &profile.Profile{}
        profile.LoadFromFile(p, inputSrc)

        // add address
        er := profile.AddParty(p, json.OrgName, json.Addresses, json.City, json.State, json.PostalCode, json.RoleId, json.PartyId)
        if er != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
            return
        }

        // give it a new uuid
        p.Id = uuid.New().String()
        profile.WriteToFile(p)

        // return file id
        c.JSON(http.StatusOK, p.Id)

    })
    r.POST("/profile/add-role-party", func(c *gin.Context) {
        var json data_models.AddRolePartyRequest
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

        // give it a new uuid
        p.Id = uuid.New().String()
        profile.WriteToFile(p)

        // return file id
        c.JSON(http.StatusOK, p.Id)
    })
    r.POST("/profile/add-control", func(c *gin.Context) {
        var json data_models.AddControlRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file
        iid := json.UUID
        dir := context.DownloadDir
        inputSrc := dir + "/" + iid + ".xml"
        inputSrc = context.ExpandPath(inputSrc)

        p := &profile.Profile{}
        profile.LoadFromFile(p, inputSrc)

        // add controls
        profile.AddControls(p, json.ControlIDs, "#catalog")

        // give it a new uuid
        p.Id = uuid.New().String()
        profile.WriteToFile(p)

        // return file id
        c.JSON(http.StatusOK, p.Id)
    })
    r.POST("/ssp/create", func(c *gin.Context) {
        var json data_models.SetTitleVersionRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        ssp := &sdk_ssp.SystemSecurityPlan{}
        ssp.Id = uuid.New().String()

        version := json.Version
        oscal_version := json.OscalVersion
        title := json.Title
        profileId := json.ProfileId
        request := data_models.SetTitleVersionRequest{Title: title, Version: version, OscalVersion: oscal_version, ProfileId: profileId}

        // operation
        sspEngine.SetTitleVersion(ssp, request)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/set-characteristic", func(c *gin.Context) {
        var json data_models.AddSystemCharacteristicReuqest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.SetSystemCharacteristic(ssp, json)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/add-inventory-item", func(c *gin.Context) {
        var json data_models.InsertInventoryItemRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.AddInventoryItem(ssp, json)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/add-by-component", func(c *gin.Context) {
        var json data_models.InsertByComponentRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        // todo: add service to sspEngine to add by-component
        
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/remove-by-component", func(c *gin.Context) {
        var json data_models.RemoveByComponentRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.RemoveByComponent(ssp, json.ControlID, json.StatementID, json.ComponentID)
        
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/edit-component-parameter", func(c *gin.Context) {
        var json data_models.EditComponentParameterRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.EditComponentParameter(ssp, json.ControlID, json.StatementID, json.ComponentID, json.ParamID, json.Value)
        
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })

    r.POST("/ssp/add-implemented-requirement", func(c *gin.Context) {
        var json data_models.InsertImplementedRequirementRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.AddImplementedRequirement(ssp, json)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })

    r.GET("/ssp/view-ssp/:fid", func(c *gin.Context) {
        fid := c.Param("fid")
        parent := context.DownloadDir
        targetFile := parent + "/" + fid
        targetFile = context.ExpandPath(targetFile)
        xmlFile := targetFile + ".xml"

        model := sspEngine.MakeSystemSecurityPlanModel(xmlFile)
        c.JSON(http.StatusOK, model)
    })
    r.POST("/ssp/remove-inventory-item", func(c *gin.Context) {
        var json data_models.RemoveElementRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        elementId := json.ElementID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.RemoveInventoryItemAt(ssp, elementId)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.POST("/ssp/remove-implemented-requirement", func(c *gin.Context) {
        var json data_models.RemoveElementRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // load from file, give a new file id
        fileId := json.FileID
        elementId := json.ElementID
        ssp := &sdk_ssp.SystemSecurityPlan{}
        sspEngine.LoadFromFileById(ssp, fileId)
        ssp.Id = uuid.New().String()

        // operation
        sspEngine.RemoveImplementedRequirementAt(ssp, elementId)
        sspEngine.WriteToFile(ssp)

        // return file id
        c.JSON(http.StatusOK, ssp.Id)
    })
    r.GET("/getparaminfo/:paramId", func(c *gin.Context) {
        paramId := c.Param("paramId")
        paramInfo := param_value.GetParamInfo(paramId)
        c.JSON(http.StatusOK, paramInfo)
    })
    r.GET("/infomation/find-component/", func(c *gin.Context) {
        filter := c.Request.URL.Query().Get("filter")
        if len(filter) < 1 {
            component := information.FindComponent("")
            c.JSON(http.StatusOK, component)
        } else {
            component := information.FindComponent(filter)
            c.JSON(http.StatusOK, component)
        }
    })
    r.GET("/information/get-component/:uuid", func(c *gin.Context) {
        id := c.Param("uuid")
        component := information.GetComponent(id)
        c.JSON(http.StatusOK, component)
    })
    r.GET("/infomation/find-inventory-item/", func(c *gin.Context) {
        filter := c.Request.URL.Query().Get("filter")
        if len(filter) < 1 {
            pv := information.FindAllInventoryItem(filter)
            c.JSON(http.StatusOK, pv)
        } else {
            pv := information.FindInventoryItem(filter)
            c.JSON(http.StatusOK, pv)
        }
    })
    r.GET("/infomation/find-party/", func(c *gin.Context) {
        filter := c.Request.URL.Query().Get("filter")
        if len(filter) < 1 {
            pv := information.FindAllParty(filter)
            c.JSON(http.StatusOK, pv)
        } else {
            pv := information.FindParty(filter)
            c.JSON(http.StatusOK, pv)
        }
    })
    r.GET("/infomation/find-user/", func(c *gin.Context) {
        filter := c.Request.URL.Query().Get("filter")
        if len(filter) < 1 {
            pv := information.FindAllUser(filter)
            c.JSON(http.StatusOK, pv)
        } else {
            pv := information.FindUser(filter)
            c.JSON(http.StatusOK, pv)
        }
    })

    r.GET("/get-params/:project_id", func(c *gin.Context) {
        project_id, err := strconv.Atoi(c.Param("project_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        pv := param_value.GetParam(project_id)
        c.JSON(http.StatusOK, pv)
    })

    r.GET("/get-component-user-party/:project_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := component_user_party.GetComponentUserParty(project_id)
        c.JSON(http.StatusOK, result)
    })

    r.POST("/remove-component-user-party/:id", func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        component_user_party.RemoveComponentUserParty(id)
    })

    r.POST("/add-component-user-party/:project_id/:component_id/:user_id/:party_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
        component_id:= c.Param("component_id")
        user_id:= c.Param("user_id")
        party_id:= c.Param("party_id")
       
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := component_user_party.AddComponentUserParty(project_id, component_id, user_id, party_id)
        c.JSON(http.StatusOK, result)
    })

    r.POST("/remove-inventory-component/:id", func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        inventory_item_component.RemoveInventoryItemComponent(id)
    })

    r.POST("/add-inventory-component/:project_id/:inventory_item_id/:component_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
        item_id := c.Param("nventory_item_id")
        component_id := c.Param("component_id")
        
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := inventory_item_component.AddInventoryItemComponent(project_id, item_id, component_id)
        c.JSON(http.StatusOK, result)
    })

    r.GET("/get-inventory-component/:project_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
       
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := inventory_item_component.GetInventoryItemComponent(project_id)
        c.JSON(http.StatusOK, result)
    })

    r.POST("/remove-inventory-user-party/:id", func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        inventory_item_user_party.RemoveInventoryItemUserParty(id)

    })

    r.POST("/add-inventory-user-party/:project_id/:inventory_item_id/:user_id/:party_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
        item_id := c.Param("inventory_item_id")
        party_id := c.Param("party_id")
        user_id := c.Param("user_id")
        
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := inventory_item_user_party.AddInventoryItemUserParty(project_id, item_id, user_id, party_id)
        c.JSON(http.StatusOK, result)
    })

    r.GET("/get-inventory-user-party/:project_id", func(c *gin.Context) {
        project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
       
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        
        result := inventory_item_user_party.GetInventoryItemUserParty(project_id)
        c.JSON(http.StatusOK, result)
    })

    // ToBeTested
    r.GET("/profile/view-profile/:fid", func(c *gin.Context) {
        fid := c.Param("fid")
        parent := context.DownloadDir
        targetFile := parent + "/" + fid
        targetFile = context.ExpandPath(targetFile)
        xmlFile := targetFile + ".xml"
        // println(xmlFile)
        model := profile.MakeProfileModel(xmlFile)
        c.JSON(http.StatusOK, model)
    })

    //r.RunTLS("gamma.infobeyondtech.com:9888", "cert.cert", "cert.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run("0.0.0.0:9050") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
