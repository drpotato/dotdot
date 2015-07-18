package main

import (
	"log"
	"path/filepath"

	"gopkg.in/fsnotify.v1"

	"github.com/drpotato/dotdot/util"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	dotDir, err := util.GetDotDir()
	if err != nil {
		log.Fatal(err)
	}

	ignoreFiles := map[string]bool{
		filepath.Join(dotDir, ".dot"):       true,
		filepath.Join(dotDir, ".gitignore"): true,
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if ignoreFiles[event.Name] {
					continue
				}
				switch event.Op {
				case fsnotify.Create:
					log.Println("dotfile added:", event.Name)
					err = util.LinkDotFile(event.Name)
				case fsnotify.Remove:
					log.Println("dotfile removed:", event.Name)
					err = util.UnLinkDotFile(event.Name)
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

	err = watcher.Add(dotDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("watching", dotDir)

	<-done
}
