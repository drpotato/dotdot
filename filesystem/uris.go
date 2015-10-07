package filesystem

import (
	"log"
	"os/user"
	"path/filepath"
)

func GetSymLinkURI(uri string) string {

	userDir := GetUserDirURI()

	_, fileName := filepath.Split(uri)
	symLinkUri := filepath.Join(userDir, fileName)

	return symLinkUri
}

func GetUserDirURI() string {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return currentUser.HomeDir
}

func GetDotDirURI() string {

	userDir := GetUserDirURI()

	return filepath.Join(userDir, ".dot")
}
