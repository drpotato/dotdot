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
		err = NotifySymLinkError(symLinkUri)
	}
	return err
}

func UnLinkDotFile(uri string) error {

	symLinkUri, err := GetSymLinkUri(uri)
	if err != nil {
		return err
	}

	targetUri, err := os.Readlink(symLinkUri)
	if err != nil {
		return err
	}
	if targetUri != uri {
		return errors.New("symlink mismatch")
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
