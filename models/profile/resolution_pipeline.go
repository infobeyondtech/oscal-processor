package profile

import (
	"bytes"
	"errors"
	"os/exec"
)

/* ResolveProfile applies a series of transformations to the input profile,
with the given transformations, jar package, writing the given output
file. */
func ResolveProfile(saxonJarPath string, xslPath string,
	inputFile string, outputFile string) error {

	javaCmd := exec.Command("java",
		"-jar", saxonJarPath,
		"-s:"+inputFile,
		"-o:"+outputFile,
		"-xsl:"+xslPath)

	javaaCmdOutput := &bytes.Buffer{}
	javaCmdErr := &bytes.Buffer{}
	javaCmd.Stdout = javaaCmdOutput
	javaCmd.Stderr = javaCmdErr

	if err := javaCmd.Run(); err != nil {
		stderr := javaCmdErr.String()
		if stderr == "" {
			return err
		}
		return errors.New(javaCmdErr.String())
	}
	return nil
}
