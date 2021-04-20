package control

import (
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
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

type Control struct {
    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
    Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
    // A partition or component of a control or part
    Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
}

func GetControl(ctrlId string) (Control) {
    // Open the DB
    db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube");
    if err != nil {
        panic(err.Error())
    }
    // Call the Stored Procedure, GetControlTree
    var queryResult string;
    query, err := db.Query(`call GetControlTree('` + ctrlId + `', @result)`)
    if err != nil {
        panic(err.Error())
    }
    // Get the result, Unmarshal, and return the Control
    query.Next()
    query.Scan(&queryResult)
    result := Control{}
    marshalError := json.Unmarshal([]byte(queryResult), &result)
    if marshalError != nil {
        panic(marshalError)
    }
    return result;
}
