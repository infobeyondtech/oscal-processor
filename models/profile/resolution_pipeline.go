// Provides the features of applying the standard XSLT transformations
// defined by OSCAL.
package profile

import (
	"bytes"
	"errors"
	"os/exec"
)

func applyResolutionProcedure(saxonJarPath string, xslPath string,
	inputFile string, outputFile string) error {

	javaCmd := exec.Command("java",
		"-jar", saxonJarPath,
		"-s:"+inputFile,
		"-o:"+outputFile,
		"-xsl:"+xslPath)

	javaaCmdOutput := &bytes.Buffer{}
	javaCmdErr := &bytes.Buffer{}
	javaCmd.Stdout = xmllintCmdOutput
	javaCmd.Stderr = xmllintCmdErr

	if err := javaCmd.Run(); err != nil {
		stderr := javaCmdErr.String()
		if stderr == "" {
			return err
		}
		return errors.New(javaCmdErr.String())
	}
	return nil
}
