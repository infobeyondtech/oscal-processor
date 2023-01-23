package information

import (
	//"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/infobeyondtech/oscal-processor/context"
)

type NullableComponent struct {
	Id           sql.NullString `json:"id,omitempty"`
	Title        sql.NullString `json:"title,omitempty"`
	UUID         sql.NullString `json:"uuid,omitempty"`
	Description  sql.NullString `json:"description,omitempty"`
	State        sql.NullString `json:"state,omitempty"`
	Type         sql.NullString `json:"type,omitempty"`
	LastModified sql.NullString `json:"last_modified,omitempty"`
	Version      sql.NullString `json:"version,omitempty"`
}

type Component struct {
	Id           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Description  string `json:"description,omitempty"`
	State        string `json:"state,omitempty"`
	Type         string `json:"type,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	Version      string `json:"version,omitempty"`
	Roles        []string `json:"roles,omitempty"`
}

type NullableInventoryItem struct {
	Id          sql.NullString `json:"id,omitempty"`
	AssetId     sql.NullString `json:"asset_id,omitempty"`
	UUID        sql.NullString `json:"uuid,omitempty"`
	Description sql.NullString `json:"description,omitempty"`
}

type InventoryItem struct {
	Id          string `json:"id,omitempty"`
	AssetId     string `json:"asset_id,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
}

type NullableParty struct {
	Id     sql.NullString `json:"id,omitempty"`
	RoleId sql.NullString `json:"role_id,omitempty"`
	UUID   sql.NullString `json:"uuid,omitempty"`
	Type   sql.NullString `json:"type,omitempty"`
}

type Party struct {
	Id     string `json:"id,omitempty"`
	RoleId string `json:"role_id,omitempty"`
	UUID   string `json:"uuid,omitempty"`
	Type   string `json:"type,omitempty"`
}

type NullableUser struct {
	Id     sql.NullString `json:"id,omitempty"`
	Title  sql.NullString `json:"title,omitempty"`
	Type   sql.NullString `json:"type,omitempty"`
	RoleId sql.NullString `json:"role_id,omitempty"`
	UUID   sql.NullString `json:"uuid,omitempty"`
}

