package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// getFileList recursively gets all markdown files from the root directory.
// It ignores files in the Unsorted folder.
func getFileList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "Unsorted/") || strings.Contains(path, "node_modules/") {
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

// updateLinks reads a file and replaces bare words that match file names with Markdown links to those files.
// It will not link to the current file.
func updateLinks(filePath string, fileNames map[string]string, sortedFileNames []string, logFile *os.File) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		lines       []string
		updates     int
		scanner     = bufio.NewScanner(file)
		markdownURL = regexp.MustCompile(`\[[^\]]*\]\([^)]*\)`) // Matches any Markdown URL
		auxArray    []string
	)

	for scanner.Scan() {
		line := scanner.Text()
		updatedLine := line
		matches := markdownURL.FindAllStringIndex(line, -1)
		// Skip the line if the first character is "!"
		if len(line) > 0 && line[0] == '!' {
			lines = append(lines, updatedLine)
			continue
		}
		// Iterate over sorted filenames
		for _, fileNameBase := range sortedFileNames {
			fileName := fileNames[fileNameBase]
			if fileName == filePath {
				continue // Skip if the filename is the same as the current file's path
			}
			pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(fileNameBase))
			wordRegex := regexp.MustCompile(pattern)
			idx := wordRegex.FindStringIndex(updatedLine)

			// Replace only if the word is not part of a Markdown URL
			if idx != nil && !isInRanges(matches, idx) {
				link := sanitizePath(fileName)
				auxArray = append(auxArray, fmt.Sprintf("[%s](%s)", fileNameBase, link))
				updatedLine = wordRegex.ReplaceAllString(updatedLine, fmt.Sprintf("&W%d", updates))
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
		if err != nil {
			return updates, err
		}

		// Replace the placeholders with the respective links
		for i, link := range auxArray {
			placeholder := fmt.Sprintf("&W%d", i)
			for j, line := range lines {
				lines[j] = strings.ReplaceAll(line, placeholder, link)
			}
		}

		// Check for the final newline and preserve it if it exists
		if len(lines) > 0 && lines[len(lines)-1] != "" {
			err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0666)
		} else {
			err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0666)
		}
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
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	root := "./" // You can set this to the root directory of your Obsidian vault

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

	// We create a sorted list of filenames by descending length to ensure longer names are replaced first
	var sortedFileNames []string
	for name := range fileNames {
		sortedFileNames = append(sortedFileNames, name)
	}
	sort.Slice(sortedFileNames, func(i, j int) bool {
		return len(sortedFileNames[i]) > len(sortedFileNames[j])
	})

	// Update links in all files
	for _, file := range files {
		updates, err := updateLinks(file, fileNames, sortedFileNames, logFile)
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
	if strings.Contains(filePath, "README.md") {
		return 0, nil // No removing links from README.md
	}
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		lines   []string
		updates int
		scanner = bufio.NewScanner(file)
		regex   = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	)

	for scanner.Scan() {
		line := scanner.Text()

		// We also avoid removing external links and images
		if strings.Contains(line, "](http") || (line[0] == '!') {
			lines = append(lines, line)
			continue
		}

		updatedLine := regex.ReplaceAllString(line, "$1") // Remove the link, but keep the word
		if line != updatedLine {
			updates++
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

	// Check for the final newline and preserve it if it exists
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0666)
	} else {
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
