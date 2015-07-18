package filesystem

import (
	"path/filepath"
	"os/user"
)

func GetSymLinkURI(uri string) (string, error) {
	userDir, err := GetUserDirURI()

	_, fileName := filepath.Split(uri)
	symLinkUri := filepath.Join(userDir, fileName)

	return symLinkUri, err
}

func GetUserDirURI() (string, error) {

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.HomeDir, nil
}

func GetDotDirURI() (string, error) {

	userDir, err := GetUserDirURI()
	if err != nil {
		return "", err
	}

	return filepath.Join(userDir, ".dot"), nil
}
