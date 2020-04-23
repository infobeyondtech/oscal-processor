package profile

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test
func TestResolveProfile(t *testing.T) {
	rules := "/home/tom/oscal_workspace/OSCAL/src/utils/util/resolver-pipeline/oscal-profile-resolve-select.xsl"
	jarPath := "/home/tom/.nanshiie_baker/jars/saxon-he-10.0.jar"
	dir := "/home/tom/oscal_processing_space"
	input := "/home/tom/oscal_workspace/OSCAL/src/utils/util/resolver-pipeline/testing/pathological-profile.xml"
	id := uuid.New().String()
	output := dir + "/output/" + id + ".xml"
	e := ResolveProfile(jarPath, rules, input, output)
	assert.Nil(t, e)
}
