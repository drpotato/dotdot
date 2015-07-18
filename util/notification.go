package util

import (
	"path/filepath"

	"github.com/deckarep/gosx-notifier"
)

func NotifySymLinkError(symLinkUri string) error {

	_, fileName := filepath.Split(symLinkUri)

	note := gosxnotifier.NewNotification(fileName + " already exists")
	note.Title = "Failed to create symbolic link"

	return note.Push()
}

func NotifyUnLinkError(symLinkUri string) error {

	_, fileName := filepath.Split(symLinkUri)

	// oldFile := linkError.Old
	note := gosxnotifier.NewNotification(fileName + " doesn't link to .dot folder")
	note.Title = "Failed to remove symbolic link"

	return note.Push()
}
