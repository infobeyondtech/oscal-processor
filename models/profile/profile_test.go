package profile

import (
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
    

    profileModel := MakeProfileModel(path)
    // assert.Equal(profileModel.Metadata.Title, "1")
    assert.Equal(profileModel.Metadata.Title, string(profile.Metadata.Title))
    assert.Equal(profileModel.Metadata.Version, string(profile.Metadata.Version))
    assert.Equal(profileModel.Metadata.OscalVersion, string(profile.Metadata.OscalVersion))

}


func TestGetProfileBaselineDiff(t *testing.T) { 
	controls := []string{"AC-1", "AC-12"} 
	ans := []string{"AC-13","AC-14","AC-15","AC-16","AC-18","AC-19","AC-2","AC-20","AC-21","AC-22","AC-23","AC-24","AC-25","AC-26","AC-27","AC-28","AC-3","AC-30","AC-31","AC-33","AC-34","AC-35","AC-36","AC-37","AC-38","AC-39","AC-40","AC-42","AC-43","AC-44","AC-46","AC-48","AC-49","AC-50","AC-52","AC-53","AC-54","AC-55","AC-56","AC-57","AC-61","AC-62","AC-63","AC-64","AC-66","AC-67","AC-68","AC-69","AC-7","AC-8","IA-8","IR-1","IR-2","IR-4","IR-5","IR-6","IR-7","IR-8","MA-1","MA-2","MA-4","MA-5","MP-1","MP-2","MP-6","MP-7","PE-1","PE-12","PE-13","PE-14","PE-15","PE-16","PE-2","PE-3","PE-6","PE-8","PL-1","PL-2","PL-4","PS-1","PS-2","PS-3","PS-4","PS-5","PS-6","PS-7","PS-8","RA-1","RA-2","RA-3","RA-5","SA-1","SA-2","SA-3","SA-4","SA-5","SA-9","SC-1","SC-12","SC-13","SC-15","SC-20","SC-21","SC-22","SC-39","SC-5","SC-7","SI-1","SI-12","SI-2","SI-3","SI-4","SI-5", "ia-2.1", "ia-2.12","ia-5.1","ia-5.11","ia-8.1","ia-8.2","ia-8.3","ia-8.4","sa-4.10"}
	res := GetProfileBaselineDiff(controls , "low")
	assert.Equal(t, res, ans) 

	controls = []string{"AC-1", "AC-12"} 
	ans = []string{"AC-10","AC-11","AC-13","AC-14","AC-15","AC-16","AC-17","AC-18","AC-19","AC-2","AC-20","AC-21","AC-22","AC-23","AC-24","AC-25","AC-26","AC-27","AC-28","AC-29","AC-3","AC-30","AC-31","AC-33","AC-34","AC-35","AC-36","AC-37","AC-38","AC-39","AC-4","AC-40","AC-42","AC-43","AC-44","AC-45","AC-46","AC-47","AC-48","AC-49","AC-5","AC-50","AC-51","AC-52","AC-53","AC-54","AC-55","AC-56","AC-57","AC-58","AC-59","AC-6","AC-60","AC-61","AC-62","AC-63","AC-64","AC-65","AC-66","AC-67","AC-68","AC-69","AC-7","AC-8","IA-8","IR-1","IR-2","IR-3","IR-4","IR-5","IR-6","IR-7","IR-8","MA-1","MA-2","MA-3","MA-4","MA-5","MA-6","MP-1","MP-2","MP-3","MP-4","MP-5","MP-6","MP-7","PE-1","PE-10","PE-11","PE-12","PE-13","PE-14","PE-15","PE-16","PE-17","PE-2","PE-3","PE-4","PE-5","PE-6","PE-8","PE-9","PL-1","PL-2","PL-4","PL-8","PS-1","PS-2","PS-3","PS-4","PS-5","PS-6","PS-7","PS-8","RA-1","RA-2","RA-3","RA-5","SA-1","SA-10","SA-11","SA-2","SA-3","SA-4","SA-5","SA-8","SA-9","SC-1","SC-10","SC-12","SC-13","SC-15","SC-17","SC-18","SC-19","SC-2","SC-20","SC-21","SC-22","SC-23","SC-28","SC-39","SC-4","SC-5","SC-7","SC-8","SI-1","SI-10","SI-11","SI-12","SI-16","SI-2","SI-3","SI-4","SI-5","SI-7","SI-8", "ac-2.1","ac-2.2","ac-2.3","ac-2.4","ac-6.1","ac-6.2","ac-6.5","ac-6.9","ac-6.10","ac-11.1","ac-17.1","ac-17.2","ac-17.3","ac-17.4","ac-18.1","ac-19.5","ac-20.1","ac-20.2","at-2.2","au-2.3","au-3.1","au-6.1","au-6.3","au-7.1","au-8.1","au-9.4","ca-2.1","ca-3.5","ca-7.1","cm-2.1","cm-2.3","cm-2.7","cm-3.2","cm-7.1","cm-7.2","cm-7.4","cm-8.1","cm-8.3","cm-8.5","cp-2.1","cp-2.3","cp-2.8","cp-4.1","cp-6.1","cp-6.3","cp-7.1","cp-7.2","cp-7.3","cp-8.1","cp-8.2","cp-9.1","cp-10.2","ia-2.1","ia-2.2","ia-2.3","ia-2.8","ia-2.11","ia-2.12","ia-5.1","ia-5.2","ia-5.3","ia-5.11","ia-8.1","ia-8.2","ia-8.3","ia-8.4","ir-3.2","ir-4.1","ir-6.1","ir-7.1","ma-3.1","ma-3.2","ma-4.2","mp-5.4","mp-7.1","pe-6.1","pe-13.3","pl-2.3","pl-4.1","ra-5.1","ra-5.2","ra-5.5","sa-4.1","sa-4.2","sa-4.9","sa-4.10","sa-9.2","sc-7.3","sc-7.4","sc-7.5","sc-7.7","sc-8.1","si-2.2","si-3.1","si-3.2","si-4.2","si-4.4","si-4.5","si-7.1","si-7.7","si-8.1","si-8.2"}
	res = GetProfileBaselineDiff(controls , "moderate")
	assert.Equal(t, res, ans) 

	controls = []string{"AC-1", "AC-12"} 
	ans = []string{"AC-10","AC-11","AC-13","AC-14","AC-15","AC-16","AC-17","AC-18","AC-19","AC-2","AC-20","AC-21","AC-22","AC-23","AC-24","AC-25","AC-26","AC-27","AC-28","AC-29","AC-3","AC-30","AC-31","AC-32","AC-33","AC-34","AC-35","AC-36","AC-37","AC-38","AC-39","AC-4","AC-40","AC-41","AC-42","AC-43","AC-44","AC-45","AC-46","AC-47","AC-48","AC-49","AC-5","AC-50","AC-51","AC-52","AC-53","AC-54","AC-55","AC-56","AC-57","AC-58","AC-59","AC-6","AC-60","AC-61","AC-62","AC-63","AC-64","AC-65","AC-66","AC-67","AC-68","AC-69","AC-7","AC-8","AC-9","IA-8","IR-1","IR-2","IR-3","IR-4","IR-5","IR-6","IR-7","IR-8","MA-1","MA-2","MA-3","MA-4","MA-5","MA-6","MP-1","MP-2","MP-3","MP-4","MP-5","MP-6","MP-7","PE-1","PE-10","PE-11","PE-12","PE-13","PE-14","PE-15","PE-16","PE-17","PE-18","PE-2","PE-3","PE-4","PE-5","PE-6","PE-8","PE-9","PL-1","PL-2","PL-4","PL-8","PS-1","PS-2","PS-3","PS-4","PS-5","PS-6","PS-7","PS-8","RA-1","RA-2","RA-3","RA-5","SA-1","SA-10","SA-11","SA-12","SA-15","SA-16","SA-17","SA-2","SA-3","SA-4","SA-5","SA-8","SA-9","SC-1","SC-10","SC-12","SC-13","SC-15","SC-17","SC-18","SC-19","SC-2","SC-20","SC-21","SC-22","SC-23","SC-24","SC-28","SC-3","SC-39","SC-4","SC-5","SC-7","SC-8","SI-1","SI-10","SI-11","SI-12","SI-16","SI-2","SI-3","SI-4","SI-5","SI-6","SI-7","SI-8", "ac-2.1","ac-2.2","ac-2.3","ac-2.4","ac-2.5","ac-2.11","ac-2.12","ac-2.13","ac-6.1","ac-6.2","ac-6.3","ac-6.5","ac-6.9","ac-6.10","ac-11.1","ac-17.1","ac-17.2","ac-17.3","ac-17.4","ac-18.1","ac-18.4","ac-18.5","ac-19.5","ac-20.1","ac-20.2","at-2.2","au-2.3","au-3.1","au-3.2","au-5.1","au-5.2","au-6.1","au-6.3","au-6.5","au-6.6","au-7.1","au-8.1","au-9.2","au-9.3","au-9.4","au-12.1","au-12.3","ca-2.1","ca-2.2","ca-3.5","ca-7.1","cm-2.1","cm-2.2","cm-2.3","cm-2.7","cm-3.1","cm-3.2","cm-4.1","cm-5.1","cm-5.2","cm-5.3","cm-6.1","cm-6.2","cm-7.1","cm-7.2","cm-7.5","cm-8.1","cm-8.2","cm-8.3","cm-8.4","cm-8.5","cp-2.1","cp-2.2","cp-2.3","cp-2.4","cp-2.5","cp-2.8","cp-3.1","cp-4.1","cp-4.2","cp-6.1","cp-6.2","cp-6.3","cp-7.1","cp-7.2","cp-7.3","cp-7.4","cp-8.1","cp-8.2","cp-8.3","cp-8.4","cp-9.1","cp-9.2","cp-9.3","cp-9.5","cp-10.2","cp-10.4","ia-2.1","ia-2.2","ia-2.3","ia-2.4","ia-2.8","ia-2.9","ia-2.11","ia-2.12","ia-5.1","ia-5.2","ia-5.3","ia-5.11","ia-8.1","ia-8.2","ia-8.3","ia-8.4","ir-2.1","ir-2.2","ir-3.2","ir-4.1","ir-4.4","ir-5.1","ir-6.1","ir-7.1","ma-2.2","ma-3.1","ma-3.2","ma-3.3","ma-4.2","ma-4.3","ma-5.1","mp-5.4","mp-6.1","mp-6.2","mp-6.3","mp-7.1","pe-3.1","pe-6.1","pe-6.4","pe-8.1","pe-11.1","pe-13.1","pe-13.2","pe-13.3","pe-15.1","pl-2.3","pl-4.1","ps-4.2","ra-5.1","ra-5.2","ra-5.4","ra-5.5","sa-4.1","sa-4.2","sa-4.9","sa-4.10","sa-9.2","sc-7.3","sc-7.4","sc-7.5","sc-7.7","sc-7.8","sc-7.18","sc-7.21","sc-8.1","sc-12.1","si-2.1","si-2.2","si-3.1","si-3.2","si-4.2","si-4.4","si-4.5","si-5.1","si-7.1","si-7.2","si-7.5","si-7.7","si-7.14","si-8.1","si-8.2"}
	res = GetProfileBaselineDiff(controls , "high")
	assert.Equal(t, res, ans) 
}

