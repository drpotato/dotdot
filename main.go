package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/fsnotify.v1"

	"github.com/drpotato/dotdot/dotignore"
	"github.com/drpotato/dotdot/filesystem"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if dotignore.ShouldIgnore(event.Name) {
					continue
				}
				switch event.Op {
				case fsnotify.Create:
					log.Println("dotfile added:", event.Name)
					err = filesystem.LinkDotFile(event.Name)
				case fsnotify.Remove:
					log.Println("dotfile removed:", event.Name)
					err = filesystem.UnLinkDotFile(event.Name)
				case fsnotify.Chmod:
					log.Println("irrelevant operation:", event)
				case fsnotify.Rename:
					log.Println("irrelevant operation:", event)
				case fsnotify.Write:
					log.Println("irrelevant operation:", event)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	dotDir := filesystem.GetDotDirURI()

	files, err := ioutil.ReadDir(dotDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if dotignore.ShouldIgnore(file.Name()) {
			continue
		}
		err = filesystem.LinkDotFile(file.Name())
		if err != nil {
			log.Print(err)
		}
	}

	err = watcher.Add(dotDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("watching", dotDir)

	<-done
}
