package util

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

func LinkDotFile(uri string) error {

	symLinkUri, err := GetSymLinkUri(uri)
	if err != nil {
		return err
	}

	err = os.Symlink(uri, symLinkUri)
	if err != nil {
		NotifySymLinkError(symLinkUri)
		return errors.New("failed to create symbolic link")
	}
	return nil
}

func UnLinkDotFile(uri string) error {

	symLinkUri, err := GetSymLinkUri(uri)
	if err != nil {
		return err
	}

	targetUri, err := os.Readlink(symLinkUri)
	if err != nil || targetUri != uri {
		NotifyUnLinkError(symLinkUri)
		return errors.New("failed to remove symbolic link")
	}

	return os.Remove(symLinkUri)
}

func GetSymLinkUri(uri string) (string, error) {
	userDir, err := GetUserDir()

	_, fileName := filepath.Split(uri)
	symLinkUri := filepath.Join(userDir, fileName)

	return symLinkUri, err
}

func GetUserDir() (string, error) {

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.HomeDir, nil
}

func GetDotDir() (string, error) {

	userDir, err := GetUserDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userDir, ".dot"), nil
}

func IsSymLink(uri string) bool {
	file, err := os.Lstat(uri)
	if err != nil {
		return false
	}
	return file.Mode()&os.ModeSymlink == os.ModeSymlink
}