func TestGetDiff(t *testing.T) {
	assert := assert.New(t)
    // set title version in metadata
    profile := &sdk_profile.Profile{}

    title := "Infobeyond BASELINE"
    version := "1.0"
    oscalVersion := "3.0"

    // request := request_models.SetTitleVersionRequest{Title: title, Version: version, OscalVersion: oscalVersion}
    SetTitleVersion(profile, title, version, oscalVersion)

    control := []string{"AC-1", "AC-12"}

    AddControls(profile, control, "../../nist.gov/SP800-53/rev4/xml/NIST_SP-800-53_rev4_MODERATE-baseline_profile.xml")
	profile.Id = "a"

	id := "catalog"
    desc := "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations"
    link := "NIST_SP-800-53_rev4_catalog.xml"
    media := "application/oscal.catalog+xml"

    AddBackMatter(profile, id, desc, link, media)
	WriteToFile(profile)

    res := GetDiff(profile.Id, "low")
	ans := []string{"AC-13","AC-14","AC-15","AC-16","AC-18","AC-19","AC-2","AC-20","AC-21","AC-22","AC-23","AC-24","AC-25","AC-26","AC-27","AC-28","AC-3","AC-30","AC-31","AC-33","AC-34","AC-35","AC-36","AC-37","AC-38","AC-39","AC-40","AC-42","AC-43","AC-44","AC-46","AC-48","AC-49","AC-50","AC-52","AC-53","AC-54","AC-55","AC-56","AC-57","AC-61","AC-62","AC-63","AC-64","AC-66","AC-67","AC-68","AC-69","AC-7","AC-8","IA-8","IR-1","IR-2","IR-4","IR-5","IR-6","IR-7","IR-8","MA-1","MA-2","MA-4","MA-5","MP-1","MP-2","MP-6","MP-7","PE-1","PE-12","PE-13","PE-14","PE-15","PE-16","PE-2","PE-3","PE-6","PE-8","PL-1","PL-2","PL-4","PS-1","PS-2","PS-3","PS-4","PS-5","PS-6","PS-7","PS-8","RA-1","RA-2","RA-3","RA-5","SA-1","SA-2","SA-3","SA-4","SA-5","SA-9","SC-1","SC-12","SC-13","SC-15","SC-20","SC-21","SC-22","SC-39","SC-5","SC-7","SI-1","SI-12","SI-2","SI-3","SI-4","SI-5", "ia-2.1", "ia-2.12","ia-5.1","ia-5.11","ia-8.1","ia-8.2","ia-8.3","ia-8.4","sa-4.10"}

	assert.Equal(len(res), len(ans)) 
	
	var notEqual = false
	for i:=0; i < len(res); i++ {
		if res[i] != ans[i] { 
			notEqual = true
		}
	} 
	
	assert.Equal(notEqual, false) 

}