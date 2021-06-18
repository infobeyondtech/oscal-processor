package profile

import (
	"fmt"
	"testing"

	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	"github.com/stretchr/testify/assert"
)

func TestMakeProfileModel(t *testing.T) {
	assert := assert.New(t)
	// set title version in metadata
	profile := &sdk_profile.Profile{}

	title := "Infobeyond BASELINE"
	version := "1.0"
	oscalVersion := "3.0"

	// request := request_models.SetTitleVersionRequest{Title: title, Version: version, OscalVersion: oscalVersion}
	SetTitleVersion(profile, title, version, oscalVersion)

	control := []string{"cp-1", "cp-3", "cp-5"}

	AddControls(profile, control, "../../nist.gov/SP800-53/rev4/xml/NIST_SP-800-53_rev4_MODERATE-baseline_profile.xml")

	id := "catalog"
	desc := "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations"
	link := "NIST_SP-800-53_rev4_catalog.xml"
	media := "application/oscal.catalog+xml"

	AddBackMatter(profile, id, desc, link, media)

	// write to file
	path := WriteToFile(profile)
	fmt.Printf("file path: " + path)

	profileModel := MakeProfileModel(path)
	// assert.Equal(profileModel.Metadata.Title, "1")
	assert.Equal(profileModel.Metadata.Title, string(profile.Metadata.Title))
	assert.Equal(profileModel.Metadata.Version, string(profile.Metadata.Version))
	assert.Equal(profileModel.Metadata.OscalVersion, string(profile.Metadata.OscalVersion))
	//assert.Equal(profileModel.Metadata.Version, "1")
	assert.Equal(profileModel.Metadata.Title, "a")
	assert.Equal(profileModel.Metadata.Version, "a")
	assert.Equal(profileModel.Metadata.OscalVersion, "a")
	// assert.Equal(profileModel.Metadata.LastModified, string(profile.Metadata.LastModified))
	// assert.Equal(profileModel.Metadata.Roles, string(profile.Metadata.Roles[0]))
	// assert.Equal(profileModel.Metadata.Title, string(profile.Metadata.Title))

}
