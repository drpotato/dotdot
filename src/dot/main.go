package dot

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/fsnotify.v1"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	currentUser, err := user.Current()

	userDir := currentUser.HomeDir
	dotDir := userDir + "/.dot"

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					_, fileName := filepath.Split(event.Name)
					newFileName := userDir + "/" + fileName
					err = os.Symlink(event.Name, newFileName)
					if err != nil {
						log.Println(err)
						err = NotifySymLinkError(newFileName)
						if err != nil {
							log.Println(err)
						}

					}
					log.Println("File added:", fileName)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dotDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("watching", dotDir)

	<-done
}
