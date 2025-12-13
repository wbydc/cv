package internal

import (
	"os"
	"strings"
)

var (
	DataFilePath   = "data.json"
	PhotoFilePath  = "photo.jpg"
	OutputFilePath = "cv.html"
	ServerPort     = ":80"
)

func init() {
	// Override defaults if Environment Variables are set
	if val := os.Getenv("DATA_FILE_PATH"); val != "" {
		DataFilePath = val
	}
	if val := os.Getenv("PHOTO_FILE_PATH"); val != "" {
		PhotoFilePath = val
	}
	if val := os.Getenv("OUTPUT_FILE_PATH"); val != "" {
		OutputFilePath = val
	}
	if val := os.Getenv("SERVER_PORT"); val != "" {
		// Ensure port starts with ":"
		if !strings.HasPrefix(val, ":") {
			ServerPort = ":" + val
		} else {
			ServerPort = val
		}
	}
}
