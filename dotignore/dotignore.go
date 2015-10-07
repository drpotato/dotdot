package dotignore

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/drpotato/dotdot/filesystem"
)

var ignoredFiles map[string]bool
var once sync.Once

func ShouldIgnore(uri string) bool {

	GetIgnoredFiles()

	dotDirURI := filesystem.GetDotDirURI()

	_, fileName := filepath.Split(uri)

	return strings.HasPrefix(uri, dotDirURI) && ignoredFiles[fileName]
}

func GetIgnoredFiles() map[string]bool {
	once.Do(func() {
		ignoredFiles = map[string]bool{
			".git":       true,
			".gitignore": true,
		}
		loadDotIgnore()
	})
	return ignoredFiles
}

func loadDotIgnore() {

	dotDirURI := filesystem.GetDotDirURI()
	dotIgnoreUri := filepath.Join(dotDirURI, ".dotignore")

	dotIgnoreFile, err := os.Open(dotIgnoreUri)
	if err != nil {
		return
	}
	defer dotIgnoreFile.Close()

	scanner := bufio.NewScanner(dotIgnoreFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		ignoredFiles[line] = true
	}
}
