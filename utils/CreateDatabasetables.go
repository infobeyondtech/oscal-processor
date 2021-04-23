package main

import (
    "os"
	"database/sql"
    //"encoding/xml"
    //"encoding/json"
    "fmt"
    //"io/ioutil"
    "github.com/docker/oscalkit/types/oscal/catalog"
    //"github.com/docker/oscalkit/types/oscal"
	_ "github.com/go-sql-driver/mysql"
    "github.com/alediaferia/stackgo"
)

type Param struct {
    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // A short name for the parameter.
    Label string `xml:"label,omitempty" json:"label,omitempty"`
}

type Part struct {
    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // Identifying the purpose and intended use of the property, part or other object.
    Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
    Prose string `xml:"prose,omitempty" json:"prose,omitempty"`
    // A partition or component of a control or part
    Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
}

type MyControl struct {
    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
    Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
    // A partition or component of a control or part
    Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
}

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

    // For all Controls
    for _, ctrl := range GetAllControls(c) {
        // For all parts that are directly children of ctrl
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                query := "INSERT INTO `controls_parts`(controlid, partid) Values(\""
                query += ctrl.Id
                query += `", "`
                query += part.Id
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

func CreateParamsToValuesTable(db *sql.DB) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `params_values`(fileid varchar(20), paramid varchar(20), value varchar(100))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }
}

