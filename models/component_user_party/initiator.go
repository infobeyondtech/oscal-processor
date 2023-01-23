package component_user_party

import (
    //"encoding/json"
    "database/sql"
    "log"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
    "github.com/infobeyondtech/oscal-processor/context"
)

type ComponentUserPartyMap struct {
    Id int64 `json:"id,omitempty"`
    ProjectId  int64 `json:"projectId,omitempty"`
    ComponentId  string `json:"componentId,omitempty"`
    UserId  string `json:"userId,omitempty"`
    PartyId  string `json:"partyId,omitempty"`
}

func GetComponentUserParty(projectid int64) []ComponentUserPartyMap {
    var userPartyMap []ComponentUserPartyMap

    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // use project id to search for user-party maps
    queryString := `SELECT id, project_id, component_id, user_id, party_id FROM component_user_party WHERE project_id = '`
    queryString += strconv.FormatInt(projectid, 10)+ `'`
    queryResult, err := db.Query(queryString)
    if err != nil {
        panic(err.Error())
    }
    userPartyMap = make([]ComponentUserPartyMap, 0)

    // iterate over the query result
    for queryResult.Next() {
        var id int64
        var project_id int64
        var component_id string
        var user_id string
        var party_id string
        err := queryResult.Scan(&id, &project_id, &component_id, &user_id, &party_id)
        if err != nil {
            log.Fatal(err)
        }
        userPartyMap = append(userPartyMap, ComponentUserPartyMap{Id:id, ProjectId: project_id, ComponentId: component_id, UserId: user_id, PartyId: party_id})
    }

    return userPartyMap
}

func AddComponentUserParty(projectId int64, componentId string, userId string, partyId string) int64{
    // do not check duplicate records, client can guard it
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // double quote might not be necessary
    query := `INSERT INTO component_user_party (project_id, component_id, user_id, party_id) Values("`
    query +=  strconv.FormatInt(projectId, 10)
    query += `", "`
    query += componentId
    query += `", "`
    query += userId
    query += `", "`
    query += partyId
    query += `")`
    res, err := db.Exec(query)
    if err != nil {
        panic(err.Error())
    }
    id, err := res.LastInsertId()
    return id
}

func RemoveComponentUserParty(id int64){
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    query := `DELETE FROM component_user_party WHERE id ='`
    query += strconv.FormatInt(id, 10)+ `'`

    _, er := db.Exec(query)
    if er != nil {
        panic(err.Error())
    }
}

func GetPartys(projectid int64, component_id string, user_id string) []string {

    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // use project id to search for user-party maps
    queryString := `SELECT party_id FROM component_user_party WHERE project_id = '`
    queryString += strconv.FormatInt(projectid, 10)+ `'AND component_id = '`
    queryString += component_id+ `'AND user_id = '`
    queryString += user_id + `'`
    queryResult, err := db.Query(queryString)
    if err != nil {
        panic(err.Error())
    }
    party := make([]string, 0)

    // iterate over the query result
    for queryResult.Next() {
        var party_id string
        err := queryResult.Scan(&party_id)
        if err != nil {
            log.Fatal(err)
        }
        party = append(party, party_id)
    } 

    return party
}

func RemoveParty(projectid int64, component_id string, user_id string, party_id string) {

    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // use project id to search for user-party maps
    queryString := `DELETE FROM component_user_party WHERE project_id = '`
    queryString += strconv.FormatInt(projectid, 10)+ `'AND component_id = '`
    queryString += component_id + `'AND user_id = '`
    queryString += user_id + `'AND party_id = '` 
    queryString += party_id + `'`
    db.Query(queryString)
    _, er := db.Exec(queryString)
    if er != nil {
        panic(err.Error())
    }
}




