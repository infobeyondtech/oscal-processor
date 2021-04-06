package information

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

func TestGetComponent(t *testing.T) {
    assert := assert.New(t)
    // Get the expected Component
    expectedData, expectedError := ioutil.
        ReadFile("test_files/components/795533ab-9427-4abe-820f-0b571bacfe6d.json")
    if expectedError != nil {
        fmt.Printf("TestGetComponent error: %v\n", expectedError)
    }
    // Marshall expected Component
    expectedComponent := Component{}
    marshalError := json.Unmarshal([]byte(expectedData), &expectedComponent)
    if marshalError != nil {
        fmt.Printf("TestGetComponent error: %v\n", marshalError)
    }
    // Get the actual Component and assert that it is equal to what is expected
    actualComponent := GetComponent("795533ab-9427-4abe-820f-0b571bacfe6d")
    assert.Equal(expectedComponent, actualComponent)
}

func TestGetInventoryItem(t *testing.T) {
    assert := assert.New(t)
    // Get the expected InventoryItem
    expectedData, expectedError := ioutil.ReadFile("test_files/inventory_items/c9c32657-a0eb-4cf2-b5c1-20928983063c.json")
    if expectedError != nil {
        fmt.Printf("TestGetInventoryItem error: %v\n", expectedError)
    }
    // Marshall expected InventoryItem
    expectedItem := InventoryItem{}
    marshalError := json.Unmarshal([]byte(expectedData), &expectedItem)
    if marshalError != nil {
        fmt.Printf("TestGetInventoryItem error: %v\n", marshalError)
    }
    // Get the actual InventoryItem and assert that it is equal to what is expected
    actualItem := GetInventoryItem("c9c32657-a0eb-4cf2-b5c1-20928983063c")
    assert.Equal(expectedItem, actualItem)
}

func TestGetParty(t *testing.T) {
    assert := assert.New(t)
    // Get the expected Party
    expectedData, expectedError := ioutil.ReadFile("test_files/partys/3b2a5599-cc37-403f-ae36-5708fa804b27.json")
    if expectedError != nil {
        fmt.Printf("TestGetParty error: %v\n", expectedError)
    }
    // Marshall expected Party
    expectedParty := Party{}
    marshalError := json.Unmarshal([]byte(expectedData), &expectedParty)
    if marshalError != nil {
        fmt.Printf("TestGetParty error: %v\n", marshalError)
    }
    // Get the actual Party and assert that it is equal to what is expected
    actualParty := GetParty("3b2a5599-cc37-403f-ae36-5708fa804b27")
    assert.Equal(expectedParty, actualParty)
}

func TestGetUser(t *testing.T) {
    assert := assert.New(t)
    // Get the expected User
    expectedData, expectedError := ioutil.ReadFile("test_files/users/9824089b-322c-456f-86c4-4111c4200f69.json")
    if expectedError != nil {
        fmt.Printf("TestGetUser error: %v\n", expectedError)
    }
    // Marshall expected User
    expectedUser := User{}
    marshalError := json.Unmarshal([]byte(expectedData), &expectedUser)
    if marshalError != nil {
        fmt.Printf("TestGetParty error: %v\n", marshalError)
    }
    // Get the actual User and assert that it is equal to what is expected
    actualUser := GetUser("9824089b-322c-456f-86c4-4111c4200f69")
    assert.Equal(expectedUser, actualUser)
}
