package component_user_party

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetComponentUserPartyMap(t *testing.T) {
	assert := assert.New(t)

	projectid := 1
	userPartyMap := GetComponentUserParty(int64(projectid))

	//fmt.Printf("length: "+ string(len(userPartyMap)))
	for _, record:= range userPartyMap{
		fmt.Printf("component id: %v, user id: %v, party id: %v\n", record.ComponentId, record.UserId, record.PartyId)
	}
	assert.Greater(len(userPartyMap), 0)
	
}

func TestAddComponentUserPartyMap(t *testing.T){
	assert := assert.New(t)

	projectid := int64(2048)
	// entities have to exist in DB
	componentid := "e00acdcf-911b-437d-a42f-b0b558cc4f03"
	userid:= "9824089b-322c-456f-86c4-4111c4200f69"
	partyid:= "eb8caad9-5fdb-46cb-b6b6-44d361ad9b5f"

    id := AddComponentUserParty(projectid, componentid, userid, partyid)
	assert.Greater(id, int64(0))
}

func TestDeleteComponentUserPartyMap(t *testing.T){
	assert := assert.New(t)

	projectid := int64(2049)	
	// entities have to exist in DB
	componentid := "e00acdcf-911b-437d-a42f-b0b558cc4f03"
	userid:= "9824089b-322c-456f-86c4-4111c4200f69"
	partyid:= "eb8caad9-5fdb-46cb-b6b6-44d361ad9b5f"

    id := AddComponentUserParty(projectid, componentid, userid, partyid)
	assert.Greater(id, int64(0))

	// remove record and examine again
	RemoveComponentUserParty(id)
	userPartyMap := GetComponentUserParty(projectid)

	assert.Equal(len(userPartyMap), 0)
}