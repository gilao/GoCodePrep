package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFile := "InterviewCollection.md"
	if err := addHyperlinks(inputFile); err != nil {
		fmt.Println("Error adding hyperlinks:", err)
		return
	}
	err := parseMarkdown(inputFile)
	if err != nil {
		fmt.Println("Error parsing markdown:", err)
	}
	fmt.Println("Markdown parsed successfully")
}
func parseMarkdown(inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		currentDir      string
		currentFile     *os.File
		currentFileName string
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") {
			title := strings.TrimPrefix(line, "# ")
			title = strings.TrimSpace(title)
			level := strings.Count(line, "#")
			if level == 1 {
				currentDir = filepath.Join(".", title)
				err := os.MkdirAll(currentDir, os.ModePerm)
				if err != nil {
					fmt.Println("Error creating directory:", err)
					return err
				}
				if currentFile != nil {
					currentFile.Close()
					currentFile = nil
				}
				fmt.Println("Created directory:", currentDir)
			} else if level == 2 {
				title = strings.TrimPrefix(title, "## ")
				if currentDir == "" {
					fmt.Println("Error: No directory found for section:", title)
					return fmt.Errorf("no directory found for section: %s", title)
				}
				currentFileName = filepath.Join(currentDir, title+".md")
				currentFile, err = os.Create(currentFileName)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return err
				}
				fmt.Println("Created file:", currentFileName)
			}
		}

		if currentFile != nil {
			_, err := currentFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return err
			}
			fmt.Println("Written to file:", currentFileName)
		}
	}
	if currentFile != nil {
		currentFile.Close()
	}
	return scanner.Err()
}

func addHyperlinks(inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var (
		currentDir      string
		currentFileName string
	)
	newFile, err := os.Create("new_" + inputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer newFile.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") {

			_, err := newFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return err
			}

			title := strings.TrimPrefix(line, "# ")
			title = strings.TrimSpace(title)
			level := strings.Count(line, "#")
			if level == 1 {
				currentDir = filepath.Join(".", title)
			} else if level == 2 {
				title = strings.TrimPrefix(title, "## ")
				if currentDir == "" {
					fmt.Println("Error: No directory found for section:", title)
					return fmt.Errorf("no directory found for section: %s", title)
				}
				currentFileName = filepath.Join(currentDir, title+".md")
				_, err = newFile.WriteString("[" + title + "](" + currentFileName + ")\n")
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return err
				}
			}
		}
	}
	return scanner.Err()
}
