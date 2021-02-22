package profile_navigator

import (
    //"encoding/xml"
    //"encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	"github.com/infobeyondtech/oscal-processor/context"
)

type ProfileNavigator struct {
    Groups map[string]*Group `json:"groups,omitempty"`
}

type Group struct {
    CtrlIds []string `json:"ctrlids,omitempty"`
}

func getControlGroupName(control_id string) string {
    var group_name string
    query := "Select families.description "
    query += "From controls "
    query += "INNER JOIN families "
    query += "ON controls.families_id = families.id "
    query += "WHERE controls.name = \""
    query += control_id
    query += "\""
    err := context.DB.QueryRow(query).Scan(&group_name)
    if err != nil {
     fmt.Println(err.Error())
     fmt.Println(query)
     fmt.Println()
    }
    return group_name
}

func (pn ProfileNavigator) addControl(control_id string) {
    group_name := getControlGroupName(control_id)
    if group_name != "" {
        _, in_group := pn.Groups[group_name]
        if !in_group {
            pn.Groups[group_name] = new(Group)
            pn.Groups[group_name].CtrlIds = make([]string, 0)
        }
        pn.Groups[group_name].CtrlIds = append(pn.Groups[group_name].CtrlIds, control_id)
    }
}

func getProfileControlIds(p *sdk_profile.Profile) []string {
    var result []string
    // For each Import in the Profile
    for _, i := range p.Imports {
        // Get the Import's include
        inc := i.Include
        // For each call/control Id in the include
        for _, c := range inc.IdSelectors {
            result = append(result, c.ControlId)
        }
    }
    return result
}

func CreateProfileNavigator(profile_navigator *ProfileNavigator, profile *sdk_profile.Profile) {

    // Get the database handle
    var err error
    context.DB, err = sql.Open("mysql", "infobeyond:1234@(192.168.1.124:3306)/cube")
    if err != nil {
        panic(err.Error())
    }

    // Open doesn't open a connection. Validate DSN data:
    err = context.DB.Ping()
    if err != nil {
        panic(err.Error())
    }

    // Initialize the pn's groups and all of the 
    // profile's controls to them
    profile_navigator.Groups = make(map[string]*Group)
    ctrl_ids := getProfileControlIds(profile)
    for _, id := range ctrl_ids {
        profile_navigator.addControl(id)
    }

    context.DB.Close()
}

func (pn ProfileNavigator) Print() {
    for group_name, group := range pn.Groups{
        fmt.Println(group_name)
        for _, ctrl := range group.CtrlIds {
            fmt.Println("\t"+ctrl)
        }
        fmt.Println()
    }
}


