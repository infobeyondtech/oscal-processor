package param_value

import (
    "testing"
    //"encoding/json"
    "github.com/stretchr/testify/assert"
    //"io/ioutil"
    //"fmt"
    //sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
    //"encoding/xml"
    //"encoding/json"
)

func TestGetParamValue(t *testing.T) {

    assert := assert.New(t)
    test_fileids := []string{"fileid1"}
    test_paramids := []string{"paramid1"}

    expected_values := []ParamValue{ParamValue{"fileid1", "paramid1", "value1"}}

    for test_ct, fileid := range test_fileids {
       // Get a Test ParamValue
       actual_pv := GetParamValue(fileid, test_paramids[test_ct])
       assert.Equal(expected_values[test_ct], actual_pv)
    }
}

func TestSetParamValue(t *testing.T) {

    assert := assert.New(t)
    test_fileids := []string{"fileid1"}
    test_paramids := []string{"paramid1"}
    test_values := []string{"value1"}

    expected_values := []ParamValue{ParamValue{"fileid1", "paramid1", "value1"}}

    for test_ct, fileid := range test_fileids {
       // Get a Test ParamValue
       actual_pv := SetParamValue(fileid, test_paramids[test_ct], test_values[test_ct])
       assert.Equal(expected_values[test_ct], actual_pv)
    }
}



