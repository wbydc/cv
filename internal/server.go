package internal

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
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Check absolute paths to ensure we match correctly
					absEventName, _ := filepath.Abs(event.Name)
					absDataPath, _ := filepath.Abs(DataFilePath)
					absPhotoPath, _ := filepath.Abs(PhotoFilePath)

					if absEventName == absDataPath || absEventName == absPhotoPath {
						log.Printf("Detected change in %s", event.Name)
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

	// Watch the directory containing the data file
	dataDir := filepath.Dir(DataFilePath)
	log.Printf("👀 Watching directory: %s", dataDir)
	err = watcher.Add(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Setup HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the generated file
		http.ServeFile(w, r, OutputFilePath)
	})

	fmt.Printf("\n🚀 Server started!\n")
	fmt.Printf("👉 Go to http://localhost%s to view your CV\n", ServerPort)
	fmt.Printf("👀 Watching for changes in %s and %s...\n\n", DataFilePath, PhotoFilePath)

	log.Fatal(http.ListenAndServe(ServerPort, nil))
}
