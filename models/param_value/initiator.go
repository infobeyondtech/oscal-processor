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
    fileid string `xml:"id,attr,omitempty" json:"id,omitempty"`
    paramid string `xml:"label,omitempty" json:"label,omitempty"`
    value string `xml:"label,omitempty" json:"label,omitempty"`
}

func SetParamValue(fileid string, paramid string, value string) (ParamValue) {
    // Open the DB
    db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube");
    if err != nil {
        panic(err.Error())
    }

    // TODO: Do we need to error check to make sure the
    //       fileid, and paramid are valid?

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
    if nullableResult.fileid.Valid {
        result.fileid = nullableResult.fileid.String
    } else {
        result.fileid = "";
    }
    if nullableResult.paramid.Valid {
        result.paramid = nullableResult.paramid.String
    } else {
        result.paramid = "";
    }
    if nullableResult.value.Valid {
        result.value = nullableResult.value.String
    } else {
        result.value = "";
    }

    return result;
}

//func main() {
//    //SetParamValue("fileid1", "paramid1", "value1")
//    fmt.Println(GetParamValue("fileid1", "paramid1"))
//}
