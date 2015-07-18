package filesystem
import (
	"errors"
	"os"
	"github.com/drpotato/dotdot/notification"
)

func LinkDotFile(uri string) error {

	symLinkUri, err := GetSymLinkURI(uri)
	if err != nil {
		return err
	}

	err = os.Symlink(uri, symLinkUri)
	if err != nil {
		notification.NotifySymLinkError(symLinkUri)
		return errors.New("failed to create symbolic link")
	}
	return nil
}

func UnLinkDotFile(uri string) error {

	symLinkUri, err := GetSymLinkURI(uri)
	if err != nil {
		return err
	}

	targetUri, err := os.Readlink(symLinkUri)
	if err != nil || targetUri != uri {
		notification.NotifyUnLinkError(symLinkUri)
		return errors.New("failed to remove symbolic link")
	}

	return os.Remove(symLinkUri)
}
