package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// getImageAsBase64 reads an image file and returns it as a base64-encoded data URL
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

	var mimeType string
	switch ext {
	case "png":
		mimeType = "image/png"
	case "gif":
		mimeType = "image/gif"
	default:
		mimeType = "image/jpeg"
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, encoded)
}

// parseDate attempts to parse a date string using multiple common formats
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

// GenerateCV reads the data file and photo, processes them, and generates the CV HTML file
func GenerateCV() {
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
