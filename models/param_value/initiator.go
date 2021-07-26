package param_value

import (
	//"encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type NullableParamValue struct {
	RecordId sql.NullInt32 `json:"recordid,omitempty"`
	ProjectId  sql.NullInt32 `json:"projectid,omitempty"`
	ComponentId  sql.NullString `json:"componentid,omitempty"`
	ParamId sql.NullString `json:"paramid,omitempty"`
	Value   sql.NullString `json:"value,omitempty"`
}

type ParamValue struct {
	RecordId  int `xml:"recordid,attr,omitempty" json:"recordid,omitempty"`
	ProjectId  int `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
	ComponentId  string `xml:"componentid,attr,omitempty" json:"componentid,omitempty"`
	ParamId string `xml:"paramid,omitempty" json:"paramid,omitempty"`
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

func UpdateParamValue(recordId int, value string) int {
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
	query := `UPDATE params_values SET value="` + value + `" WHERE record_id="` + strconv.Itoa(recordId) + `";`
	queryResult, err := db.Exec(query)
	fmt.Println(queryResult)

	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
	return recordId
}

func CreateParamValue(project_id int, component_id string, param_id string, value string) int {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `INSERT INTO params_values(project_id, component_id, param_id, value) Values("`
	query += strconv.Itoa(project_id)
	query += `", "`
	query += component_id
	query += `", "`
	query += param_id
	query += `", "`
	query += value
	query += `"); `
	queryResult, err := db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
	lastInsertId, err := queryResult.LastInsertId()
	return int(lastInsertId)
}

func DeleteParamValue(record_id int) {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `DELETE FROM params_values WHERE record_id="` + strconv.Itoa(record_id) + `";`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
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

func GetParamValue(record_id int) ParamValue {
	var result ParamValue
	var nullableResult NullableParamValue
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `SELECT record_id, project_id, component_id, param_id, params_values.value FROM params_values WHERE record_id = "`
	query += strconv.Itoa(record_id)
	query += `";`
	err = db.QueryRow(query).
		Scan(&nullableResult.RecordId, &nullableResult.ProjectId, &nullableResult.ComponentId, &nullableResult.ParamId, &nullableResult.Value)
	if err != nil {
		panic(err.Error())
	}
	// Validate the query response
	if nullableResult.RecordId.Valid {
		result.RecordId = int(nullableResult.RecordId.Int32)
	} else {
		result.RecordId = -1
	}
	if nullableResult.ProjectId.Valid {
		result.ProjectId = int(nullableResult.ProjectId.Int32)
	} else {
		result.ProjectId = -1
	}
	if nullableResult.ComponentId.Valid {
		result.ComponentId = nullableResult.ComponentId.String
	} else {
		result.ComponentId = ""
	}
	if nullableResult.ParamId.Valid {
		result.ParamId = nullableResult.ParamId.String
	} else {
		result.ParamId = ""
	}
	if nullableResult.Value.Valid {
		result.Value = nullableResult.Value.String
	} else {
		result.Value = ""
	}
	return result
}

func GetParam(project_id int) []ParamValue {
	results := make([]ParamValue, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	qs := `SELECT record_id, project_id, component_id, param_id, value ` +
		  `FROM params_values ` +
		  `WHERE project_id="` + strconv.Itoa(project_id) + `";`
	rows, err := db.Query(qs)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var result ParamValue
		var nullableResult NullableParamValue
		err := rows.Scan(&nullableResult.RecordId, &nullableResult.ProjectId, &nullableResult.ComponentId, &nullableResult.ParamId, &nullableResult.Value)
		if err != nil {
			panic(err.Error())
		}
		if nullableResult.RecordId.Valid {
			result.RecordId = int(nullableResult.RecordId.Int32)
		} else {
			result.RecordId = -1
		}
		if nullableResult.ProjectId.Valid {
			result.ProjectId = int(nullableResult.ProjectId.Int32)
		} else {
			result.ProjectId = -1
		}
		if nullableResult.ComponentId.Valid {
			result.ComponentId = nullableResult.ComponentId.String
		} else {
			result.ComponentId = ""
		}
		if nullableResult.ParamId.Valid {
			result.ParamId = nullableResult.ParamId.String
		} else {
			result.ParamId = ""
		}
		if nullableResult.Value.Valid {
			result.Value = nullableResult.Value.String
		} else {
			result.Value = ""
		}
		results = append(results, result)
	}
	return results
}
