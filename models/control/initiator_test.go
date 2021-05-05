package control

import (
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "sort"
    "testing"
)

func TestGetControl(t *testing.T) {

    assert := assert.New(t)
    testControls := []string{"ac-1",
                              "ac-2",
                              "ca-2",
                              "cp-8",
                              "cp-9"}

    expectedFileDir := "test_files/controls/"
    expectedCtrlFiles := []string{"ac-1.json",
                              "ac-2.json",
                              "ca-2.json",
                              "cp-8.json",
                              "cp-9.json"}

    for testCt, ctrl := range testControls {
       // Get a Test Control and Marshal it to JSON
       actualCtrl := GetControl(ctrl, false)
       actualCtrlJson, err1 := json.Marshal(actualCtrl)
       if err1 != nil {
           fmt.Printf("TestGetControl error 1: %v\n", err1)
       }

       // Read the expected Control from test_files/
       expectedData, err2 := ioutil.ReadFile(expectedFileDir + expectedCtrlFiles[testCt])
       if err2 != nil {
           fmt.Printf("TestGetControl error 2: %v\n", err2)
       }
       // Unmarshall expected Control and Marshal it to JSON
       expectedCtrl := Control{}
       err3 := json.Unmarshal([]byte(expectedData), &expectedCtrl)
       if err3 != nil {
           fmt.Printf("TestGetControl error 3: %v\n", err3)
       }
       expectedCtrlJson, err4 := json.Marshal(expectedCtrl)
       if err4 != nil {
           fmt.Printf("TestGetControl error 4: %v\n", err4)
       }
       assert.Equal(expectedCtrlJson, actualCtrlJson)
    }
}

func TestGetEnhancement(t *testing.T) {

    assert := assert.New(t)
    testEnhancements := []string{"ac-1.1",
                              "ac-2.2",
                              "ca-2.1",
                              "cp-8.5",
                              "cp-9.3"}

    expectedFileDir := "test_files/enhancements/"
    expectedEnhFiles := []string{"ac-1.1.json",
                              "ac-2.2.json",
                              "ca-2.1.json",
                              "cp-8.5.json",
                              "cp-9.3.json"}

    for testCt, enh := range testEnhancements {
       // Get a Test Control and Marshal it to JSON
       actualEnh := GetControl(enh, true)
       actualEnhJson, err1 := json.Marshal(actualEnh)
       if err1 != nil {
           fmt.Printf("TestGetEnhancement error 1: %v\n", err1)
       }
       // Read the expected Control from test_files/
       expectedData, err2 := ioutil.ReadFile(expectedFileDir + expectedEnhFiles[testCt])
       if err2 != nil {
           fmt.Printf("TestGetEnhancement error 2: %v\n", err2)
       }
       // Unmarshall expected Control and Marshal it to JSON
       expectedEnh := Control{}
       err3 := json.Unmarshal([]byte(expectedData), &expectedEnh)
       if err3 != nil {
           fmt.Printf("TestGetEnhancement error 3: %v\n", err3)
       }
       expectedEnhJson, err4 := json.Marshal(expectedEnh)
       if err4 != nil {
           fmt.Printf("TestGetEnhancement error 4: %v\n", err4)
       }
       assert.Equal(expectedEnhJson, actualEnhJson)
    }
}

func TestGetControlEnhancement(t *testing.T) {

    assert := assert.New(t)
    testEnhancements := []string{"ac-1.1",
                              "ac-2.2",
                              "ca-2.1",
                              "cp-8.5",
                              "cp-9.3"}

    expectedFileDir := "test_files/enhancements/"
    expectedEnhFiles := []string{"ac-1.1.json",
                              "ac-2.2.json",
                              "ca-2.1.json",
                              "cp-8.5.json",
                              "cp-9.3.json"}

    for testCt, enh := range testEnhancements {
       // Get a Test Control and Marshal it to JSON
       actualEnh := GetControl(enh, true)
       actualEnhJson, err1 := json.Marshal(actualEnh)
       if err1 != nil {
           fmt.Printf("TestGetControlEnhancement error 1: %v\n", err1)
       }
       // Read the expected Control from test_files/
       expectedData, err2 := ioutil.ReadFile(expectedFileDir + expectedEnhFiles[testCt])
       if err2 != nil {
           fmt.Printf("TestGetControlEnhancement error 2: %v\n", err2)
       }
       // Unmarshall expected Control and Marshal it to JSON
       expectedEnh := Control{}
       err3 := json.Unmarshal([]byte(expectedData), &expectedEnh)
       if err3 != nil {
           fmt.Printf("TestGetControlEnhancement error 3: %v\n", err3)
       }
       expectedEnhJson, err4 := json.Marshal(expectedEnh)
       if err4 != nil {
           fmt.Printf("TestGetControlEnhancement error 4: %v\n", err4)
       }
       assert.Equal(expectedEnhJson, actualEnhJson)
    }
}

func TestGetControlEnhancementIds(t *testing.T) {
    assert := assert.New(t)
    testCtrlIds := []string{
        "ac-1",
        "ac-2",
        "ca-2",
        "cp-8",
        "cp-9",
    }

    expectedEnhSets := [][]string{
        {},
        {"ac-2.1", "ac-2.2", "ac-2.3", "ac-2.4", "ac-2.5", "ac-2.6", "ac-2.7", "ac-2.8", "ac-2.9", "ac-2.10", "ac-2.11", "ac-2.12", "ac-2.13"},
        {"ca-2.1", "ca-2.2", "ca-2.3"},
        {"cp-8.1", "cp-8.2", "cp-8.3", "cp-8.4", "cp-8.5"},
        {"cp-9.1", "cp-9.2", "cp-9.3", "cp-9.4", "cp-9.5", "cp-9.6", "cp-9.7"},
    }

    for testCt, ctrlId := range testCtrlIds {
        actualEnhSet := GetControlEnhancementIds(ctrlId)
        expectedEnhSet := expectedEnhSets[testCt]
        sort.Strings(actualEnhSet)
        sort.Strings(expectedEnhSet)
        assert.Equal(expectedEnhSet, actualEnhSet)
    }
}

func TestSearchPartsWithKeyword(t *testing.T) {
    assert := assert.New(t)
    testKeywords := []string{"spam", "documenting", "qsdfalkjsdf"}
    expectedControlSets := [][]string{{"si-8"}, {"au-6", "au-10", "ir-5", "mp-6", "pl-8"}, {}}
    for testCt, keyword := range testKeywords {
        actualControlSet := SearchPartsWithKeyword(keyword)
        expectedControlSet := expectedControlSets[testCt]
        sort.Strings(expectedControlSet)
        sort.Strings(actualControlSet)
        assert.Equal(expectedControlSet, actualControlSet)
    }
}
