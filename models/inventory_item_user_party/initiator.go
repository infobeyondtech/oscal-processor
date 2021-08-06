package inventory_item_user_party

import (
	//"encoding/json"
	"database/sql"
	"strconv"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type InventoryItemUserPartyMap struct {
	Id int64 `json:"id,omitempty"`
	ProjectId  int64 `json:"projectId,omitempty"`
	InventoryItemId string `json:"inventoryItemId,omitempty"`
	UserId string `json:"userId,omitempty"`
	PartyId string `json:"partyId,omitempty"`
}

func AddInventoryItemUserParty(projectId int64, inventoryItemId string, userId string, partyId string) int64{
 	// do not check duplicate records, client can guard it
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := `INSERT INTO inventory_item_user_party (project_id, inventory_item_id, user_id, party_id) Values("`
	query +=  strconv.FormatInt(projectId, 10)
	query += `", "`
	query += inventoryItemId
	query += `", "`
	query += userId
	query += `", "`
	query += partyId
	query += `")`
	res, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	return id
}

func RemoveInventoryItemUserParty(id int64){
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := `DELETE FROM inventory_item_user_party WHERE id ='`
	query += strconv.FormatInt(id, 10)+ `'`

	_, er := db.Exec(query)
	if er != nil {
		panic(err.Error())
	}
}

func GetInventoryItemUserParty(projectid int64) []InventoryItemUserPartyMap{
	var iupmap []InventoryItemUserPartyMap

	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// use project id to search for user-party maps
	queryString := `SELECT id, project_id, inventory_item_id, user_id, party_id FROM inventory_item_component WHERE project_id = '`
	queryString += strconv.FormatInt(projectid, 10)+ `'`
	queryResult, err := db.Query(queryString)
	if err != nil {
		panic(err.Error())
	}

	// iterate over the query result
	for queryResult.Next() {
		var id int64
		var project_id int64
		var inventory_item_id string
		var user_id string
		var party_id string
		err := queryResult.Scan(&id, &project_id, &inventory_item_id, &user_id, &party_id)
		if err != nil {
			log.Fatal(err)
		}
		iupmap = append(iupmap, InventoryItemUserPartyMap{Id:id, ProjectId: project_id, InventoryItemId: inventory_item_id, UserId: user_id, PartyId: party_id})
	}

	return iupmap	
}