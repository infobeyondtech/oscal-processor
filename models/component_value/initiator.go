package component_value

//import (
//	"database/sql"
//	"github.com/infobeyondtech/oscal-processor/context"
//	"strconv"
//)
import (
	//"encoding/json"
	"database/sql"
	"strings"

	//"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type NullableComponentValue struct {
	RecordId    sql.NullInt32  `json:"recordid,omitempty"`
	ProjectId   sql.NullInt32  `json:"projectid,omitempty"`
	StatementId sql.NullString `json:"statementid,omitempty"`
	ComponentId sql.NullString `json:"componentid,omitempty"`
}

type ComponentValue struct {
	RecordId  int `xml:"recordid,attr,omitempty" json:"recordid,omitempty"`
	ProjectId  int `xml:"projectid,attr,omitempty" json:"projectid,omitempty"`
	StatementId string `xml:"statmentid,attr,omitempty" json:"statementid,omitempty"`
	ComponentId  string `xml:"componentid,attr,omitempty" json:"componentid,omitempty"`
}

func UpdateComponentValue(recordId int, componentId string) int {
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
	query := `UPDATE components_values SET componentId="` + componentId + `" WHERE recordId="` + strconv.Itoa(recordId) + `";`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
	return recordId
}

func CreateComponentValue(projectId int, statementId string, componentId string) int {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `INSERT INTO components_values(projectId, statementId, componentId) Values("`
	query += strconv.Itoa(projectId)
	query += `", "`
	query += statementId
	query += `", "`
	query += componentId
	query += `"); `
	fmt.Println(query)
	queryResult, err := db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
	lastInsertId, err := queryResult.LastInsertId()
	return int(lastInsertId)
}

func DeleteComponentValue(recordId int) {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `DELETE FROM components_values WHERE recordId="` + strconv.Itoa(recordId) + `";`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Caused by: " + query)
		panic(err.Error())
	}
}

func GetComponentValue(recordId int) ComponentValue {
	var result ComponentValue
	var nullableResult NullableComponentValue
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `SELECT recordId, projectId, statementId, componentId FROM components_values WHERE recordId = "`
	query += strconv.Itoa(recordId)
	query += `";`
	err = db.QueryRow(query).
		Scan(&nullableResult.RecordId, &nullableResult.ProjectId, &nullableResult.StatementId, &nullableResult.ComponentId)
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
	if nullableResult.StatementId.Valid {
		result.StatementId = nullableResult.StatementId.String
	} else {
		result.StatementId = ""
	}
	if nullableResult.ComponentId.Valid {
		result.ComponentId = nullableResult.ComponentId.String
	} else {
		result.ComponentId = ""
	}
	return result
}

func GetControlToStatementMap(projectId int) map[string][]string {
	controlToStatementMap := make(map[string][]string)
	cvs := GetComponent(projectId)
	for _, cv := range cvs {
		ctrlId := strings.Split(cv.StatementId, "_")[0]
		if statementIDs, ok := controlToStatementMap[ctrlId]; !ok {
			controlToStatementMap[ctrlId] = make([]string, 0)
			controlToStatementMap[ctrlId] = append(controlToStatementMap[ctrlId], cv.StatementId)
		} else {
			exists := false
			for _, s := range statementIDs {
				if s == cv.StatementId {
					exists = true
				}
			}
			if !exists {
				controlToStatementMap[ctrlId] = append(controlToStatementMap[ctrlId], cv.StatementId)
			}
		}
	}
	return controlToStatementMap
}

func GetStatementToComponentMap(projectId int) map[string][]string {
	statementToComponentMap := make(map[string][]string)
	cvs := GetComponent(projectId)
	for _, cv := range cvs {
		if _, exists := statementToComponentMap[cv.StatementId]; !exists {
			statementToComponentMap[cv.StatementId] = make([]string, 0)
		}
		statementToComponentMap[cv.StatementId] = append(statementToComponentMap[cv.StatementId], cv.ComponentId)
	}
	return statementToComponentMap
}

func GetComponent(projectId int) []ComponentValue {
	results := make([]ComponentValue, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	qs := `SELECT recordId, projectId, statementId, componentId ` +
		  `FROM components_values ` +
		  `WHERE projectId="` + strconv.Itoa(projectId) + `";`
	rows, err := db.Query(qs)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var result ComponentValue
		var nullableResult NullableComponentValue
		err := rows.Scan(&nullableResult.RecordId, &nullableResult.ProjectId, &nullableResult.StatementId, &nullableResult.ComponentId)
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
		if nullableResult.StatementId.Valid {
			result.StatementId = nullableResult.StatementId.String
		} else {
			result.ComponentId = ""
		}
		if nullableResult.ComponentId.Valid {
			result.ComponentId = nullableResult.ComponentId.String
		} else {
			result.ComponentId = ""
		}
		results = append(results, result)
	}
	return results
}
