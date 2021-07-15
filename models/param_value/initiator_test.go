package param_value

//func TestParamValueIntegration(t *testing.T) {
//    assert := assert.New(t)
//    test_project_ids := []int{1}
//    test_component_ids := []string{"test_component_1"}
//    test_param_ids := []string{"test_param_1"}
//    test_values := []string{"value1"}
//
//}

//func TestSetParamValue(t *testing.T) {
//func TestCreateParamValue(t *testing.T) {
//    //assert := assert.New(t)
//    test_project_ids := []int{1}
//    test_component_ids := []string{"test_component_1"}
//    test_param_ids := []string{"test_param_1"}
//    test_values := []string{"value1"}
//
//    expected_values := make([]ParamValue, 0)
//    for idx, _ := range test_project_ids {
//        expected_values = append(expected_values, ParamValue{0, test_project_ids[idx],test_component_ids[idx],test_param_ids[idx],test_values[idx]})
//    }
//
//    for idx, _ := range expected_values {
//       // Get a Test ParamValue
//       //actual_pv := CreateParamValue(expected_values[idx].ProjectId, expected_values[idx].ComponentId, expected_values[idx].ParamId, expected_values[idx].Value)
//       r := CreateParamValue(expected_values[idx].ProjectId, expected_values[idx].ComponentId, expected_values[idx].ParamId, expected_values[idx].Value)
//       fmt.Println(r)
//       UpdateParamValue(r, "rawr")
//       fmt.Println(GetParamValue(r))
//       fmt.Println(GetParam(1))
//       DeleteParamValue(2)
//
//       //assert.Equal(expected_values[idx], actual_pv)
//    }
//}

//func TestGetParamValue(t *testing.T) {
//
//    assert := assert.New(t)
//    test_project_ids := []int{1}
//    test_component_ids := []string{"test_component_1"}
//    test_param_ids := []string{"test_param_1"}
//    test_values := []string{"value1"}
//
//    expected_values := make([]ParamValue, 0)
//    for idx, _ := range test_project_ids {
//        expected_values = append(expected_values, ParamValue{test_project_ids[idx],test_component_ids[idx],test_param_ids[idx],test_values[idx]})
//    }
//
//    for idx, _ := range expected_values {
//       // Get a Test ParamValue
//       actual_pv := GetParamValue(expected_values[idx].ProjectId, expected_values[idx].ComponentId, expected_values[idx].ParamId)
//       assert.Equal(expected_values[idx], actual_pv)
//    }
//}
//
//func TestGetParam(t *testing.T) {
//    fmt.Println(GetParam(1))
//}



