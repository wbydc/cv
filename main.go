package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	DataFilename   = "data.json"
	PhotoFilename  = "photo.jpg"
	OutputFilename = "cv.html"
	ServerPort     = ":80"
)

// --- Helper Functions ---

func getImageAsBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		// Don't crash the server if photo is missing, just log it
		log.Printf("Warning: Could not read photo at %s: %v", path, err)
		return ""
	}

	ext := strings.ToLower(filepath.Ext(path))
	if len(ext) > 0 {
		ext = ext[1:]
	}

	mimeType := "image/jpeg"
	if ext == "png" {
		mimeType = "image/png"
	} else if ext == "gif" {
		mimeType = "image/gif"
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, encoded)
}

func parseDate(dateStr string) time.Time {
	formats := []string{
		"2006-01-02",
		"Jan 2006",
		"January 2006",
		"2006",
	}
	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t
		}
	}
	return time.Time{}
}

// --- Core Logic ---

func generateCV() {
	log.Println("🔄 Regenerating CV...")

	// 1. Read Data
	jsonData, err := os.ReadFile(DataFilename)
	if err != nil {
		log.Printf("Error reading %s: %v", DataFilename, err)
		return
	}

	var cvData CVData
	if err := json.Unmarshal(jsonData, &cvData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	// 2. Read Photo
	photoBase64 := getImageAsBase64(PhotoFilename)

	// 3. Process Data
	emailEncoded := base64.StdEncoding.EncodeToString([]byte(cvData.Email))
	phoneEncoded := base64.StdEncoding.EncodeToString([]byte(cvData.Phone))

	sort.Slice(cvData.Experience, func(i, j int) bool {
		a := cvData.Experience[i]
		b := cvData.Experience[j]
		isCurrentA := a.EndDate == "" || strings.ToLower(strings.TrimSpace(a.EndDate)) == "present"
		isCurrentB := b.EndDate == "" || strings.ToLower(strings.TrimSpace(b.EndDate)) == "present"

		if isCurrentA && !isCurrentB {
			return true
		}
		if !isCurrentA && isCurrentB {
			return false
		}
		return parseDate(a.StartDate).After(parseDate(b.StartDate))
	})

	ctx := TemplateContext{
		CVData:       cvData,
		PhotoBase64:  template.URL(photoBase64),
		EmailEncoded: emailEncoded,
		PhoneEncoded: phoneEncoded,
	}

	// 4. Generate HTML
	tmpl, err := template.New("cv").Parse(HTMLTemplate)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return
	}

	outputFile, err := os.Create(OutputFilename)
	if err != nil {
		log.Printf("Error creating output file: %v", err)
		return
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, ctx); err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}

	log.Println("✅ CV Generated successfully.")
}

// --- Main Execution ---

func main() {
	// 1. Initial Generation
	generateCV()

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
						generateCV()
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
