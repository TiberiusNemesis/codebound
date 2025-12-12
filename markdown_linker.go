package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Replacement struct {
	Start   int
	End     int
	NewText string
	Match   string
}

type LinkMatch struct {
	TargetFile string
	Position   string
	LineNumber int
}

type Relationship struct {
	Source        string   `json:"source"`
	Target        string   `json:"target"`
	Weight        float64  `json:"weight"`
	Frequency     int      `json:"frequency"`
	Positions     []string `json:"positions"`
	Bidirectional bool     `json:"bidirectional"`
}

type RelationshipOutput struct {
	Generated     string         `json:"generated"`
	TotalFiles    int            `json:"total_files"`
	Relationships []Relationship `json:"relationships"`
}

var positionWeights = map[string]float64{
	"title":  3.0,
	"header": 2.0,
	"body":   1.0,
}

func getFileList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "Unsorted/") || strings.Contains(path, "node_modules/") {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func sanitizePath(path string) string {
	return strings.ReplaceAll(path, " ", "%20")
}

func detectPosition(line string, lineNum int) string {
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "# ") && lineNum <= 3 {
		return "title"
	}
	if strings.HasPrefix(trimmed, "##") {
		return "header"
	}
	return "body"
}

func isInRanges(ranges [][]int, start, end int) bool {
	for _, r := range ranges {
		if start >= r[0] && end <= r[1] {
			return true
		}
	}
	return false
}

func calculateDocumentFrequency(files []string, fileRegexes map[string]*regexp.Regexp) map[string]int {
	docFreq := make(map[string]int)

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			continue
		}

		content, err := os.ReadFile(filePath)
		file.Close()
		if err != nil {
			continue
		}

		contentStr := string(content)
		seen := make(map[string]bool)

		for term, regex := range fileRegexes {
			if !seen[term] && regex.MatchString(contentStr) {
				docFreq[term]++
				seen[term] = true
			}
		}
	}

	return docFreq
}

func updateLinks(
	filePath string,
	fileNames map[string]string,
	sortedFileNames []string,
	fileRegexes map[string]*regexp.Regexp,
	markdownURLRegex *regexp.Regexp,
) (int, []LinkMatch, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	var (
		lines       []string
		matches     []LinkMatch
		updates     int
		scanner     = bufio.NewScanner(file)
		lineNum     = 0
		inCodeBlock = false
	)

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCodeBlock = !inCodeBlock
			lines = append(lines, line)
			continue
		}

		if inCodeBlock {
			lines = append(lines, line)
			continue
		}

		if len(line) > 0 && line[0] == '!' {
			lines = append(lines, line)
			continue
		}

		position := detectPosition(line, lineNum)

		if position == "title" {
			lines = append(lines, line)
			continue
		}

		existingLinks := markdownURLRegex.FindAllStringIndex(line, -1)

		var replacements []Replacement

		for _, fileNameBase := range sortedFileNames {
			fileName := fileNames[fileNameBase]
			if fileName == filePath {
				continue
			}

			regex := fileRegexes[fileNameBase]
			allMatches := regex.FindAllStringIndex(line, -1)

			for _, idx := range allMatches {
				if isInRanges(existingLinks, idx[0], idx[1]) {
					continue
				}

				alreadyReplaced := false
				for _, r := range replacements {
					if (idx[0] >= r.Start && idx[0] < r.End) || (idx[1] > r.Start && idx[1] <= r.End) {
						alreadyReplaced = true
						break
					}
				}
				if alreadyReplaced {
					continue
				}

				originalText := line[idx[0]:idx[1]]
				link := sanitizePath(fileName)
				newText := fmt.Sprintf("[%s](%s)", originalText, link)

				replacements = append(replacements, Replacement{
					Start:   idx[0],
					End:     idx[1],
					NewText: newText,
					Match:   fileNameBase,
				})

				matches = append(matches, LinkMatch{
					TargetFile: fileNameBase,
					Position:   position,
					LineNumber: lineNum,
				})
				updates++
			}
		}

		sort.Slice(replacements, func(i, j int) bool {
			return replacements[i].Start > replacements[j].Start
		})

		updatedLine := line
		for _, r := range replacements {
			updatedLine = updatedLine[:r.Start] + r.NewText + updatedLine[r.End:]
		}

		lines = append(lines, updatedLine)
	}

	if err := scanner.Err(); err != nil {
		return updates, matches, err
	}

	if updates > 0 {
		content := strings.Join(lines, "\n")
		if len(lines) > 0 && lines[len(lines)-1] != "" {
			content += "\n"
		}
		err = os.WriteFile(filePath, []byte(content), 0666)
	}

	return updates, matches, err
}

