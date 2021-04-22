package profile_navigator

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "fmt"
    sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
    "encoding/xml"
    "encoding/json"
)

func TestCreateProfileNavigator(t *testing.T) {

    assert := assert.New(t)

    test_file_dir := "test_files/SP800-53/rev4/xml/"
    profile_files := []string{"NIST_SP-800-53_rev4_HIGH-baseline_profile.xml", 
                              "NIST_SP-800-53_rev4_MODERATE-baseline_profile.xml",
                              "NIST_SP-800-53_rev4_LOW-baseline_profile.xml"}

    expected_file_dir := "test_files/SP800-53/rev4/expected_json/"
    expected_pn_files := []string{"NIST_SP-800-53_rev4_HIGH-baseline_profile_navigator.json", 
                              "NIST_SP-800-53_rev4_MODERATE-baseline_profile_navigator.json",
                              "NIST_SP-800-53_rev4_LOW-baseline_profile_navigator.json"}

    for test_ct, pf := range profile_files {
        // Get a Profile
        profile_data, profile_e := ioutil.ReadFile(test_file_dir + pf)
        if profile_e != nil {
            fmt.Printf("TestCreateProfileNavigator error: %v\n", profile_e)
        }
        // Marshall the Profile
        p := sdk_profile.Profile{}
        profile_marshalError := xml.Unmarshal([]byte(profile_data), &p)
        if profile_marshalError != nil {
            fmt.Printf("TestCreateProfileNavigator error: %v\n", profile_marshalError)
        }
        // Get the Profile's actual ProfileNavigator
        pn := ProfileNavigator{}
        CreateProfileNavigator(&pn, &p)

        // Get the expected ProfileNavigator
        expected_data, expected_e := ioutil.ReadFile(expected_file_dir + expected_pn_files[test_ct])
        if expected_e != nil {
            fmt.Printf("TestCreateProfileNavigator error: %v\n", profile_e)
        }
        // Marshall expected ProfileNavigator
        expected_pn := ProfileNavigator{}
        expected_marshalError := json.Unmarshal([]byte(expected_data), &expected_pn)
        if expected_marshalError != nil {
            fmt.Printf("TestCreateProfileNavigator erro: %v\n", profile_marshalError)
        }

        assert.Equal(expected_pn, pn)
    }
}