func main() {

    //toLoad := "NIST_SP-800-53_rev4_catalog.xml"
    //toLoad := "ac-1-control.json";
    //toLoad := "NIST_SP-800-53_rev4_catalog.json"

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

    CreateParamsToValuesTable(db)

    //var result string;
    //query, err := db.Query(`call GetControlTree('ac-2', @result)`)
    //if err != nil {
    //    fmt.Println(err);
    //}
    //query.Next();
    //query.Scan(&result);
    //fmt.Println(result);
    //fmt.Println();
    //query, _ := db.Query(`select @result`)//.Next().scan(&result);
    //query.Next()
    //query.Scan(&result);
    //fmt.Println(result);

    //data, e := ioutil.ReadFile(toLoad)
    //if e != nil {
    // fmt.Printf("error 1: %v\n", e)
    // return
    //}

    //data:= `{ "id": "ac-1", "parameters": [ { "id": "ac-1_prm_1", "label": "organization-defined personnel or roles" }, { "id": "ac-1_prm_2", "label": "organization-defined frequency" }, { "id": "ac-1_prm_3", "label": "organization-defined frequency" } ], "parts": [ { "id": "ac-1_smt", "name": "statement", "prose": "The organization", "parts": [ { "id": "ac-1_smt.a", "name": "item", "prose": "Develops, documents, and disseminates to {{ ac-1_prm_1 }}", "parts": [ { "id": "ac-1_smt.a.1", "name": "item", "prose": "An access control policy that addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; an", "parts": [] }, { "id": "ac-1_smt.a.2", "name": "item", "prose": "Procedures to facilitate the implementation of the access control policy and associated access controls; an", "parts": [] } ] }, { "id": "ac-1_smt.b", "name": "item", "prose": "Reviews and updates the current", "parts": [ { "id": "ac-1_smt.b.1", "name": "item", "prose": "Access control policy {{ ac-1_prm_2 }}; an", "parts": [] }, { "id": "ac-1_smt.b.2", "name": "item", "prose": "Access control procedures {{ ac-1_prm_3 }}", "parts": [] } ] } ] }, { "id": "ac-1_gdn", "name": "guidance", "prose": "This control addresses the establishment of policy and procedures for the effective implementation of selected security controls and control enhancements in the AC family. Policy and procedures reflect applicable federal laws, Executive Orders, directives, regulations, policies, standards, and guidance. Security program policies and procedures at the organization level may make the need for system-specific policies and procedures unnecessary. The policy can be included as part of the general information security policy for organizations or conversely, can be represented by multiple policies reflecting the complex nature of certain organizations. The procedures can be established for the security program in general and for particular information systems, if needed. The organizational risk management strategy is a key factor in establishing policy and procedures", "parts": [] }, { "id": "ac-1_obj", "name": "objective", "prose": "Determine if the organization", "parts": [ { "id": "ac-1.a_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.a.1_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.a.1_obj.1", "name": "objective", "prose": "develops and documents an access control policy that addresses", "parts": [ { "id": "ac-1.a.1_obj.1.a", "name": "objective", "prose": "purpose", "parts": [] }, { "id": "ac-1.a.1_obj.1.b", "name": "objective", "prose": "scope", "parts": [] }, { "id": "ac-1.a.1_obj.1.c", "name": "objective", "prose": "roles", "parts": [] }, { "id": "ac-1.a.1_obj.1.d", "name": "objective", "prose": "responsibilities", "parts": [] }, { "id": "ac-1.a.1_obj.1.e", "name": "objective", "prose": "management commitment", "parts": [] }, { "id": "ac-1.a.1_obj.1.f", "name": "objective", "prose": "coordination among organizational entities", "parts": [] }, { "id": "ac-1.a.1_obj.1.g", "name": "objective", "prose": "compliance", "parts": [] } ] }, { "id": "ac-1.a.1_obj.2", "name": "objective", "prose": "defines personnel or roles to whom the access control policy are to be disseminated", "parts": [] }, { "id": "ac-1.a.1_obj.3", "name": "objective", "prose": "disseminates the access control policy to organization-defined personnel or roles", "parts": [] } ] }, { "id": "ac-1.a.2_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.a.2_obj.1", "name": "objective", "prose": "develops and documents procedures to facilitate the implementation of the access control policy and associated access control controls", "parts": [] }, { "id": "ac-1.a.2_obj.2", "name": "objective", "prose": "defines personnel or roles to whom the procedures are to be disseminated", "parts": [] }, { "id": "ac-1.a.2_obj.3", "name": "objective", "prose": "disseminates the procedures to organization-defined personnel or roles", "parts": [] } ] } ] }, { "id": "ac-1.b_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.b.1_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.b.1_obj.1", "name": "objective", "prose": "defines the frequency to review and update the current access control policy", "parts": [] }, { "id": "ac-1.b.1_obj.2", "name": "objective", "prose": "reviews and updates the current access control policy with the organization-defined frequency", "parts": [] } ] }, { "id": "ac-1.b.2_obj", "name": "objective", "prose": "", "parts": [ { "id": "ac-1.b.2_obj.1", "name": "objective", "prose": "defines the frequency to review and update the current access control procedures; an", "parts": [] }, { "id": "ac-1.b.2_obj.2", "name": "objective", "prose": "reviews and updates the current access control procedures with the organization-defined frequency", "parts": [] } ] } ] } ] } ] }`
    ////data := `{"id": "ac-1", "class": "test"}`
    ////data := `{ "id": "ac-1", "class": "SP800-53", "title": "Access Control Policy and Procedures", "parameters": [ { "id": "ac-1_prm_1", "label": "organization-defined personnel or roles" }, { "id": "ac-1_prm_2", "label": "organization-defined frequency" }, { "id": "ac-1_prm_3", "label": "organization-defined frequency" } ], "properties": [ { "name": "label", "value": "AC-1" }, { "name": "sort-id", "value": "ac-01" } ], "links": [ { "href": "#ref050", "rel": "reference", "text": "NIST Special Publication 800-12" }, { "href": "#ref044", "rel": "reference", "text": "NIST Special Publication 800-100" } ], "parts": [ { "id": "ac-1_smt", "name": "statement", "prose": "The organization:", "parts": [ { "id": "ac-1_smt.a", "name": "item", "properties": [ { "name": "label", "value": "a." } ], "prose": "Develops, documents, and disseminates to {{ ac-1_prm_1 }}:", "parts": [ { "id": "ac-1_smt.a.1", "name": "item", "properties": [ { "name": "label", "value": "1." } ], "prose": "An access control policy that addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and" }, { "id": "ac-1_smt.a.2", "name": "item", "properties": [ { "name": "label", "value": "2." } ], "prose": "Procedures to facilitate the implementation of the access control policy and associated access controls; and" } ] }, { "id": "ac-1_smt.b", "name": "item", "properties": [ { "name": "label", "value": "b." } ], "prose": "Reviews and updates the current:", "parts": [ { "id": "ac-1_smt.b.1", "name": "item", "properties": [ { "name": "label", "value": "1." } ], "prose": "Access control policy {{ ac-1_prm_2 }}; and" }, { "id": "ac-1_smt.b.2", "name": "item", "properties": [ { "name": "label", "value": "2." } ], "prose": "Access control procedures {{ ac-1_prm_3 }}." } ] } ] }, { "id": "ac-1_gdn", "name": "guidance", "prose": "This control addresses the establishment of policy and procedures for the effective implementation of selected security controls and control enhancements in the AC family. Policy and procedures reflect applicable federal laws, Executive Orders, directives, regulations, policies, standards, and guidance. Security program policies and procedures at the organization level may make the need for system-specific policies and procedures unnecessary. The policy can be included as part of the general information security policy for organizations or conversely, can be represented by multiple policies reflecting the complex nature of certain organizations. The procedures can be established for the security program in general and for particular information systems, if needed. The organizational risk management strategy is a key factor in establishing policy and procedures.", "links": [ { "href": "#pm-9", "rel": "related", "text": "PM-9" } ] }, { "id": "ac-1_obj", "name": "objective", "prose": "Determine if the organization:", "parts": [ { "id": "ac-1.a_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)" } ], "parts": [ { "id": "ac-1.a.1_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)" } ], "parts": [ { "id": "ac-1.a.1_obj.1", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1]" } ], "prose": "develops and documents an access control policy that addresses:", "parts": [ { "id": "ac-1.a.1_obj.1.a", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][a]" } ], "prose": "purpose;" }, { "id": "ac-1.a.1_obj.1.b", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][b]" } ], "prose": "scope;" }, { "id": "ac-1.a.1_obj.1.c", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][c]" } ], "prose": "roles;" }, { "id": "ac-1.a.1_obj.1.d", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][d]" } ], "prose": "responsibilities;" }, { "id": "ac-1.a.1_obj.1.e", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][e]" } ], "prose": "management commitment;" }, { "id": "ac-1.a.1_obj.1.f", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][f]" } ], "prose": "coordination among organizational entities;" }, { "id": "ac-1.a.1_obj.1.g", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[1][g]" } ], "prose": "compliance;" } ] }, { "id": "ac-1.a.1_obj.2", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[2]" } ], "prose": "defines personnel or roles to whom the access control policy are to be disseminated;" }, { "id": "ac-1.a.1_obj.3", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(1)[3]" } ], "prose": "disseminates the access control policy to organization-defined personnel or roles;" } ] }, { "id": "ac-1.a.2_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(2)" } ], "parts": [ { "id": "ac-1.a.2_obj.1", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(2)[1]" } ], "prose": "develops and documents procedures to facilitate the implementation of the access control policy and associated access control controls;" }, { "id": "ac-1.a.2_obj.2", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(2)[2]" } ], "prose": "defines personnel or roles to whom the procedures are to be disseminated;" }, { "id": "ac-1.a.2_obj.3", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(a)(2)[3]" } ], "prose": "disseminates the procedures to organization-defined personnel or roles;" } ] } ] }, { "id": "ac-1.b_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)" } ], "parts": [ { "id": "ac-1.b.1_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(1)" } ], "parts": [ { "id": "ac-1.b.1_obj.1", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(1)[1]" } ], "prose": "defines the frequency to review and update the current access control policy;" }, { "id": "ac-1.b.1_obj.2", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(1)[2]" } ], "prose": "reviews and updates the current access control policy with the organization-defined frequency;" } ] }, { "id": "ac-1.b.2_obj", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(2)" } ], "parts": [ { "id": "ac-1.b.2_obj.1", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(2)[1]" } ], "prose": "defines the frequency to review and update the current access control procedures; and" }, { "id": "ac-1.b.2_obj.2", "name": "objective", "properties": [ { "name": "label", "value": "AC-1(b)(2)[2]" } ], "prose": "reviews and updates the current access control procedures with the organization-defined frequency." } ] } ] } ] }, { "name": "assessment", "properties": [ { "name": "method", "value": "EXAMINE" } ], "parts": [ { "name": "objects", "prose": "Access control policy and procedures\\n\\nother relevant documents or records" } ] }, { "name": "assessment", "properties": [ { "name": "method", "value": "INTERVIEW" } ], "parts": [ { "name": "objects", "prose": "Organizational personnel with access control responsibilities\\n\\norganizational personnel with information security responsibilities" } ] } ] }`
    //c := MyControl{}
    //c := catalog.Catalog{}
    //marshalError := json.Unmarshal([]byte(result), &c)
    //if marshalError != nil {
    // fmt.Printf("error 2: %v\n", marshalError)
    // return
    //}

    //fmt.Print(c);

    //CreatePartsTable(db, c)
    //CreateParamsTable(db, c)
    //CreateControlsToPartsTable(db, c)
    //CreateControlsToParamsTable(db, c)
    //CreatePartsToPartsTable(db, c)
    //CreatePartsToParagraphsTable(db, c)

}
