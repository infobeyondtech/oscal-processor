package profile

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/infobeyondtech/oscal-processor/context"
)

// Test
func TestResolveProfile(t *testing.T) {
	rules := context.OSCALRepo +
		"/src/utils/util/resolver-pipeline/oscal-profile-resolve-select.xsl"
	rules = context.ExpandPath(rules)

	jarPath := context.JarLibDir + "/saxon-he-10.0.jar"
	jarPath = context.ExpandPath(jarPath)

	output := context.TempDir
	id := uuid.New().String()
	output = output + "/" + id + ".xml"
	output = context.ExpandPath(output)

	input := context.OSCALRepo +
		"/src/utils/util/resolver-pipeline/testing/pathological-profile.xml"
	input = context.ExpandPath(input)

	e := ResolveProfile(jarPath, rules, input, output)
	assert.Nil(t, e)
}
