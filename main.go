package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var dirPath string
var fileExtension string
var prefix string
var verbose bool

func init() {
	dirPath = "."
	fileExtension = ""
	prefix = ""

	flag.StringVar(&dirPath, "dir", ".", "Directory to list")
	flag.StringVar(&fileExtension, "ext", "", "Filter files by extension (e.g., .txt)")
	flag.StringVar(&prefix, "prefix", "", "Prefix to add to files (e.g., 'log_')")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()
}

func main() {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contents of directory '%s':\n", dirPath)

	var filteredFiles []string
	var checkedExtension = checkFileExtension(fileExtension)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) == checkedExtension {
			filteredFiles = append(filteredFiles, entry.Name())
		}

		if checkedExtension == "" {
			filteredFiles = append(filteredFiles, entry.Name())
		}
	}

	totalFiles := len(filteredFiles)

	if totalFiles == 0 {
		if fileExtension != "" {
			fmt.Printf("No files found\n")
		} else {
			fmt.Printf("No files found with extension '%s'\n", fileExtension)
		}
		return
	}

	fmt.Printf("Total files with extension '%s': %d\n", fileExtension, totalFiles)
	fmt.Print("Are you sure you want to rename these files? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
		fmt.Println("Editing cancelled.")
		return
	}

	for i, file := range filteredFiles {
		newFileName := prefix + "_" + addZeroPrefix(totalFiles, i+1) + checkedExtension
		if verbose {
			fmt.Printf("Renaming '%s' to '%s'\n", file, newFileName)
		}
		err := os.Rename(filepath.Join(dirPath, file), filepath.Join(dirPath, newFileName))
		if err != nil {
			fmt.Printf("Error renaming file '%s': %v\n", file, err)
		}
	}

	fmt.Println("Renaming completed.")
}

func checkFileExtension(fileExtension string) string {
	if fileExtension == "" {
		return fileExtension
	}
	if isFileNameHasDotPrefix(fileExtension) {
		return fileExtension
	} else {
		return "." + fileExtension
	}

}

func isFileNameHasDotPrefix(fileName string) bool {
	return len(fileName) > 0 && fileName[0] == '.'
}

func addZeroPrefix(totalFiles int, index int) string {
	digits := len(fmt.Sprintf("%d", totalFiles))
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, index)
}
