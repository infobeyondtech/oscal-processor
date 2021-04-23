package param_value

import (
    //"encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

type NullableParamValue struct {
    fileid sql.NullString `json:"fileid,omitempty"`
    paramid sql.NullString `json:"paramid,omitempty"`
    value sql.NullString `json:"value,omitempty"`
}

type ParamValue struct {
    Fileid string `xml:"fileid,attr,omitempty" json:"fileid,omitempty"`
    Paramid string `xml:"paramid,omitempty" json:"paramid,omitempty"`
    Value string `xml:"value,omitempty" json:"value,omitempty"`
}

func SetParamValue(fileid string, paramid string, value string) (ParamValue) {
    // Open the DB
    db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube");
    if err != nil {
        panic(err.Error())
    }

    // TODO: Do we need to error check to make sure the
    //       fileid, and paramid are valid?

    // TODO: Check to see if value already exisits in DB
    //       Update if so

    query := `INSERT INTO params_values(fileid, paramid, value) Values("`
    query += fileid
    query += `", "`
    query += paramid
    query += `", "`
    query += value
    query += `")`
    _,err = db.Exec(query)
    if err != nil {
        fmt.Println("Caused by: " + query)
        panic(err.Error())
    }
    result := ParamValue{fileid, paramid, value}
    return result;
}

func GetParamValue(fileid string, paramid string) (ParamValue) {
    var result ParamValue
    var nullableResult NullableParamValue
    // Open the DB
    db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube");
    if err != nil {
        panic(err.Error())
    }
    query := `SELECT * FROM params_values WHERE fileid = "`
    query += fileid
    query += `" and paramid = "`
    query += paramid
    query += `";`
    err = db.QueryRow(query).
        Scan(&nullableResult.fileid, &nullableResult.paramid, &nullableResult.value)
    if err != nil {
        panic(err.Error())
    }
    // Validate the query response
    if nullableResult.fileid.Valid {
        result.Fileid = nullableResult.fileid.String
    } else {
        result.Fileid = "";
    }
    if nullableResult.paramid.Valid {
        result.Paramid = nullableResult.paramid.String
    } else {
        result.Paramid = "";
    }
    if nullableResult.value.Valid {
        result.Value = nullableResult.value.String
    } else {
        result.Value = "";
    }

    return result;
}
