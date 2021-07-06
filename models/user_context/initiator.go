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

// Given a set of controls, a set of catalogs, and a baseline,
// generate a unique ID, which can be used for the following operations.
func SetUserContext(projectId string, profileFileId string) UserContext {
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