type User struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
	RoleId string `json:"role_id,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

func GetComponent(UUID string) Component {
	var result Component
	var nullableResult NullableComponent
	users := make([]string, 0)

	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//rows, err := db.Query()
		qs := `SELECT id, title, uuid, description, state, type, last_modified, version, user ` +
                          `FROM oscal_component ` +
                          `INNER JOIN components_users on oscal_component.uuid=components_users.component ` +
                          `WHERE uuid="` + UUID + `";`
		rows, err := db.Query(qs)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var nullableRole sql.NullString
		err = rows.Scan(
			&nullableResult.Id,
			&nullableResult.Title,
			&nullableResult.UUID,
			&nullableResult.Description,
			&nullableResult.State,
			&nullableResult.Type,
			&nullableResult.LastModified,
			&nullableResult.Version,
			&nullableRole,
		)
		if nullableRole.Valid {
			users = append(users, nullableRole.String)
		}
	}
	if nullableResult.Id.Valid {
		result.Id = nullableResult.Id.String
	} else {
		result.Id = ""
	}
	if nullableResult.Title.Valid {
		result.Title = nullableResult.Title.String
	} else {
		result.Title = ""
	}
	if nullableResult.UUID.Valid {
		result.UUID = nullableResult.UUID.String
	} else {
		result.UUID = ""
	}
	if nullableResult.Description.Valid {
		result.Description = nullableResult.Description.String
	} else {
		result.Description = ""
	}
	if nullableResult.State.Valid {
		result.State = nullableResult.State.String
	} else {
		result.State = ""
	}
	if nullableResult.Type.Valid {
		result.Type = nullableResult.Type.String
	} else {
		result.Type = ""
	}
	if nullableResult.LastModified.Valid {
		result.LastModified = nullableResult.LastModified.String
	} else {
		result.LastModified = ""
	}
	if nullableResult.Version.Valid {
		result.Version = nullableResult.Version.String
	} else {
		result.Version = ""
	}
	result.Roles = users
	return result
}

func GetInventoryItem(UUID string) InventoryItem {
	var result InventoryItem
	var nullableResult NullableInventoryItem
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.QueryRow(`SELECT id, asset_id, uuid, description FROM oscal_inventory_item WHERE uuid = "`+UUID+`";`).
		Scan(&nullableResult.Id, &nullableResult.AssetId, &nullableResult.UUID, &nullableResult.Description)
	if err != nil {
		panic(err.Error())
	}
	if nullableResult.Id.Valid {
		result.Id = nullableResult.Id.String
	} else {
		result.Id = ""
	}
	if nullableResult.AssetId.Valid {
		result.AssetId = nullableResult.AssetId.String
	} else {
		result.AssetId = ""
	}
	if nullableResult.UUID.Valid {
		result.UUID = nullableResult.UUID.String
	} else {
		result.UUID = ""
	}
	if nullableResult.Description.Valid {
		result.Description = nullableResult.Description.String
	} else {
		result.Description = ""
	}
	return result
}

func GetParty(UUID string) Party {
	var result Party
	var nullableResult NullableParty
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.QueryRow(`SELECT id, role_id, uuid, type FROM oscal_party WHERE uuid = "`+UUID+`";`).
		Scan(&nullableResult.Id, &nullableResult.RoleId, &nullableResult.UUID, &nullableResult.Type)
	if err != nil {
		panic(err.Error())
	}
	if nullableResult.Id.Valid {
		result.Id = nullableResult.Id.String
	} else {
		result.Id = ""
	}
	if nullableResult.RoleId.Valid {
		result.RoleId = nullableResult.RoleId.String
	} else {
		result.RoleId = ""
	}
	if nullableResult.UUID.Valid {
		result.UUID = nullableResult.UUID.String
	} else {
		result.UUID = ""
	}
	if nullableResult.Type.Valid {
		result.Type = nullableResult.Type.String
	} else {
		result.Type = ""
	}
	return result
}

func GetUser(UUID string) User {
	var result User
	var nullableResult NullableUser
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.QueryRow(`SELECT id, title, type, role_id, uuid FROM oscal_user WHERE uuid = "`+UUID+`";`).
		Scan(&nullableResult.Id, &nullableResult.Title, &nullableResult.Type, &nullableResult.RoleId, &nullableResult.UUID)
	if err != nil {
		panic(err.Error())
	}
	if nullableResult.Id.Valid {
		result.Id = nullableResult.Id.String
	} else {
		result.Id = ""
	}
	if nullableResult.Title.Valid {
		result.Title = nullableResult.Title.String
	} else {
		result.Title = ""
	}
	if nullableResult.Type.Valid {
		result.Type = nullableResult.Type.String
	} else {
		result.Type = ""
	}
	if nullableResult.RoleId.Valid {
		result.RoleId = nullableResult.RoleId.String
	} else {
		result.RoleId = ""
	}
	if nullableResult.UUID.Valid {
		result.UUID = nullableResult.UUID.String
	} else {
		result.UUID = ""
	}
	return result
}

func FindComponent(Input string) []Component {
	results := make([]Component, 0)
	componentMap := make(map[string]Component)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	qs := `SELECT id, title, uuid, description, state, type, last_modified, version, user ` +
		  `FROM oscal_component ` +
		  `INNER JOIN components_users ON oscal_component.uuid=components_users.component ` +
		  `WHERE title like "%` + Input + `%" OR description like "%` + Input + `%";`
	rows, err := db.Query(qs)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result Component
		var nullableResult NullableComponent
		var nullableRole sql.NullString
		var currRole string
		err := rows.Scan(
			&nullableResult.Id,
			&nullableResult.Title,
			&nullableResult.UUID,
			&nullableResult.Description,
			&nullableResult.State,
			&nullableResult.Type,
			&nullableResult.LastModified,
			&nullableResult.Version,
			&nullableRole,
		)
		if err != nil {
			panic(err.Error())
		}
		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.Title.Valid {
			result.Title = nullableResult.Title.String
		} else {
			result.Title = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		if nullableResult.Description.Valid {
			result.Description = nullableResult.Description.String
		} else {
			result.Description = ""
		}
		if nullableResult.State.Valid {
			result.State = nullableResult.State.String
		} else {
			result.State = ""
		}
		if nullableResult.Type.Valid {
			result.Type = nullableResult.Type.String
		} else {
			result.Type = ""
		}
		if nullableResult.LastModified.Valid {
			result.LastModified = nullableResult.LastModified.String
		} else {
			result.LastModified = ""
		}
		if nullableResult.Version.Valid {
			result.Version = nullableResult.Version.String
		} else {
			result.Version = ""
		}
		if nullableRole.Valid {
			currRole = nullableRole.String
		}
		if _, present := componentMap[result.UUID]; !present {
			result.Roles = make([]string, 0)
		}
		result.Roles = append(componentMap[result.UUID].Roles, currRole)
		componentMap[result.UUID] = result
	}

	for _, c := range componentMap {
		results = append(results, c)
	}

	return results
}

//func FindAllComponent(Input string) []Component {
//	results := make([]Component, 0)
//	db, err := sql.Open("mysql", context.DBSource)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//	qs := `SELECT id, title, uuid, description, state, type, last_modified, version, user ` +
//		`FROM oscal_component ` +
//		`INNER JOIN components_users ON oscal_component.uuid=components_users.component;`
//	rows, err := db.Query(qs)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var result Component
//		var nullableResult NullableComponent
//
//		err := rows.Scan(&nullableResult.Id, &nullableResult.Title, &nullableResult.UUID, &nullableResult.Description, &nullableResult.State, &nullableResult.Type, &nullableResult.LastModified, &nullableResult.Version)
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		if nullableResult.Id.Valid {
//			result.Id = nullableResult.Id.String
//		} else {
//			result.Id = ""
//		}
//		if nullableResult.Title.Valid {
//			result.Title = nullableResult.Title.String
//		} else {
//			result.Title = ""
//		}
//		if nullableResult.UUID.Valid {
//			result.UUID = nullableResult.UUID.String
//		} else {
//			result.UUID = ""
//		}
//		if nullableResult.Description.Valid {
//			result.Description = nullableResult.Description.String
//		} else {
//			result.Description = ""
//		}
//		if nullableResult.State.Valid {
//			result.State = nullableResult.State.String
//		} else {
//			result.State = ""
//		}
//		if nullableResult.Type.Valid {
//			result.Type = nullableResult.Type.String
//		} else {
//			result.Type = ""
//		}
//		if nullableResult.LastModified.Valid {
//			result.LastModified = nullableResult.LastModified.String
//		} else {
//			result.LastModified = ""
//		}
//		if nullableResult.Version.Valid {
//			result.Version = nullableResult.Version.String
//		} else {
//			result.Version = ""
//		}
//		results = append(results, result)
//	}
//	return results
//}

func FindInventoryItem(Input string) []InventoryItem {
	results := make([]InventoryItem, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, asset_id, uuid, description FROM oscal_inventory_item WHERE asset_id like "%` + Input + `%" OR description like "% ` + Input + `%";`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result InventoryItem
		var nullableResult NullableInventoryItem

		err := rows.Scan(&nullableResult.Id, &nullableResult.AssetId, &nullableResult.UUID, &nullableResult.Description)

		if err != nil {
			panic(err.Error())
		}

		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.AssetId.Valid {
			result.AssetId = nullableResult.AssetId.String
		} else {
			result.AssetId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		if nullableResult.Description.Valid {
			result.Description = nullableResult.Description.String
		} else {
			result.Description = ""
		}
		results = append(results, result)
	}
	return results
}

