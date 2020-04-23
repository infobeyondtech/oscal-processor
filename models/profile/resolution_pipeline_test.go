package profile

import "testing"

// Test
func TestRun(t *testing.T) {
	rules := "/home/tom/oscal_workspace/OSCAL/src/utils/util/resolver-pipeline/oscal-profile-resolve-select.xsl"
	jarPath := "/home/tom/.nanshiie_baker/jars/saxon-he-10.0.jar"
	dir := "/home/tom/oscal_processing_space"
	input := "/home/tom/oscal_workspace/OSCAL/src/utils/util/resolver-pipeline/testing/pathological-profile.xml"
	output := dir + "/output/selected.xml"
	e := ResolveProfile(jarPath, rules, input, output)
	t.Error(e == nil)
}
