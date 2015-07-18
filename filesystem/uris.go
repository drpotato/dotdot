package filesystem

import (
	"log"
	"os/user"
	"path/filepath"
	"sync"
)

func GetSymLinkURI(uri string) string {

	userDir := GetUserDirURI()

	_, fileName := filepath.Split(uri)
	symLinkUri := filepath.Join(userDir, fileName)

	return symLinkUri
}

var userDirOnce sync.Once
var userDirUri string

func GetUserDirURI() string {

	userDirOnce.Do(func() {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		userDirUri = currentUser.HomeDir
	})

	return userDirUri
}

var dotDirOnce sync.Once
var dotDirUri string

func GetDotDirURI() string {

	userDir := GetUserDirURI()

	dotDirOnce.Do(func() {
		dotDirUri = filepath.Join(userDir, ".dot")
	})

	return dotDirUri
}