func FindAllInventoryItem(Input string) []InventoryItem {
	results := make([]InventoryItem, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, asset_id, uuid, description FROM oscal_inventory_item;`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result InventoryItem
		var nullableResult NullableInventoryItem

		err := rows.Scan(&nullableResult.Id, &nullableResult.AssetId, &nullableResult.UUID, &nullableResult.Description)

		if err != nil {
			panic(err.Error())
		}

		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.AssetId.Valid {
			result.AssetId = nullableResult.AssetId.String
		} else {
			result.AssetId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		if nullableResult.Description.Valid {
			result.Description = nullableResult.Description.String
		} else {
			result.Description = ""
		}
		results = append(results, result)
	}
	return results
}

func FindParty(Input string) []Party {
	results := make([]Party, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, role_id, uuid, type FROM oscal_party WHERE role_id like "%` + Input + `%";`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result Party
		var nullableResult NullableParty

		err := rows.Scan(&nullableResult.Id, &nullableResult.RoleId, &nullableResult.UUID, &nullableResult.Type)

		if err != nil {
			panic(err.Error())
		}

		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.RoleId.Valid {
			result.RoleId = nullableResult.RoleId.String
		} else {
			result.RoleId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		if nullableResult.Type.Valid {
			result.Type = nullableResult.Type.String
		} else {
			result.Type = ""
		}
		results = append(results, result)
	}
	return results
}

func FindAllParty(Input string) []Party {
	results := make([]Party, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, role_id, uuid, type FROM oscal_party;`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result Party
		var nullableResult NullableParty
		err := rows.Scan(&nullableResult.Id, &nullableResult.RoleId, &nullableResult.UUID, &nullableResult.Type)
		if err != nil {
			panic(err.Error())
		}
		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.RoleId.Valid {
			result.RoleId = nullableResult.RoleId.String
		} else {
			result.RoleId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		if nullableResult.Type.Valid {
			result.Type = nullableResult.Type.String
		} else {
			result.Type = ""
		}
		results = append(results, result)
	}
	return results
}

func FindUser(Input string) []User {
	results := make([]User, 0)
	db, err := sql.Open("mysql", context.DBSource)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, title, type, role_id, uuid FROM oscal_user WHERE title like '%` + Input + `%' OR role_id like '%` + Input + `%';`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var result User
		var nullableResult NullableUser

		err := rows.Scan(&nullableResult.Id, &nullableResult.Title, &nullableResult.Type, &nullableResult.RoleId, &nullableResult.UUID)

		if err != nil {
			panic(err.Error())
		}

		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.Title.Valid {
			result.Title = nullableResult.Title.String
		} else {
			result.Title = ""
		}
		if nullableResult.Type.Valid {
			result.Type = nullableResult.Type.String
		} else {
			result.Type = ""
		}
		if nullableResult.RoleId.Valid {
			result.RoleId = nullableResult.RoleId.String
		} else {
			result.RoleId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}
		results = append(results, result)
	}
	return results
}

func FindAllUser(Input string) []User {
	results := make([]User, 0)
	db, err := sql.Open("mysql", context.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, title, type, role_id, uuid FROM oscal_user;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result User
		var nullableResult NullableUser
		err := rows.Scan(&nullableResult.Id, &nullableResult.Title, &nullableResult.Type, &nullableResult.RoleId, &nullableResult.UUID)
		if err != nil {
			panic(err.Error())
		}

		if nullableResult.Id.Valid {
			result.Id = nullableResult.Id.String
		} else {
			result.Id = ""
		}
		if nullableResult.Title.Valid {
			result.Title = nullableResult.Title.String
		} else {
			result.Title = ""
		}
		if nullableResult.Type.Valid {
			result.Type = nullableResult.Type.String
		} else {
			result.Type = ""
		}
		if nullableResult.RoleId.Valid {
			result.RoleId = nullableResult.RoleId.String
		} else {
			result.RoleId = ""
		}
		if nullableResult.UUID.Valid {
			result.UUID = nullableResult.UUID.String
		} else {
			result.UUID = ""
		}

		results = append(results, result)
	}
	return results
}
