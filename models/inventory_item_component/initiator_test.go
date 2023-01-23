package inventory_item_component

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAddInventoryItemComponent(t *testing.T) {
	assert := assert.New(t)
	project_id := int64(4096)
	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	component_id := "4938767c-dd8b-4ea4-b74a-fafffd48ac99"
	id:= AddInventoryItemComponent(project_id, inventory_item_id, component_id)

    assert.Greater(id, 0)
}

func TestRemoveInventoryItemComponent(t *testing.T) {
	assert := assert.New(t)

	project_id := int64(4097)
	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	component_id := "4938767c-dd8b-4ea4-b74a-fafffd48ac99"
	id:= AddInventoryItemComponent(project_id, inventory_item_id, component_id)

	RemoveInventoryItemComponent(id)
	icMap := GetInventoryItemComponent(project_id)

	assert.Equal(len(icMap), 0)
}

func TestGetInventoryItemComponent(t *testing.T){
	assert := assert.New(t)

	project_id := int64(4098)
	inventory_item_id := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	component_id := "4938767c-dd8b-4ea4-b74a-fafffd48ac99"
	AddInventoryItemComponent(project_id, inventory_item_id, component_id)
	
	icMap := GetInventoryItemComponent(project_id)
	assert.Greater(len(icMap), 0)
}




