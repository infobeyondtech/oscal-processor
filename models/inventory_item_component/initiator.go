package inventory_item_component

import (
    //"encoding/json"
    "database/sql"
    "strconv"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/infobeyondtech/oscal-processor/context"
)

type InventoryItemComponentMap struct {
    Id int64 `json:"id,omitempty"`
    ProjectId  int64 `json:"projectId,omitempty"`
    InventoryItemId string `json:"inventoryItemId,omitempty"`
    ComponentId string `json:"componentId,omitempty"`
} 

type InventoryValue struct {
    Id int64 `json:"id,omitempty"`
    ProjectId  int64 `json:"projectId,omitempty"`
    InventoryItemId string `json:"inventoryItemId,omitempty"`
} 

type NullableInventoryValue struct {
    Id sql.NullInt64 `json:"id,omitempty"`
    ProjectId  sql.NullInt64 `json:"projectId,omitempty"`
    InventoryItemId sql.NullString `json:"inventoryItemId,omitempty"`
}

func AddInventoryItemComponent(projectid int64, inventoryItemId string, componentid string) int64{
    // do not check duplicate records, client can guard it
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    query := `INSERT INTO inventory_item_component (project_id, inventory_item_id, component_id) Values("`
    query +=  strconv.FormatInt(projectid, 10)
    query += `", "`
    query += inventoryItemId
    query += `", "`
    query += componentid
    query += `")`
    res, err := db.Exec(query)
    if err != nil {
        panic(err.Error())
    }
    id, err := res.LastInsertId()
    return id
}

func RemoveInventoryItemComponent(id int64){
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    query := `DELETE FROM inventory_item_component WHERE id ='`
    query += strconv.FormatInt(id, 10)+ `'`

    _, er := db.Exec(query)
    if er != nil {
        panic(err.Error())
    }
}

func GetInventoryItemComponent(projectid int64) []InventoryItemComponentMap {
    var icmap []InventoryItemComponentMap

    db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // use project id to search for user-party maps
    queryString := `SELECT id, project_id, inventory_item_id, component_id FROM inventory_item_component WHERE project_id = '`
    queryString += strconv.FormatInt(projectid, 10)+ `'`
    queryResult, err := db.Query(queryString)
    if err != nil {
        panic(err.Error())
    }

    // iterate over the query result
    for queryResult.Next() {
        var id int64
        var project_id int64
        var inventory_item_id string
        var component_id string
        err := queryResult.Scan(&id, &project_id, &inventory_item_id, &component_id)
        if err != nil {
            log.Fatal(err)
        }
        icmap = append(icmap, InventoryItemComponentMap{Id:id, ProjectId: project_id, InventoryItemId: inventory_item_id, ComponentId: component_id})
    }

    return icmap
}

func GetInventoryItems(projectId int) []InventoryValue {
    results := make([]InventoryValue, 0)
    db, err := sql.Open("mysql", context.DBSource)
    if err != nil { 
        panic(err.Error())
    }
    defer db.Close()
    qs := `SELECT id, project_id, inventory_item_id ` +
          `FROM inventory_item_component` +
          `WHERE project_id="` + strconv.Itoa(projectId) + `";`
    rows, err := db.Query(qs)
    if err != nil {
        panic(err.Error())
    } 

    defer rows.Close()
    for rows.Next() {
        var result InventoryValue
        var nullableResult NullableInventoryValue
        err := rows.Scan(&nullableResult.Id, &nullableResult.ProjectId, &nullableResult.InventoryItemId)
        if err != nil {
            panic(err.Error())
        }
        if nullableResult.Id.Valid {
            result.Id = nullableResult.Id.Int64
        } else {
            result.Id = -1
        }
        if nullableResult.ProjectId.Valid {
            result.ProjectId = nullableResult.ProjectId.Int64
        } else {
            result.ProjectId = -1
        }
        if nullableResult.InventoryItemId.Valid {
            result.InventoryItemId = nullableResult.InventoryItemId.String
        } else {
            result.InventoryItemId = ""
        }

        results = append(results, result)
    }
    return results
}

