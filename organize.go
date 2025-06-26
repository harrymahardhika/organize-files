package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var extensionMap = map[string][]string{
	"images":    {"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp"},
	"documents": {"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "txt", "md"},
	"videos":    {"mp4", "avi", "mov", "mkv", "flv"},
	"audio":     {"mp3", "wav", "aac", "flac"},
	"archives":  {"zip", "tar", "gz", "rar", "7z"},
}

func toSnakeCase(name string) string {
	// Remove extension first
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)

	// Replace non-alphanumeric with underscores
	reg := regexp.MustCompile(`[^\w]+`)
	snake := reg.ReplaceAllString(base, "_")

	// Remove leading/trailing underscores
	snake = strings.Trim(snake, "_")

	return strings.ToLower(snake) + strings.ToLower(ext)
}

func toSnakeCaseDir(name string) string {
	// Replace non-alphanumeric with underscores
	reg := regexp.MustCompile(`[^\w]+`)
	snake := reg.ReplaceAllString(name, "_")

	// Remove leading/trailing underscores
	snake = strings.Trim(snake, "_")

	return strings.ToLower(snake)
}

func getFolderForExt(ext string) string {
	ext = strings.TrimPrefix(strings.ToLower(ext), ".")
	for folder, exts := range extensionMap {
		for _, e := range exts {
			if e == ext {
				return folder
			}
		}
	}
	return "others"
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// First pass: rename directories to snake_case
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		oldName := entry.Name()
		newName := toSnakeCaseDir(oldName)
		
		// Skip if already in snake_case
		if oldName == newName {
			continue
		}

		oldPath := filepath.Join(dir, oldName)
		newPath := filepath.Join(dir, newName)

		// Handle conflicts by appending _1, _2, etc.
		counter := 1
		baseName := newName
		for fileExists(newPath) {
			newName = fmt.Sprintf("%s_%d", baseName, counter)
			newPath = filepath.Join(dir, newName)
			counter++
		}

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Error renaming directory:", err)
			continue
		}

		fmt.Printf("Renamed directory '%s' -> '%s'\n", oldName, newName)
	}

	// Reload directory entries after renaming directories
	entries, err = os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Second pass: organize files
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		oldName := entry.Name()
		oldPath := filepath.Join(dir, oldName)

		newName := toSnakeCase(oldName)
		ext := filepath.Ext(newName)
		folder := getFolderForExt(ext)
		destDir := filepath.Join(dir, folder)

		// Create folder if not exists
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			if err := os.Mkdir(destDir, 0755); err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}
		}

		destPath := filepath.Join(destDir, newName)
		// Avoid overwrite by adding _1, _2, etc.
		counter := 1
		base := strings.TrimSuffix(newName, ext)
		for fileExists(destPath) {
			destPath = filepath.Join(destDir, fmt.Sprintf("%s_%d%s", base, counter, ext))
			counter++
		}

		err := os.Rename(oldPath, destPath)
		if err != nil {
			fmt.Println("Error moving file:", err)
			continue
		}

		relDest, _ := filepath.Rel(dir, destPath)
		fmt.Printf("Moved '%s' -> '%s'\n", oldName, relDest)
	}
}

