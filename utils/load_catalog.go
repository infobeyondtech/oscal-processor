package main

import (
    "os"
	"database/sql"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "github.com/docker/oscalkit/types/oscal/catalog"
	_ "github.com/go-sql-driver/mysql"
    "github.com/alediaferia/stackgo"
)

func GetAllControls(c catalog.Catalog) []catalog.Control {

    result := make([]catalog.Control, 0)
    visited := make(map[string]bool)
    control_stack := stackgo.NewStack()

    // Add 0-depth controls from catalog to stack
    for _, group := range c.Groups {
        for _, ctrl := range group.Controls {
            // Skipping controls without an Id
            if ctrl.Id != "" {
                control_stack.Push(ctrl)
            }
        }
    }

    for control_stack.Size() > 0 {
        curr_control := control_stack.Pop().(catalog.Control)
        _, isVisited := visited[curr_control.Id]
        if !isVisited {
            visited[curr_control.Id] = true
            result = append(result, curr_control)
            for _, child_control := range curr_control.Controls {
                // Skipping controls without an Id
                if child_control.Id != "" {
                    _, isVisitedChild := visited[child_control.Id]
                    if !isVisitedChild {
                        control_stack.Push(child_control)
                    } else {
                        fmt.Println("GetAllControls Warning: Trying to add already visited child_control: " + child_control.Id)
                    }
                }
            }
        } else {
            fmt.Println("GetAllControls Warning: Trying to add already visited control: " + curr_control.Id)
        }
    }

    return result

}

func CreatePartsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS parts(partid varchar(20), name varchar(20), PRIMARY KEY(partid))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    visited := make(map[string]bool)

    for _, ctrl := range GetAllControls(c) {
        part_stack := stackgo.NewStack()
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                part_stack.Push(part)
            }
        }
        for part_stack.Size() > 0 {
            curr_part := part_stack.Pop().(catalog.Part)
            _, isVisited := visited[curr_part.Id]
            if !isVisited {
                visited[curr_part.Id] = true
                query := `INSERT INTO parts(partid, name) Values("`
                query += curr_part.Id
                query += `", "`
                query += curr_part.Name
                query += `")`
                _,err = db.Exec(query)
                if err != nil {
                    fmt.Println(err.Error())
                    fmt.Println("Caused by: " + query)
                    os.Exit(0)
                }
                for _, child_part := range curr_part.Parts {
                    if child_part.Id != "" {
                        _, isVisitedChild := visited[child_part.Id]
                        if !isVisitedChild {
                            part_stack.Push(child_part)
                        } else {
                            fmt.Println("CreatePartsTable Warning: Trying to add already visited child_part: " + child_part.Id)
                        }
                    }
                }
            } else {
                fmt.Println("CreatePartsTable Warning: Trying to add already visited curr_part: " + curr_part.Id)
            }
        }
    }
}

func CreateParamsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS params(paramid varchar(20), label varchar(300), PRIMARY KEY(paramid))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    for _, ctrl := range GetAllControls(c) {
        for _, param := range ctrl.Parameters {
            query := `INSERT INTO params(paramid, label) Values("`
            query += param.Id
            query += `", "`
            query += string(param.Label)
            query += `")`
            _,err = db.Exec(query)
            if err != nil {
                fmt.Println(err.Error())
                fmt.Println("Caused by: " + query)
                os.Exit(0)
            }
        }
    }
}

func CreateControlsToPartsTable(db *sql.DB, c catalog.Catalog) {
    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `controls_parts`(controlid varchar(20), partid varchar(20))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    visited := make(map[string]bool)

    for _, ctrl := range GetAllControls(c) {
        part_stack := stackgo.NewStack()
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                part_stack.Push(part)
            }
        }
        for part_stack.Size() > 0 {
            curr_part := part_stack.Pop().(catalog.Part)
            _, isVisited := visited[curr_part.Id]
            if !isVisited {
                visited[curr_part.Id] = true
                query := "INSERT INTO `controls_parts`(controlid, partid) Values(\""
                query += ctrl.Id
                query += `", "`
                query += curr_part.Id
                query += `")`
                _,err = db.Exec(query)
                if err != nil {
                    fmt.Println(err.Error())
                    fmt.Println("Caused by: " + query)
                    os.Exit(0)
                }
                for _, child_part := range curr_part.Parts {
                    if child_part.Id != "" {
                        _, isVisitedChild := visited[child_part.Id]
                        if !isVisitedChild {
                            part_stack.Push(child_part)
                        } else {
                            fmt.Println("CreateControlsToPartsTable Warning: Trying to add already visited child_part: " + child_part.Id)
                        }
                    }
                }
            } else {
                fmt.Println("CreateControlsToPartsTable Warning: Trying to add already visited curr_part: " + curr_part.Id)
            }
        }
    }
}

func CreateControlsToParamsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `controls_params`(controlid varchar(20), paramid varchar(20), PRIMARY KEY(paramid))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    for _, ctrl := range GetAllControls(c) {
        for _, param := range ctrl.Parameters {
            query := "INSERT INTO `controls_params`(controlid, paramid) Values(\""
            query += ctrl.Id
            query += `", "`
            query += param.Id
            query += `")`
            _,err = db.Exec(query)
            if err != nil {
                fmt.Println(err.Error())
                fmt.Println("Caused by: " + query)
                os.Exit(0)
            }
        }
    }
}

func CreatePartsToPartsTable(db *sql.DB, c catalog.Catalog) {
    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `parts_parts`(`parent_partid` varchar(20), `child_partid` varchar(20))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    visited := make(map[string]bool)

    for _, ctrl := range GetAllControls(c) {
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                part_stack := stackgo.NewStack()
                part_stack.Push(part)
                for part_stack.Size() > 0 {
                    curr_part := part_stack.Pop().(catalog.Part)
                    _, isVisited := visited[curr_part.Id]
                    if !isVisited {
                        visited[curr_part.Id] = true
                        for _, child_part := range curr_part.Parts {
                            if child_part.Id != "" {
                                _, isVisitedChild := visited[child_part.Id]
                                if !isVisitedChild {
                                    query := "INSERT INTO `parts_parts`(`parent_partid`, `child_partid`) Values(\""
                                    query += curr_part.Id
                                    query += `", "`
                                    query += child_part.Id
                                    query += `")`
                                    _,err = db.Exec(query)
                                    if err != nil {
                                        fmt.Println(err.Error())
                                        fmt.Println("Caused by: " + query)
                                        os.Exit(0)
                                    }
                                    part_stack.Push(child_part)
                                } else {
                                    fmt.Println("CreatePartsToPartsTable Warning: Trying to add already visited child_part: " + child_part.Id)
                                }
                            }
                        }
                    } else {
                        fmt.Println("CreatePartsToPartsTable Warning: Trying to add already visited curr_part: " + curr_part.Id)
                    }
                }
            }
        }
    }
}

func CreatePartsToParagraphsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `parts_paragraph`(partid varchar(20), paragraph varchar(300))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    visited := make(map[string]bool)

    for _, ctrl := range GetAllControls(c) {
        part_stack := stackgo.NewStack()
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                part_stack.Push(part)
            }
        }
        for part_stack.Size() > 0 {
            curr_part := part_stack.Pop().(catalog.Part)
            _, isVisited := visited[curr_part.Id]
            if !isVisited {
                visited[curr_part.Id] = true
                fmt.Println(curr_part)
                //query := "INSERT INTO `parts-`(controlid, partid) Values(\""
                //query += ctrl.Id
                //query += `", "`
                //query += curr_part.Id
                //query += `")`
                //_,err = db.Exec(query)
                //if err != nil {
                //    fmt.Println(err.Error())
                //    fmt.Println("Caused by: " + query)
                //    os.Exit(0)
                //}
                for _, child_part := range curr_part.Parts {
                    if child_part.Id != "" {
                        _, isVisitedChild := visited[child_part.Id]
                        if !isVisitedChild {
                            part_stack.Push(child_part)
                        } else {
                            fmt.Println("CreateControlsToPartsTable Warning: Trying to add already visited child_part: " + child_part.Id)
                        }
                    }
                }
            } else {
                fmt.Println("CreateControlsToPartsTable Warning: Trying to add already visited curr_part: " + curr_part.Id)
            }
        }
    }
}

func main() {

    toLoad := "NIST_SP-800-53_rev4_catalog.xml"

    db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    data, e := ioutil.ReadFile(toLoad)
    if e != nil {
     fmt.Printf("error 1: %v\n", e)
     return
    }

    c := catalog.Catalog{}
    marshalError := xml.Unmarshal([]byte(data), &c)
    if marshalError != nil {
     fmt.Printf("error 2: %v\n", marshalError)
     return
    }

    CreatePartsTable(db, c)
    CreateParamsTable(db, c)
    CreateControlsToPartsTable(db, c)
    CreateControlsToParamsTable(db, c)
    CreatePartsToPartsTable(db, c)
    //CreatePartsToParagraphsTable(db, c)

}
