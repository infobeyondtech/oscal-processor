package main

import (
	"database/sql"
	"github.com/docker/oscalkit/types/oscal/catalog"
	"github.com/infobeyondtech/oscal-processor/context"
	"reflect"

	//"github.com/docker/oscalkit/types/oscal/catalog"
	//"encoding/json"
	"encoding/xml"

	//"encoding/xml"
	//"github.com/docker/oscalkit/types/oscal/catalog"
	//"github.com/docker/oscalkit/types/oscal/profile"
	//"github.com/infobeyondtech/oscal-processor/models/profile"

	//"github.com/docker/oscalkit/types/oscal/catalog"
	//"github.com/docker/oscalkit/types/oscal"
	//"github.com/docker/oscalkit/types/oscal/catalog"

	"os"
	"strconv"
	"strings"

	"fmt"
	"github.com/alediaferia/stackgo"
	"io/ioutil"
	//"github.com/docker/oscalkit/types/oscal/catalog"
	//catalog "github.com/docker/oscalkit/types/oscal/catalog"
	_ "github.com/go-sql-driver/mysql"
	//"io/ioutil"
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
							P    string `xml:"p"`
							Part []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Name string `xml:"name,attr"`
								Prop struct {
									Text string `xml:",chardata"`
									Name string `xml:"name,attr"`
								} `xml:"prop"`
								P    string `xml:"p"`
								Part []struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Name string `xml:"name,attr"`
									Prop struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"prop"`
									P string `xml:"p"`
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
								P    string `xml:"p"`
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
									P string `xml:"p"`
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
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS controls_enhancements(controlid varchar(20), enhid varchar(20))")
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
				_, err = db.Exec(query)
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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS parts(partid varchar(20), name varchar(20), prose TEXT, PRIMARY KEY(partid))")
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
				_, err = db.Exec(query)
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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS params(paramid varchar(20), label varchar(300), PRIMARY KEY(paramid))")
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
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Caused by: " + query)
				os.Exit(0)
			}
		}
	}
}

func CreateControlsToPartsTable(db *sql.DB, c catalog.Catalog) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `controls_parts`(controlid varchar(20), partid varchar(20))")
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
				_, err = db.Exec(query)
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
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `enhancements_parts`(enhid varchar(20), partid varchar(20))")
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
				_, err = db.Exec(query)
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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `enhancements_params`(enhid varchar(20), paramid varchar(20), PRIMARY KEY(paramid))")
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
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Caused by: " + query)
				os.Exit(0)
			}
		}
	}
}

func CreateControlsToParamsTable(db *sql.DB, c catalog.Catalog) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `controls_params`(controlid varchar(20), paramid varchar(20), PRIMARY KEY(paramid))")
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
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Caused by: " + query)
				os.Exit(0)
			}
		}
	}
}

func CreatePartsToPartsTable(db *sql.DB, c catalog.Catalog) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `parts_parts`(`parent_partid` varchar(20), `child_partid` varchar(20))")
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
									_, err = db.Exec(query)
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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `parts_paragraph`(partid varchar(20), paragraph varchar(300))")
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
    qs := "CREATE TABLE IF NOT EXISTS `params_values`(record_id MEDIUMINT NOT NULL AUTO_INCREMENT, project_id INT, component_id VARCHAR(30), param_id VARCHAR(30), value TEXT, PRIMARY KEY (record_id));"
    fmt.Println(qs)
	_, err := db.Exec(qs)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB tabled created successfully..")
	}
}

func CreateNist800_53_rev4_enhancements(db *sql.DB) {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `nist800_53_rev4_enhancements`(enh_id TEXT, low INT DEFAULT NULL, moderate INT DEFAULT NULL, high INT DEFAULT NULL)")
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
				if strings.Compare(impact, "LOW") == 0 {
					low = 1
				} else if strings.Compare(impact, "MODERATE") == 0 {
					moderate = 1
				} else if strings.Compare(impact, "HIGH") == 0 {
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
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Caused by: " + query)
				os.Exit(0)
			}
		}
	}
}

func isEmpty(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}

func CreateControlToRelatedControlsTable(db *sql.DB, c Catalog) {
	db.Exec("CREATE TABLE IF NOT EXISTS `control_related_controls`(controlid varchar(20), relatedcontrolid varchar(20))")
	for _, group := range c.Group {
		for _, ctrl := range group.Control {
			for _, part := range ctrl.Part {
				if part.Name == "guidance" {
					//fmt.Println(ctrl.ID)
					for _, link := range part.Link {
						//fmt.Println("\t" + strings.ToLower(link.Text))
						query := `INSERT INTO control_related_controls(controlid, relatedcontrolid) Values("`
						query += ctrl.ID
						query += `", "`
						query += strings.ToLower(link.Text)
						query += `")`
						_, err := db.Exec(query)
						if err != nil {
							fmt.Println("Error on: " + query)
						}
					}
				}
			}
		}
	}

}

