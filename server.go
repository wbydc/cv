package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

// StartServer initializes the file watcher and starts the HTTP server
func StartServer() {
	// 1. Initial Generation
	GenerateCV()

	// 2. Setup File Watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Check if the event is a Write (save) operation
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Check if the modified file is one we care about
					// We use filepath.Base to handle paths correctly
					fileName := filepath.Base(event.Name)
					if fileName == DataFilename || fileName == PhotoFilename {
						log.Printf("Detected change in %s", fileName)
						// Small delay to ensure file write is complete (editors can be atomic)
						time.Sleep(100 * time.Millisecond)
						GenerateCV()
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	// Watch the current directory
	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Setup HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the generated file
		http.ServeFile(w, r, OutputFilename)
	})

	fmt.Printf("\n🚀 Server started!\n")
	fmt.Printf("👉 Go to http://localhost%s to view your CV\n", ServerPort)
	fmt.Printf("👀 Watching for changes in %s and %s...\n\n", DataFilename, PhotoFilename)

	log.Fatal(http.ListenAndServe(ServerPort, nil))
}
