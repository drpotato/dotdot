package util

import (
	"github.com/deckarep/gosx-notifier"
)

func NotifySymLinkError(newFileName string) error {

	// oldFile := linkError.Old
	note := gosxnotifier.NewNotification(newFileName + " already exists")
	note.Title = "Failed to create symbolic link"

	return note.Push()
}