//
//func AddEnhancementImpact(db *sql.DB) {
//
//	//_, err := db.Exec("CREATE TABLE IF NOT EXISTS `nist800_53_rev4_enhancements`(enh_id TEXT, low INT DEFAULT NULL, moderate INT DEFAULT NULL, high INT DEFAULT NULL)")
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//} else {
//	//	fmt.Println("DB tabled created successfully..")
//	//}
//	toLoad := "800-53-rev4-controls.xml"
//	data, e := ioutil.ReadFile(toLoad)
//	if e != nil {
//		fmt.Printf("error 1: %v\n", e)
//		return
//	}
//	controls := Controls{}
//	marshalError := xml.Unmarshal([]byte(data), &controls)
//	if marshalError != nil {
//		fmt.Printf("error 2: %v\n", marshalError)
//		return
//	}
//
//	for _, ctrls := range controls.Control {
//		for _, enh := range ctrls.ControlEnhancements.ControlEnhancement {
//			for _, impact := range enh.BaselineImpact {
//				if impact == "LOW" {
//					query := `INSERT INTO control_related_controls(controlid, relatedcontrolid) Values("`
//					query += ctrl.ID
//					query += `", "`
//					query += strings.ToLower(link.Text)
//					query += `")`
//					_, err := db.Exec(query)
//
//				}
//				if impact == "MODERATE" {
//
//				}
//				if impact == "HIGH" {
//
//				}
//			}
//        }
//    }
//}

func CreateComponentsToUsersTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `components_users`(component TEXT, user TEXT)")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB tabled created successfully..")
	}
}

func CreateComponentsValues(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `components_values`(recordId MEDIUMINT NOT NULL AUTO_INCREMENT, projectId INT, statementId VARCHAR(30), componentId VARCHAR(100), PRIMARY KEY (recordId))")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB tabled created successfully..")
	}
}


