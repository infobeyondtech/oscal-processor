package inventory_item_user_party

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAddInventoryItemUserParty(t *testing.T){
	assert := assert.New(t)
	project_id := int64(8192)

	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	userid:= "9824089b-322c-456f-86c4-4111c4200f69"
	partyid:= "eb8caad9-5fdb-46cb-b6b6-44d361ad9b5f"

	id:= AddInventoryItemUserParty(project_id, inventory_item_id, userid, partyid)

	assert.Greater(id, 0)
}

func TestRemoveInventoryItemUserParty(t *testing.T){
	assert := assert.New(t)
	project_id := int64(8193)

	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	userid:= "9824089b-322c-456f-86c4-4111c4200f69"
	partyid:= "eb8caad9-5fdb-46cb-b6b6-44d361ad9b5f"

	id:= AddInventoryItemUserParty(project_id, inventory_item_id, userid, partyid)
	assert.Greater(id, 0)

	RemoveInventoryItemUserParty(id)

	iiupMap := GetInventoryItemUserParty(project_id)

	assert.Equal(len(iiupMap), 0)
}

func TestGetInventoryItemUserParty(t *testing.T){
	assert := assert.New(t)
	project_id := int64(8194)

	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	userid:= "9824089b-322c-456f-86c4-4111c4200f69"
	partyid:= "eb8caad9-5fdb-46cb-b6b6-44d361ad9b5f"

	id:= AddInventoryItemUserParty(project_id, inventory_item_id, userid, partyid)
	assert.Greater(id, 0)

	iiupMap := GetInventoryItemUserParty(project_id)

	assert.Greater(len(iiupMap), 0)
}

