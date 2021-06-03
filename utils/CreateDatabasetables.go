package main

import (
	"database/sql"
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"fmt"
	"github.com/alediaferia/stackgo"
	"io/ioutil"
	//"github.com/docker/oscalkit/types/oscal/catalog"
	catalog "github.com/docker/oscalkit/types/oscal/catalog"
	_ "github.com/go-sql-driver/mysql"
	//"io/ioutil"
)

type Controls struct {
	XMLName        xml.Name `xml:"controls"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Controls       string   `xml:"controls,attr"`
	Xhtml          string   `xml:"xhtml,attr"`
	Xsi            string   `xml:"xsi,attr"`
	PubDate        string   `xml:"pub_date,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Control        []struct {
		Text           string   `xml:",chardata"`
		Family         string   `xml:"family"`
		Number         string   `xml:"number"`
		Title          string   `xml:"title"`
		Priority       string   `xml:"priority"`
		BaselineImpact []string `xml:"baseline-impact"`
		Statement      struct {
			Text        string `xml:",chardata"`
			Description string `xml:"description"`
			Statement   []struct {
				Text        string `xml:",chardata"`
				Number      string `xml:"number"`
				Description string `xml:"description"`
				Statement   []struct {
					Text        string `xml:",chardata"`
					Number      string `xml:"number"`
					Description string `xml:"description"`
				} `xml:"statement"`
			} `xml:"statement"`
		} `xml:"statement"`
		SupplementalGuidance struct {
			Text        string   `xml:",chardata"`
			Description string   `xml:"description"`
			Related     []string `xml:"related"`
		} `xml:"supplemental-guidance"`
		References struct {
			Text      string `xml:",chardata"`
			Reference []struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
					Href string `xml:"href,attr"`
				} `xml:"item"`
			} `xml:"reference"`
		} `xml:"references"`
		ControlEnhancements struct {
			Text               string `xml:",chardata"`
			ControlEnhancement []struct {
				Text           string   `xml:",chardata"`
				Number         string   `xml:"number"`
				Title          string   `xml:"title"`
				BaselineImpact []string `xml:"baseline-impact"`
				Statement      struct {
					Text        string `xml:",chardata"`
					Description string `xml:"description"`
					Statement   []struct {
						Text        string `xml:",chardata"`
						Number      string `xml:"number"`
						Description string `xml:"description"`
						Statement   []struct {
							Text        string `xml:",chardata"`
							Number      string `xml:"number"`
							Description string `xml:"description"`
						} `xml:"statement"`
					} `xml:"statement"`
				} `xml:"statement"`
				SupplementalGuidance struct {
					Text        string   `xml:",chardata"`
					Description string   `xml:"description"`
					Related     []string `xml:"related"`
				} `xml:"supplemental-guidance"`
				Withdrawn struct {
					Text             string   `xml:",chardata"`
					IncorporatedInto []string `xml:"incorporated-into"`
				} `xml:"withdrawn"`
			} `xml:"control-enhancement"`
		} `xml:"control-enhancements"`
		Withdrawn struct {
			Text             string   `xml:",chardata"`
			IncorporatedInto []string `xml:"incorporated-into"`
		} `xml:"withdrawn"`
	} `xml:"control"`
}

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
    //visited := make(map[string]bool)
    //control_stack := stackgo.NewStack()

    // Add 0-depth controls from catalog to stack
    for _, group := range c.Groups {
        for _, ctrl := range group.Controls {
            // Skipping controls without an Id
            if ctrl.Id != "" {
                //control_stack.Push(ctrl)
                result = append(result, ctrl)
            }
        }
    }

    //for control_stack.Size() > 0 {
    //    curr_control := control_stack.Pop().(catalog.Control)
    //    _, isVisited := visited[curr_control.Id]
    //    if !isVisited {
    //        visited[curr_control.Id] = true
    //        result = append(result, curr_control)
    //        for _, child_control := range curr_control.Controls {
    //            // Skipping controls without an Id
    //            if child_control.Id != "" {
    //                _, isVisitedChild := visited[child_control.Id]
    //                if !isVisitedChild {
    //                    control_stack.Push(child_control)
    //                } else {
    //                    fmt.Println("GetAllControls Warning: Trying to add already visited child_control: " + child_control.Id)
    //                }
    //            }
    //        }
    //    } else {
    //        fmt.Println("GetAllControls Warning: Trying to add already visited control: " + curr_control.Id)
    //    }
    //}

    return result

}


func GetAllEnhancements(c catalog.Catalog) []catalog.Control {

    result := make([]catalog.Control, 0)
    for _, curr_ctrl := range GetAllControls(c) {
        for _, child_ctrl := range curr_ctrl.Controls {
            result = append(result, child_ctrl)
        }

    }
    return result
}


func CreateControlstoEnhancementsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS controls_enhancements(controlid varchar(20), enhid varchar(20))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    //visited := make(map[string]bool)

    for _, curr_ctrl := range GetAllControls(c) {
        for _, child_control := range curr_ctrl.Controls {
            if child_control.Id != "" {
                query := `INSERT INTO controls_enhancements(controlid, enhid) Values("`
                query += curr_ctrl.Id
                query += `", "`
                query += child_control.Id
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

// Not Correct
func CreatePartsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS parts(partid varchar(20), name varchar(20), prose TEXT, PRIMARY KEY(partid))")
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
                query := `INSERT INTO parts(partid, name, prose) Values("`
                query += curr_part.Id
                query += `", "`
                query += curr_part.Name
                query += `", "`
                if curr_part.Prose == nil {
                    query += ""
                } else {
                    query += curr_part.Prose.Raw
                }
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

func CreateEnhancementsToPartsTable(db *sql.DB, c catalog.Catalog) {
    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `enhancements_parts`(enhid varchar(20), partid varchar(20))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    // For all Enhancements
    for _, ctrl := range GetAllEnhancements(c) {
        // For all parts that are directly children of ctrl
        for _, part := range ctrl.Parts {
            // Skipping Parts without Ids
            if part.Id != "" {
                query := "INSERT INTO `enhancements_parts`(enhid, partid) Values(\""
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


func CreateEnhancementsToParamsTable(db *sql.DB, c catalog.Catalog) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `enhancements_params`(enhid varchar(20), paramid varchar(20), PRIMARY KEY(paramid))")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }

    for _, ctrl := range GetAllEnhancements(c) {
        for _, param := range ctrl.Parameters {
            query := "INSERT INTO `enhancements_params`(enhid, paramid) Values(\""
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

func CreateNist800_53_rev4_enhancements(db *sql.DB) {

    _,err := db.Exec("CREATE TABLE IF NOT EXISTS `nist800_53_rev4_enhancements`(enh_id TEXT, low INT DEFAULT NULL, moderate INT DEFAULT NULL, high INT DEFAULT NULL)")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB tabled created successfully..")
    }
	toLoad := "800-53-rev4-controls.xml"
	data, e := ioutil.ReadFile(toLoad)
	if e != nil {
		fmt.Printf("error 1: %v\n", e)
		return
	}
	controls := Controls{}
	marshalError := xml.Unmarshal([]byte(data), &controls)
	if marshalError != nil {
		fmt.Printf("error 2: %v\n", marshalError)
		return
	}

	for _, ctrls := range controls.Control {
		for _, enh := range ctrls.ControlEnhancements.ControlEnhancement {
			ctrlId := strings.Split(enh.Number, " ")[0]
			enhId := strings.ReplaceAll(strings.Split(enh.Number, " ")[1], "(", "")
			enhId = strings.ReplaceAll(enhId, ")", "")
			enhId = strings.ToLower(ctrlId + "." + enhId)
			low, moderate, high := 0, 0, 0
			for _, impact := range enh.BaselineImpact {
				if (strings.Compare(impact, "LOW") == 0) {
					low = 1
				} else if (strings.Compare(impact, "MODERATE") == 0) {
					moderate = 1
				} else if (strings.Compare(impact, "HIGH") == 0) {
					high = 1
				}
			}
			query := `INSERT INTO nist800_53_rev4_enhancements(enh_id, low, moderate, high) Values("`
			query += enhId
			query += `", "`
			query += strconv.Itoa(low)
			query += `", "`
			query += strconv.Itoa(moderate)
			query += `", "`
			query += strconv.Itoa(high)
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

func main() {

	//toLoad := "NIST_SP-800-53_rev4_catalog.xml"
    //toLoad := "ac-1-control.json";
    //toLoad := "NIST_SP-800-53_rev4_catalog.json"
	//toLoad := "800-53-rev4-controls.xml"

	db, err := sql.Open("mysql", "infobeyond:1234@(192.168.1.124:3306)/cube")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    CreateNist800_53_rev4_enhancements(db)

    //CreateParamsToValuesTable(db)

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
    ////fmt.Println(data)

    //controls := Controls{}
	//marshalError := xml.Unmarshal([]byte(data), &controls)
	//if marshalError != nil {
	// fmt.Printf("error 2: %v\n", marshalError)
	// return
	//}
	//fmt.Println(controls.Control[1].ControlEnhancements.ControlEnhancement[0].BaselineImpact)

	//for _, val := range controls.

	////fmt.Println(data)

	//c := catalog.Catalog{}
    ////var c string

    ////var c catalog.Catalog
    //marshalError := json.Unmarshal([]byte(data), &c)
    //if marshalError != nil {
    // fmt.Printf("error 2: %v\n", marshalError)
    // return
    //}

    ////fmt.Println(data)

    ////fmt.Print(data)

    //fmt.Println(c)

    //fmt.Print(c);

    //CreatePartsTable(db, c)
    //CreateParamsTable(db, c)
    //CreateControlsToPartsTable(db, c)
    //CreateControlsToParamsTable(db, c)
    //CreatePartsToPartsTable(db, c)
    //CreatePartsToParagraphsTable(db, c)

    //CreateControlstoEnhancementsTable(db, c)
    //CreateEnhancementsToPartsTable(db, c)
    //CreateEnhancementsToParamsTable(db, c)

}