func main() {
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//CreateComponentsValues(db)
    //CreateComponentsToUsersTable(db)
    //CreateParamsToValuesTable(db)
	//AddEnhancementImpact(db)

	toLoad := "NIST_SP-800-53_rev4_catalog.xml"
	//data, _ := ioutil.ReadFile(toLoad)
	//c := Catalog{}
	//marshalError := xml.Unmarshal([]byte(data), &c)
	//if marshalError != nil {
	//	fmt.Printf("error 2: %v\n", marshalError)
	//	return
	//}
	//CreateControlToRelatedControlsTable(db, c)
	//toLoad := "NIST_SP-800-53_rev4_catalog.json"
	//db, err := sql.Open("mysql", "root_master:root@(216.84.167.166:3306)/cube")
	//db, err := sql.Open("mysql", context.DBSource)
	//if err != nil {
	//    panic(err.Error())
	//}
	//defer db.Close()
	//// Open doesn't open a connection. Validate DSN data:
	//err = db.Ping()
	//if err != nil {
	//    panic(err.Error())
	//}

	//db.Exec("CREATE TABLE IF NOT EXISTS `user_context`(projectId varchar(20), profileFileId varchar(20))")

	//query := `INSERT INTO param_info(paramid, type, description) Values("`
	//query += enhId
	//query += `", "`
	//query += strconv.Itoa(low)
	//query += `", "`
	//query += strconv.Itoa(moderate)
	//query += `", "`
	//query += strconv.Itoa(high)
	//query += `")`
	//_, err = db.Exec(query)

	db.Exec("CREATE TABLE IF NOT EXISTS `param_info`(paramid varchar(20), label TEXT, sort varchar(20), description TEXT)")
	data, _ := ioutil.ReadFile(toLoad)
	c := Catalog{}
	marshalError := xml.Unmarshal([]byte(data), &c)
	if marshalError != nil {
		fmt.Printf("error 2: %v\n", marshalError)
		return
	} else {
		//var paramMap map[string]string
		paramMap := make(map[string]string)
		for _, group := range c.Group {
			for _, ctrl := range group.Control {
				for _, param := range ctrl.Param {
					paramMap[param.ID] = param.Label
				}
			}
		}
		for _, group := range c.Group {
			for _, ctrl := range group.Control {
				for _, param := range ctrl.Param {
					if !isEmpty(param.Select) {
						query := `INSERT INTO param_info(paramid, label, sort, description) Values("`
						query += param.ID
						query += `", "`
						query += param.Label
						query += `", "`
						query += `selection`
						query += `", `
						query += `'{ "HowMany": "`
						strippedHowMany := strings.ReplaceAll(param.Select.HowMany, " ", "")
						strippedHowMany = strings.ReplaceAll(strippedHowMany, "\n", "")
						strippedHowMany = strings.ReplaceAll(strippedHowMany, "\t", "")
						if strippedHowMany != "" {
							query += param.Select.HowMany
							query += `", `
						} else {
							query += `one", `
						}
						query += `"choices": [`
						firstChoice := true
						for _, c := range param.Select.Choice {
							if firstChoice {
								firstChoice = false
							} else {
								query += `, `
							}
							query += `{ "Text": "`
							strippedChoiceText := strings.ReplaceAll(c.Text, " ", "")
							strippedChoiceText = strings.ReplaceAll(strippedChoiceText, "\n", "")
							strippedChoiceText = strings.ReplaceAll(strippedChoiceText, "\t", "")
							if strippedChoiceText != "" {
								query += strings.TrimSpace(c.Text)
								query += `", `
							} else {
								query += `", `
							}
							query += `"Insert": "`
							if !isEmpty(c.Insert) {
								query += c.Insert.ParamID
								query += `", "InsertLabel": "`
								query += paramMap[c.Insert.ParamID]
								query += `"`
								//fmt.Println("insert paramid: " + c.Insert.ParamID)
							} else {
								query += `", "InsertLabel": ""`
							}
							query += ` }`
						}
						query += `] }'`
						query += `)`
						//fmt.Println(query)
						_, err = db.Exec(query)
						if err != nil {
							fmt.Print(err)
							fmt.Println()
							fmt.Println(query)
						}
					} else {
						query := `INSERT INTO param_info(paramid, label, sort, description) Values("`
						query += param.ID
						query += `", "`
						query += param.Label
						query += `", "`
						query += `assignment`
						query += `", "`
						query += `{}`
						query += `")`
						//fmt.Println(query)
						_, err = db.Exec(query)
						if err != nil {
							fmt.Print(err)
							fmt.Println()
							fmt.Println(query)
						}
					}
				}
			}
		}
	}
}

//for _, curr_ctrl := range GetAllControls(c) {
//	if curr_ctrl.Parameters != nil {
//		for _, param := range curr_ctrl.Parameters {
//			fmt.Print(param)
//			fmt.Println()
//		}
//	}
//}
//for _, group := range c.Catalog.Groups {
//	fmt.Print(group)
//    //for _, ctrl := range group.Controls {
//    //    fmt.Println(ctrl)
//    //    //for _, part := range ctrl.Parts {
//    //    //    fmt.Println(part)
//    //    //}
//    //}
//}

//toLoad := "NIST_SP-800-53_rev4_catalog.xml"
//data, _ := ioutil.ReadFile(toLoad)

//toLoad := "ac-1-control.json"; //toLoad := "NIST_SP-800-53_rev4_catalog.json"
//toLoad := "800-53-rev4-controls.xml"

//db, err := sql.Open("mysql", "infobeyond:1234@(192.168.1.124:3306)/cube")
//if err != nil {
//    panic(err.Error())
//}
//defer db.Close()

//// Open doesn't open a connection. Validate DSN data:
//err = db.Ping()
//if err != nil {
//panic(err.Error())
//}

//CreateNist800_53_rev4_enhancements(db)

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
//xml.Unmarshal([]byte(data), &controls)
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

