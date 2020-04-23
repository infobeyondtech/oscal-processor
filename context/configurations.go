package Context

import (
	"os/user"
	"path/filepath"
	"strings"
)

const (
	AppData     = "~/oscal_processing_space"
	TempDir     = "~/oscal_processing_space/tmp"
	UploadDir   = "~/oscal_processing_space/uploads"
	DownloadDir = "~/oscal_processing_space/downloads"
	JarLibDir   = "~/.nanshiie_baker/jars"
)

var usr, _ = user.Current()
var dir = usr.HomeDir

func ExpandPath(path string) string {
	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		path = filepath.Join(dir, path[2:])
	}
	return path
}
