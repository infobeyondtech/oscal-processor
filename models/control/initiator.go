package control

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
	"log"
	"strings"
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
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
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

type ControlSummary struct {
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty"`
}

func GetControl(ctrlId string, isEnh bool) Control {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
    defer db.Close()
	// Call the Stored Procedure, GetControlTree or GetEnhancementTree
	var queryResult string
	var queryString string
	if isEnh {
		queryString = `call GetEnhancementTree('` + ctrlId + `', @result)`
	} else {
		queryString = `call GetControlTree('` + ctrlId + `', @result)`
	}
	query, err := db.Query(queryString)
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
	return result
}

func GetControlEnhancementIds(ctrlId string) []string {
	// Open the DB
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
    defer db.Close()
	queryString := `SELECT controlid, enhid FROM controls_enhancements WHERE controls_enhancements.controlid = '`
	queryString += ctrlId + `'`
	queryResult, err := db.Query(queryString)
	if err != nil {
		panic(err.Error())
	}
	result := make([]string, 0)
	for queryResult.Next() {
		var currControlId string
		var currEnhId string
		err := queryResult.Scan(&currControlId, &currEnhId)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, currEnhId)
	}
	return result
}

func SearchPartsWithKeyword(keyword string) []string {
    // Open the DB
	db, err := sql.Open("mysql", context.DBSource)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    // query all parts' proses for the keyword
    queryString := `SELECT partid, name, prose FROM parts WHERE parts.prose LIKE '`
    queryString += `%` + keyword + `%`
    queryString += `'`
    queryResult, err := db.Query(queryString)
    if err != nil {
        panic(err.Error())
    }
    var id string
    var name string
    var prose string
    partIds := make([]string, 0)
    // Get all the parts that contain the keyword
    for queryResult.Next() {
        err := queryResult.Scan(&id, &name, &prose)
        if err != nil {
            log.Fatal(err)
        } else {
        	partIds = append(partIds, id)
		}
	}
	// Make a map (set) for the controls that contain parts with the provided keyword.
	ctrlMap := make(map[string]bool)
	for _, partId := range partIds {
		// Parse the partId to get its control ancestor
		ctrlId := strings.Split(partId, `.`)[0]
		ctrlId = strings.Split(ctrlId, `_`)[0]
		ctrlMap[ctrlId] = true
	}
	result := make([]string, 0)
	for key,_ := range ctrlMap {
		result = append(result, key)
	}
	return result
}
