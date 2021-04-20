package control

import (
    "testing"
    //"encoding/json"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "fmt"
    //sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
    //"encoding/xml"
    "encoding/json"
)

func TestGetControl(t *testing.T) {

    assert := assert.New(t)
    test_controls := []string{"ac-1",
                              "ac-2",
                              "ca-2",
                              "cp-8",
                              "cp-9"}

    expected_file_dir := "test_files/"
    expected_ctrl_files := []string{"ac-1.json",
                              "ac-2.json",
                              "ca-2.json",
                              "cp-8.json",
                              "cp-9.json"}

    for test_ct, cf := range test_controls {
       // Get a Test Control and Marshal it to JSON
       actual_c := GetControl(cf)
       actual_c_json, err1 := json.Marshal(actual_c)
       if err1 != nil {
           fmt.Printf("TestGetControl error 1: %v\n", err1)
       }

       // Read the expected Control from test_files/
       expected_data, err2 := ioutil.ReadFile(expected_file_dir + expected_ctrl_files[test_ct])
       if err2 != nil {
           fmt.Printf("TestGetControl error 2: %v\n", err2)
       }
       // Unmarshall expected Control and Marshal it to JSON
       expected_ctrl := Control{}
       err3 := json.Unmarshal([]byte(expected_data), &expected_ctrl)
       if err3 != nil {
           fmt.Printf("TestGetControl error 3: %v\n", err3)
       }
       expected_ctrl_json, err4 := json.Marshal(expected_ctrl)
       if err4 != nil {
           fmt.Printf("TestGetControl error 4: %v\n", err4)
       }
       assert.Equal(expected_ctrl_json, actual_c_json)
    }
}



