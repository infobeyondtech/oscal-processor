package param_value

import (
	//"encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type NullableParamValue struct {
	fileid  sql.NullString `json:"fileid,omitempty"`
	paramid sql.NullString `json:"paramid,omitempty"`
	value   sql.NullString `json:"value,omitempty"`
}

type ParamValue struct {
	Fileid  string `xml:"fileid,attr,omitempty" json:"fileid,omitempty"`
	Paramid string `xml:"paramid,omitempty" json:"paramid,omitempty"`
	Value   string `xml:"value,omitempty" json:"value,omitempty"`
}

type Key struct {
	UUID    string
	Paramid string
}

type NullableKey struct {
	UUID    sql.NullString
	Paramid sql.NullString
}

type NullableParamInfo struct {
	paramid sql.NullString `json:"paramid,omitempty"`
	label sql.NullString `json:"label,omitempty"`
	sort sql.NullString `json:"sort,omitempty"`
	description sql.NullString `json:"description,omitempty"`
}

type SelectionChoice struct {
	Text string `json:"Text,omitempty"`
	Insert string `json:"Insert,omitempty"`
	InsertLabel string `json:"InsertLabel,omitempty"`
}

type ParamDescription struct {
	HowMany string `json:"HowMany,omitempty"`
	Choices []SelectionChoice `json:"choices,omitempty"`
}

type ParamInfo struct {
	Paramid string
	Label string
	Sort string
	Description ParamDescription
}

func SetParamValue(fileid string, paramid string, value string) ParamValue {
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
	query := `INSERT INTO params_values(fileid, paramid, value) Values("`
	query += fileid
	query += `", "`
	query += paramid
	query += `", "`
	query += value
	query += `")`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
	result := ParamValue{fileid, paramid, value}
	return result
}

func GetParamInfo(paramid string) ParamInfo {
	var result ParamInfo
	var nullableResult NullableParamInfo
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `SELECT paramid, label, sort, description FROM param_info WHERE paramid = "`
	query += paramid
	query += `";`
	err = db.QueryRow(query).
		Scan(&nullableResult.paramid, &nullableResult.label, &nullableResult.sort, &nullableResult.description)
	if err != nil {
		panic(err.Error())
	}
	// Validate the query response
	if nullableResult.paramid.Valid {
		result.Paramid = nullableResult.paramid.String
	} else {
		result.Paramid = ""
	}
	if nullableResult.label.Valid {
		result.Label = nullableResult.label.String
	} else {
		result.Label = ""
	}
	if nullableResult.sort.Valid {
		result.Sort = nullableResult.sort.String
	} else {
		result.Sort = ""
	}
	if nullableResult.description.Valid {
		var desc ParamDescription
		json.Unmarshal([]byte(nullableResult.description.String), &desc)
		result.Description = desc
		//result.Description = json.Unmarshal([]byte(nullableResult.description.String), )//nil//nullableResult.description.String
	} else {
		result.Description = ParamDescription{}
	}
	return result
}

func GetParamValue(fileid string, paramid string) ParamValue {
	var result ParamValue
	var nullableResult NullableParamValue
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `SELECT fileid, paramid, params_values.value FROM params_values WHERE fileid = "`
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
		result.Fileid = ""
	}
	if nullableResult.paramid.Valid {
		result.Paramid = nullableResult.paramid.String
	} else {
		result.Paramid = ""
	}
	if nullableResult.value.Valid {
		result.Value = nullableResult.value.String
	} else {
		result.Value = ""
	}
	return result
}

func GetParam(fileid string) []ParamValue {
	results := make([]ParamValue, 0)
	db, err := sql.Open("mysql", context.DBSource)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(`SELECT fileid, paramid, value FROM params_values WHERE fileid = "` + fileid + `";`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result ParamValue
		var nullableResult NullableParamValue

		err := rows.Scan(&nullableResult.fileid, &nullableResult.paramid, &nullableResult.value)
		if err != nil {
			panic(err.Error())
		}

		if nullableResult.fileid.Valid {
			result.Fileid = nullableResult.fileid.String
		} else {
			result.Fileid = ""
		}
		if nullableResult.paramid.Valid {
			result.Paramid = nullableResult.paramid.String
		} else {
			result.Paramid = ""
		}

		if nullableResult.value.Valid {
			result.Value = nullableResult.value.String
		} else {
			result.Value = ""
		}
		results = append(results, result)
	}

	return results
}
