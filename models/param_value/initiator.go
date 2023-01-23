package param_value

import (
	//"encoding/json"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	data_models "github.com/infobeyondtech/oscal-processor/models/data_models/requests_model"
	"io/ioutil"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type Catalog struct {
	XMLName  xml.Name `xml:"catalog"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	ID       string   `xml:"id,attr"`
	Metadata struct {
		Text         string `xml:",chardata"`
		Title        string `xml:"title"`
		LastModified string `xml:"last-modified"`
		Version      string `xml:"version"`
		OscalVersion string `xml:"oscal-version"`
		Prop         struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"prop"`
		Link []struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Role []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Title string `xml:"title"`
		} `xml:"role"`
		Party struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Org  struct {
				Text    string `xml:",chardata"`
				OrgName string `xml:"org-name"`
				Address struct {
					Text       string   `xml:",chardata"`
					AddrLine   []string `xml:"addr-line"`
					City       string   `xml:"city"`
					State      string   `xml:"state"`
					PostalCode string   `xml:"postal-code"`
				} `xml:"address"`
				Email string `xml:"email"`
			} `xml:"org"`
		} `xml:"party"`
		ResponsibleParty []struct {
			Text    string `xml:",chardata"`
			RoleID  string `xml:"role-id,attr"`
			PartyID string `xml:"party-id"`
		} `xml:"responsible-party"`
	} `xml:"metadata"`
	Group []struct {
		Text    string `xml:",chardata"`
		Class   string `xml:"class,attr"`
		ID      string `xml:"id,attr"`
		Title   string `xml:"title"`
		Control []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			ID    string `xml:"id,attr"`
			Title string `xml:"title"`
			Param []struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				DependsOn string `xml:"depends-on,attr"`
				Label     string `xml:"label"`
				Select    struct {
					Text    string `xml:",chardata"`
					HowMany string `xml:"how-many,attr"`
					Choice  []struct {
						Text   string `xml:",chardata"`
						Insert struct {
							Text    string `xml:",chardata"`
							ParamID string `xml:"param-id,attr"`
						} `xml:"insert"`
					} `xml:"choice"`
				} `xml:"select"`
			} `xml:"param"`
			Prop []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"prop"`
			Link []struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
				Rel  string `xml:"rel,attr"`
			} `xml:"link"`
			Part []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Name string `xml:"name,attr"`
				P    struct {
					Text   string `xml:",chardata"`
					Insert []struct {
						Text    string `xml:",chardata"`
						ParamID string `xml:"param-id,attr"`
					} `xml:"insert"`
				} `xml:"p"`
				Part []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Name string `xml:"name,attr"`
					Prop struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"prop"`
					P []struct {
						Text   string `xml:",chardata"`
						Insert []struct {
							Text    string `xml:",chardata"`
							ParamID string `xml:"param-id,attr"`
						} `xml:"insert"`
					} `xml:"p"`
					Part []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Name string `xml:"name,attr"`
						Prop struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"prop"`
						P struct {
							Text   string `xml:",chardata"`
							Insert []struct {
								Text    string `xml:",chardata"`
								ParamID string `xml:"param-id,attr"`
							} `xml:"insert"`
						} `xml:"p"`
						Part []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Name string `xml:"name,attr"`
							Prop struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"prop"`
							P struct {
								Text   string `xml:",chardata"`
								Insert []struct {
									Text    string `xml:",chardata"`
									ParamID string `xml:"param-id,attr"`
								} `xml:"insert"`
							} `xml:"p"`
							Part []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Name string `xml:"name,attr"`
								Prop struct {
									Text string `xml:",chardata"`
									Name string `xml:"name,attr"`
								} `xml:"prop"`
								P struct {
									Text   string `xml:",chardata"`
									Insert []struct {
										Text    string `xml:",chardata"`
										ParamID string `xml:"param-id,attr"`
									} `xml:"insert"`
								} `xml:"p"`
								Part []struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Name string `xml:"name,attr"`
									Prop struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"prop"`
									P struct {
										Text   string `xml:",chardata"`
										Insert []struct {
											Text    string `xml:",chardata"`
											ParamID string `xml:"param-id,attr"`
										} `xml:"insert"`
									} `xml:"p"`
								} `xml:"part"`
							} `xml:"part"`
						} `xml:"part"`
					} `xml:"part"`
				} `xml:"part"`
				Link []struct {
					Text string `xml:",chardata"`
					Rel  string `xml:"rel,attr"`
					Href string `xml:"href,attr"`
				} `xml:"link"`
				Prop struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"prop"`
			} `xml:"part"`
			Control []struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
				ID    string `xml:"id,attr"`
				Title string `xml:"title"`
				Prop  []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"prop"`
				Part []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Name string `xml:"name,attr"`
					P    struct {
						Text   string `xml:",chardata"`
						Insert []struct {
							Text    string `xml:",chardata"`
							ParamID string `xml:"param-id,attr"`
						} `xml:"insert"`
					} `xml:"p"`
					Prop struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"prop"`
					Part []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						ID   string `xml:"id,attr"`
						P    []struct {
							Text   string `xml:",chardata"`
							Insert []struct {
								Text    string `xml:",chardata"`
								ParamID string `xml:"param-id,attr"`
							} `xml:"insert"`
						} `xml:"p"`
						Prop struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"prop"`
						Part []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Name string `xml:"name,attr"`
							Prop struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"prop"`
							P struct {
								Text   string `xml:",chardata"`
								Insert struct {
									Text    string `xml:",chardata"`
									ParamID string `xml:"param-id,attr"`
								} `xml:"insert"`
							} `xml:"p"`
							Link struct {
								Text string `xml:",chardata"`
								Rel  string `xml:"rel,attr"`
								Href string `xml:"href,attr"`
							} `xml:"link"`
							Part []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Name string `xml:"name,attr"`
								Prop struct {
									Text string `xml:",chardata"`
									Name string `xml:"name,attr"`
								} `xml:"prop"`
								P struct {
									Text   string `xml:",chardata"`
									Insert []struct {
										Text    string `xml:",chardata"`
										ParamID string `xml:"param-id,attr"`
									} `xml:"insert"`
								} `xml:"p"`
								Link struct {
									Text string `xml:",chardata"`
									Rel  string `xml:"rel,attr"`
									Href string `xml:"href,attr"`
								} `xml:"link"`
								Part []struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Name string `xml:"name,attr"`
									Prop struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"prop"`
									P struct {
										Text   string `xml:",chardata"`
										Insert []struct {
											Text    string `xml:",chardata"`
											ParamID string `xml:"param-id,attr"`
										} `xml:"insert"`
									} `xml:"p"`
								} `xml:"part"`
							} `xml:"part"`
						} `xml:"part"`
						Link struct {
							Text string `xml:",chardata"`
							Rel  string `xml:"rel,attr"`
							Href string `xml:"href,attr"`
						} `xml:"link"`
					} `xml:"part"`
					Link []struct {
						Text string `xml:",chardata"`
						Rel  string `xml:"rel,attr"`
						Href string `xml:"href,attr"`
					} `xml:"link"`
				} `xml:"part"`
				Param []struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"id,attr"`
					DependsOn string `xml:"depends-on,attr"`
					Select    struct {
						Text    string `xml:",chardata"`
						HowMany string `xml:"how-many,attr"`
						Choice  []struct {
							Text   string `xml:",chardata"`
							Insert []struct {
								Text    string `xml:",chardata"`
								ParamID string `xml:"param-id,attr"`
							} `xml:"insert"`
						} `xml:"choice"`
					} `xml:"select"`
					Label string `xml:"label"`
				} `xml:"param"`
				Link []struct {
					Text string `xml:",chardata"`
					Rel  string `xml:"rel,attr"`
					Href string `xml:"href,attr"`
				} `xml:"link"`
			} `xml:"control"`
		} `xml:"control"`
	} `xml:"group"`
	BackMatter struct {
		Text     string `xml:",chardata"`
		Resource []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Title    string `xml:"title"`
			Citation struct {
				Chardata string `xml:",chardata"`
				Text     string `xml:"text"`
			} `xml:"citation"`
			Rlink struct {
				Text      string `xml:",chardata"`
				Href      string `xml:"href,attr"`
				MediaType string `xml:"media-type,attr"`
			} `xml:"rlink"`
			DocID struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"doc-id"`
		} `xml:"resource"`
	} `xml:"back-matter"`
}

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

func getParamsFromString(input string, controlID string) []string {
	var result []string
	toSearch := input
	for strings.Index(toSearch, controlID+"_prm_") != -1 {
		idx := strings.Index(toSearch, controlID+"_prm_")
		c := string(toSearch[idx])
		pId := ""
		for c != "}" {
			pId += c
			idx += 1
			c = string(toSearch[idx])
		}
		result = append(result, pId)
		toSearch = toSearch[idx:]
	}
	return result
}

func GetStmtCompToParamMap(projectId int) map[string][]data_models.SetParameter {
	paramToPartMap := getParamToPartMap()
	stmtCompToParamMap := make(map[string][]data_models.SetParameter)
	pvs := GetParam(projectId)
	for _, pv := range pvs {
		statementID := paramToPartMap[pv.ParamId]
		if _, exists := stmtCompToParamMap[statementID + pv.ComponentId]; !exists {
			stmtCompToParamMap[statementID + pv.ComponentId] = make([]data_models.SetParameter, 0)
		}
		sp := data_models.SetParameter{pv.ParamId, pv.Value}
		stmtCompToParamMap[statementID+pv.ComponentId] = append(stmtCompToParamMap[statementID+pv.ComponentId], sp)
	}
	return stmtCompToParamMap
}

func getParamToPartMap() map[string]string {
	toLoad := "../models/param_value/NIST_SP-800-53_rev4_catalog.xml"
	data, _ := ioutil.ReadFile(toLoad)
	c := Catalog{}
	marshalError := xml.Unmarshal([]byte(data), &c)
	if marshalError != nil {
		fmt.Printf("error 2: %v\n", marshalError)
		return nil
	}
	paramToPartMap := make(map[string]string)
	for _, group := range c.Group {
		for _, ctrl := range group.Control {
			for _, part := range ctrl.Part {
				//paramToPartMap[part.ID] = getParamsFromString(fmt.Sprintf("%s", part.P), ctrl.ID)
				params := getParamsFromString(fmt.Sprintf("%s", part.P), ctrl.ID)
				for _, p := range params {
					paramToPartMap[p] = part.ID
				}
				for _, child1 := range part.Part {
					//paramToPartMap[child1.ID] = getParamsFromString(fmt.Sprintf("%s", child1.P), ctrl.ID)
					params = getParamsFromString(fmt.Sprintf("%s", child1.P), ctrl.ID)
					for _, p := range params {
						paramToPartMap[p] = child1.ID
					}
					for _, child2 := range child1.Part {
						//paramToPartMap[child2.ID] = getParamsFromString(fmt.Sprintf("%s", child2.P), ctrl.ID)
						params = getParamsFromString(fmt.Sprintf("%s", child2.P), ctrl.ID)
						for _, p := range params {
							paramToPartMap[p] = child2.ID
						}
						for _, child3 := range child2.Part {
							//paramToPartMap[child3.ID] = getParamsFromString(fmt.Sprintf("%s", child3.P), ctrl.ID)
							params = getParamsFromString(fmt.Sprintf("%s", child3.P), ctrl.ID)
							for _, p := range params {
								paramToPartMap[p] = child3.ID
							}
							for _, child4 := range child3.Part {
								//paramToPartMap[child4.ID] = getParamsFromString(fmt.Sprintf("%s", child4.P), ctrl.ID)
								params = getParamsFromString(fmt.Sprintf("%s", child4.P), ctrl.ID)
								for _, p := range params {
									paramToPartMap[p] = child4.ID
								}
								for _, child5 := range child4.Part {
									//paramToPartMap[child5.ID] = getParamsFromString(fmt.Sprintf("%s", child5.P), ctrl.ID)
									params = getParamsFromString(fmt.Sprintf("%s", child5.P), ctrl.ID)
									for _, p := range params {
										paramToPartMap[p] = child5.ID
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return paramToPartMap
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
