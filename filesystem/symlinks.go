package filesystem

import (
	"errors"
	// "github.com/drpotato/dotdot/notification"
	"os"
)

func LinkDotFile(uri string) error {

	symLinkUri := GetSymLinkURI(uri)

	err := os.Symlink(uri, symLinkUri)
	if err != nil {
		// notification.NotifySymLinkError(symLinkUri)
	}
	return err
}

func UnLinkDotFile(uri string) error {

	symLinkUri := GetSymLinkURI(uri)

	targetUri, err := os.Readlink(symLinkUri)
	if err != nil || targetUri != uri {
		// notification.NotifyUnLinkError(symLinkUri)
		return errors.New("failed to remove symbolic link")
	}

	return os.Remove(symLinkUri)
}
