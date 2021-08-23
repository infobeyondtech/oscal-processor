package user_context

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/infobeyondtech/oscal-processor/context"
)

type UserContext struct {
    ProjectId  string `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
    ProfileFileId string `xml:"profilefileid,omitempty" json:"profilefileid,omitempty"`
}

type NullableUserContext struct {
    ProjectId  sql.NullString `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
    ProfileFileId sql.NullString `xml:"profilefileid,omitempty" json:"profilefileid,omitempty"`
}

type UserSsp struct{
    ProjectId string `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
    SspFileId string `xml:"sspFileId,omitempty" json:"sspFileId,omitempty"`
}

type NullableUserSsp struct{
    ProjectId sql.NullString `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
    SspFileId sql.NullString `xml:"sspFileId,omitempty" json:"sspFileId,omitempty"`
}

// Given a set of controls, a set of catalogs, and a baseline,
// generate a unique ID, which can be used for the following operations.
func AddUserContext(projectId string, profileFileId string) UserContext {
        // Open the DB
        db, err := sql.Open("mysql", context.DBSource)
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()
        // TODO: Do we need to error check to make sure the
        //       fileid, and paramid are valid?
        // TODO: Check to see if value already exisits in DB
        //       Update if so
        query := `INSERT INTO user_context(projectId, profileFileId) Values("`
        query += projectId
        query += `", "`
        query += profileFileId
        query += `")`
        _, err = db.Exec(query)
        if err != nil {
            fmt.Println("Caused by: " + query)
            panic(err.Error())
        }

        result := UserContext{projectId, profileFileId}
        return result
}

func GetProfileFileId(projectId string) UserContext {
    var result UserContext
    var nullableResult NullableUserContext
    // Open the DB
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    query := `SELECT projectId, profileFileId From user_context WHERE projectId = "`
    query += projectId
    query += `";`
    err = db.QueryRow(query).
        Scan(&nullableResult.ProjectId,&nullableResult.ProfileFileId)
    if err != nil {
        panic(err.Error())
    }
    // Validate the query response

    if nullableResult.ProjectId.Valid {
        result.ProjectId = nullableResult.ProjectId.String
    } else {
        result.ProjectId = ""
    }
    
    if nullableResult.ProfileFileId.Valid {
        result.ProfileFileId = nullableResult.ProfileFileId.String
    } else {
        result.ProfileFileId = ""
    }

    return result
}

func UpdateUserContext(projectId string, profileFileId string) UserContext{
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    query := `UPDATE user_context Set profileFileId = "`
    query += profileFileId
    query += `" WHERE projectId = "`
    query += projectId
    query += `");`
    _, err = db.Exec(query)
    if err != nil {
        fmt.Println("Caused by: " + query)
        panic(err.Error())
    }

    result := UserContext{projectId, profileFileId}
    return result
}

func SetUserContext(projectId string, profileFileId string) UserContext{
    // one project is linked to one profile file at most
    currentProfile := GetProfileFileId(projectId)
    if currentProfile.ProfileFileId == "" {
        return AddUserContext(projectId, profileFileId)
    }else{
        return UpdateUserContext(projectId, profileFileId)
    }
}



func AddUserSsp(projectId string, sspFileId string) UserSsp {
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    query := `INSERT INTO user_ssp(projectId, sspFileId) Values("`
    query += projectId
    query += `", "`
    query += sspFileId
    query += `")`
    _, err = db.Exec(query)
    if err != nil {
        fmt.Println("Caused by: " + query)
        panic(err.Error())
    }

    result := UserSsp{projectId, sspFileId}
    return result
}

func GetSspFileId(projectId string) UserSsp {
    var result UserSsp
    var nullableResult NullableUserSsp
    
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    query := `SELECT projectId, sspFileId From user_ssp WHERE projectId = "`
    query += projectId
    query += `";`
    err = db.QueryRow(query).
        Scan(&nullableResult.ProjectId, &nullableResult.SspFileId)
    if err != nil {
        panic(err.Error())
    }

    if nullableResult.ProjectId.Valid {
        result.ProjectId = nullableResult.ProjectId.String
    } else {
        result.ProjectId = ""
    }
    
    if nullableResult.SspFileId.Valid {
        result.SspFileId = nullableResult.SspFileId.String
    } else {
        result.SspFileId = ""
    }

    fmt.Println("GetSspFileId Result")
    fmt.Println(result)

    // if no record is found, return an empty record
    return result
}

func UpdateUserSsp(projectId string, sspFileId string) UserSsp{
    // update user ssp record disreguard what the current ssp is
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    query := `UPDATE user_ssp Set sspFileId = "`
    query += sspFileId
    query += `" WHERE projectId = "`
    query += projectId
    query += `";`
    _, err = db.Exec(query)
    if err != nil {
        fmt.Println("Caused by: " + query)
        panic(err.Error())
    }

    result := UserSsp{projectId, sspFileId}
    return result
}

func SetUserSsp(projectId string, sspFileId string) UserSsp{
    // one project is linked to one ssp file at most
    currentSsp := GetSspFileId(projectId)
    if currentSsp.SspFileId == "" {
        fmt.Println("Adding")
        return AddUserSsp(projectId, sspFileId)
    }else{
        fmt.Println("Updating")
        return UpdateUserSsp(projectId, sspFileId)
    }
}




