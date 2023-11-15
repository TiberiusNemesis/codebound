package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// getFileList recursively gets all markdown files from the root directory.
func getFileList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "Unsorted/") {
			return nil // Ignore files in the Unsorted folder
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// sanitizePath makes a path URL-friendly, especially for spaces in folder names.
func sanitizePath(path string) string {
	return strings.ReplaceAll(path, " ", "%20")
}

// updateLinks reads a file and replaces bare words that match file names with Markdown links.
func updateLinks(filePath string, fileNames map[string]string, logFile *os.File) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		lines       []string
		updates     int
		scanner     = bufio.NewScanner(file)
		markdownURL = regexp.MustCompile(`\[[^\]]*\]\([^)]*\)`)   // Matches any Markdown URL
	)

	for scanner.Scan() {
		line := scanner.Text()
		updatedLine := line
		matches := markdownURL.FindAllStringIndex(line, -1)

		for _, fileName := range fileNames {
			if fileName == filePath {
				continue // Skip if the filename is the same as the current file's path
			}
			fileNameBase := strings.TrimSuffix(filepath.Base(fileName), ".md")
			pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(fileNameBase))
			wordRegex := regexp.MustCompile(pattern)
			idx := wordRegex.FindStringIndex(updatedLine)

			// Replace only if the word is not part of a Markdown URL
			if idx != nil && !isInRanges(matches, idx) {
				link := sanitizePath(fileName)
				updatedLine = wordRegex.ReplaceAllString(updatedLine, fmt.Sprintf("[%s](%s)", fileNameBase, link))
				updates++
			}
		}
		lines = append(lines, updatedLine)
	}

	if err := scanner.Err(); err != nil {
		return updates, err
	}

	// Write the updated content back to the file if there were updates
	if updates > 0 {
		err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0666)
	}
	return updates, err
}

// isInRanges checks if a given index is within any of the ranges provided.
func isInRanges(ranges [][]int, index []int) bool {
	for _, r := range ranges {
		if index[0] >= r[0] && index[1] <= r[1] {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	logFileName := "updates.log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	root := "./" // Set the root directory of your Obsidian vault

	// Get a list of all markdown files
	files, err := getFileList(root)
	if err != nil {
		panic(err)
	}
	// removeLinksFromFiles(files) // Uncomment this line to remove all links from all files


	// Create a map of file names without extension to their relative paths
	fileNames := make(map[string]string)
	for _, file := range files {
		fileNames[strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))] = file
	}

	// Update links in all files
	for _, file := range files {
		updates, err := updateLinks(file, fileNames, logFile)
		if err != nil {
			fmt.Fprintf(logFile, "Error updating links in file %s: %s\n", file, err)
		} else if updates > 0 {
			fmt.Fprintf(logFile, "Updated %d links in file %s\n", updates, file)
		}
	}

	duration := time.Since(start)
	fmt.Fprintf(logFile, "Links updated successfully in %v.\n", duration)

	fmt.Println("Links updated successfully. Check the log file for details.")
}

func removeLinks(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		lines   []string
		updates int
		scanner = bufio.NewScanner(file)
		regex   = regexp.MustCompile(`\[[^\]]+\]\([^)]+\)`) // Matches Markdown links
	)

	for scanner.Scan() {
		line := scanner.Text()
		updatedLine := regex.ReplaceAllStringFunc(line, func(match string) string {
			updates++
			return ""
		})
		lines = append(lines, updatedLine)
	}

	if err := scanner.Err(); err != nil {
		return updates, err
	}

	// Write the updated content back to the file if there were updates
	if updates > 0 {
		err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0666)
	}
	return updates, err
}

// removeLinksFromFiles calls removeLinks on each file in the provided list.
func removeLinksFromFiles(files []string) {
	for _, file := range files {
		updates, err := removeLinks(file)
		if err != nil {
			fmt.Printf("Error removing links from file %s: %s\n", file, err)
		} else if updates > 0 {
			fmt.Printf("Removed %d links from file %s\n", updates, file)
		}
	}
}