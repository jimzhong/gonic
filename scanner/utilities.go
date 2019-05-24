package scanner

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/dhowden/tag"
)

var trackExtensions = map[string]string{
	"mp3":  "audio/mpeg",
	"flac": "audio/x-flac",
	"aac":  "audio/x-aac",
	"m4a":  "audio/m4a",
	"ogg":  "audio/ogg",
}

func isTrack(fullPath string) (string, string, bool) {
	ext := filepath.Ext(fullPath)[1:]
	mine, ok := trackExtensions[ext]
	if !ok {
		return "", "", false
	}
	return mine, ext, true
}

var coverFilenames = map[string]bool{
	"cover.png":   true,
	"cover.jpg":   true,
	"cover.jpeg":  true,
	"folder.png":  true,
	"folder.jpg":  true,
	"folder.jpeg": true,
	"album.png":   true,
	"album.jpg":   true,
	"album.jpeg":  true,
	"front.png":   true,
	"front.jpg":   true,
	"front.jpeg":  true,
}

func isCover(fullPath string) bool {
	_, filename := path.Split(fullPath)
	_, ok := coverFilenames[strings.ToLower(filename)]
	return ok
}

func readTags(path string) (tag.Metadata, error) {
	trackData, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("when tags from disk: %v", err)
	}
	defer trackData.Close()
	tags, err := tag.ReadFrom(trackData)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func logElapsed(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("finished %s in %s\n", name, elapsed)
}