func buildRelationships(
	allMatches map[string][]LinkMatch,
	docFreq map[string]int,
	totalFiles int,
) []Relationship {
	type relKey struct {
		source, target string
	}

	relMap := make(map[relKey]*Relationship)

	for source, matches := range allMatches {
		for _, m := range matches {
			key := relKey{source: source, target: m.TargetFile}
			if rel, exists := relMap[key]; exists {
				rel.Frequency++
				rel.Positions = append(rel.Positions, m.Position)
			} else {
				relMap[key] = &Relationship{
					Source:    source,
					Target:    m.TargetFile,
					Frequency: 1,
					Positions: []string{m.Position},
				}
			}
		}
	}

	for key, rel := range relMap {
		reverseKey := relKey{source: key.target, target: key.source}
		if _, exists := relMap[reverseKey]; exists {
			rel.Bidirectional = true
		}

		var positionSum float64
		for _, pos := range rel.Positions {
			positionSum += positionWeights[pos]
		}
		avgPositionWeight := positionSum / float64(len(rel.Positions))

		idf := 1.0
		if df, exists := docFreq[rel.Target]; exists && df > 0 {
			idf = math.Log(float64(totalFiles) / float64(df))
			if idf < 0.1 {
				idf = 0.1
			}
		}

		weight := float64(rel.Frequency) * avgPositionWeight * idf
		if rel.Bidirectional {
			weight *= 1.5
		}

		rel.Weight = math.Round(weight*100) / 100
	}

	var relationships []Relationship
	for _, rel := range relMap {
		relationships = append(relationships, *rel)
	}

	sort.Slice(relationships, func(i, j int) bool {
		return relationships[i].Weight > relationships[j].Weight
	})

	return relationships
}

func writeRelationshipsJSON(relationships []Relationship, totalFiles int) error {
	output := RelationshipOutput{
		Generated:     time.Now().Format(time.RFC3339),
		TotalFiles:    totalFiles,
		Relationships: relationships,
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("relationships.json", data, 0644)
}

func main() {
	start := time.Now()
	logFileName := "updates.log"
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	root := "./"

	files, err := getFileList(root)
	if err != nil {
		panic(err)
	}

	fileNames := make(map[string]string)
	for _, file := range files {
		fileNames[strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))] = file
	}

	var sortedFileNames []string
	for name := range fileNames {
		sortedFileNames = append(sortedFileNames, name)
	}
	sort.Slice(sortedFileNames, func(i, j int) bool {
		return len(sortedFileNames[i]) > len(sortedFileNames[j])
	})

	fileRegexes := make(map[string]*regexp.Regexp)
	for _, name := range sortedFileNames {
		pattern := fmt.Sprintf(`(?i)\b%s\b`, regexp.QuoteMeta(name))
		fileRegexes[name] = regexp.MustCompile(pattern)
	}

	markdownURLRegex := regexp.MustCompile(`\[[^\]]*\]\([^)]*\)`)

	fmt.Fprintf(logFile, "=== Link Update Run: %s ===\n", start.Format(time.RFC3339))
	fmt.Fprintf(logFile, "Total files to process: %d\n\n", len(files))

	docFreq := calculateDocumentFrequency(files, fileRegexes)

	allMatches := make(map[string][]LinkMatch)
	totalUpdates := 0

	for _, file := range files {
		baseName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		updates, matches, err := updateLinks(file, fileNames, sortedFileNames, fileRegexes, markdownURLRegex)
		if err != nil {
			fmt.Fprintf(logFile, "ERROR: %s: %s\n", file, err)
		} else if updates > 0 {
			fmt.Fprintf(logFile, "Updated %d links in %s\n", updates, file)
			allMatches[baseName] = matches
			totalUpdates += updates
		}
	}

	relationships := buildRelationships(allMatches, docFreq, len(files))

	if err := writeRelationshipsJSON(relationships, len(files)); err != nil {
		fmt.Fprintf(logFile, "\nERROR writing relationships.json: %s\n", err)
	} else {
		fmt.Fprintf(logFile, "\nGenerated relationships.json with %d relationships\n", len(relationships))
	}

	bidirectionalCount := 0
	for _, r := range relationships {
		if r.Bidirectional {
			bidirectionalCount++
		}
	}

	duration := time.Since(start)
	fmt.Fprintf(logFile, "\n=== Summary ===\n")
	fmt.Fprintf(logFile, "Files processed: %d\n", len(files))
	fmt.Fprintf(logFile, "Total links created: %d\n", totalUpdates)
	fmt.Fprintf(logFile, "Unique relationships: %d\n", len(relationships))
	fmt.Fprintf(logFile, "Bidirectional relationships: %d\n", bidirectionalCount/2)
	fmt.Fprintf(logFile, "Duration: %v\n", duration)

	fmt.Printf("Links updated successfully in %v.\n", duration)
	fmt.Printf("  - %d links created across %d files\n", totalUpdates, len(files))
	fmt.Printf("  - %d relationships found (%d bidirectional)\n", len(relationships), bidirectionalCount/2)
	fmt.Println("Check updates.log and relationships.json for details.")
}

func removeLinks(filePath string) (int, error) {
	if strings.Contains(filePath, "README.md") {
		return 0, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		lines       []string
		updates     int
		scanner     = bufio.NewScanner(file)
		regex       = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
		inCodeBlock = false
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCodeBlock = !inCodeBlock
			lines = append(lines, line)
			continue
		}

		if inCodeBlock {
			lines = append(lines, line)
			continue
		}

		if strings.Contains(line, "](http") || (len(line) > 0 && line[0] == '!') {
			lines = append(lines, line)
			continue
		}

		updatedLine := regex.ReplaceAllString(line, "$1")
		if line != updatedLine {
			updates++
		}
		lines = append(lines, updatedLine)
	}

	if err := scanner.Err(); err != nil {
		return updates, err
	}

	if updates > 0 {
		content := strings.Join(lines, "\n")
		if len(lines) > 0 && lines[len(lines)-1] != "" {
			content += "\n"
		}
		err = os.WriteFile(filePath, []byte(content), 0666)
	}

	return updates, err
}

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
