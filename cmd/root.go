package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ahmadtheswe/file_renamer/internal/service"
)

var dirPath string
var fileExtension string
var prefix string
var verbose bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: file_renamer --dir <path> --ext <extension> --prefix <prefix> [options]\n\n")
		fmt.Fprintf(os.Stderr, "Required flags:\n")
		fmt.Fprintf(os.Stderr, "  --dir string\n\tDirectory containing files to rename\n")
		fmt.Fprintf(os.Stderr, "  --ext string\n\tFile extension to filter by (e.g., .txt)\n")
		fmt.Fprintf(os.Stderr, "  --prefix string\n\tPrefix to add to renamed files (e.g., log_)\n\n")
		fmt.Fprintf(os.Stderr, "Optional flags:\n")
		fmt.Fprintf(os.Stderr, "  --v\tVerbose output\n")
		fmt.Fprintf(os.Stderr, "  --help\tShow this help message\n")
	}

	flag.StringVar(&dirPath, "dir", "", "Directory to list (required)")
	flag.StringVar(&fileExtension, "ext", "", "Filter files by extension, e.g., .txt (required)")
	flag.StringVar(&prefix, "prefix", "", "Prefix to add to files, e.g., 'log_' (required)")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()

	var missing []string
	if dirPath == "" {
		missing = append(missing, "-dir")
	}
	if fileExtension == "" {
		missing = append(missing, "-ext")
	}
	if prefix == "" {
		missing = append(missing, "-prefix")
	}
	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "error: missing required flags: %s\n\n", strings.Join(missing, ", "))
		flag.Usage()
		os.Exit(1)
	}
}

func Execute() {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contents of directory '%s':\n", dirPath)
	fileSvc := service.NewFileService()

	var checkedExtension = fileSvc.CheckFileExtension(fileExtension)
	var filteredFiles []string = fileSvc.FilterFilesByExtensionAndPrefix(entries, checkedExtension, prefix)
	totalFiles := len(filteredFiles)

	if totalFiles == 0 {
		if fileExtension != "" || prefix != "" {
			fmt.Printf("No files found with extension '%s' or prefix '%s'\n", checkedExtension, prefix)
		} else {
			fmt.Printf("No files found\n")
		}
		return
	}

	correctedPrefixCount := fileSvc.CountFileWithCorrectPrefix(entries, prefix)
	if correctedPrefixCount > 0 {
		fmt.Printf("Warning: %d files already have the prefix '%s'.\n", correctedPrefixCount, prefix)
	}

	fmt.Printf("Total files with extension '%s' that will be renamed: %d\n", checkedExtension, totalFiles)
	fmt.Print("Are you sure you want to rename these files? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
		fmt.Println("Editing cancelled.")
		return
	}

	for i, file := range filteredFiles {
		index := correctedPrefixCount + i + 1
		newFileName := fileSvc.BuildNewFileName(prefix, totalFiles, correctedPrefixCount, index, checkedExtension)
		fileSvc.RenameFiles(file, newFileName, dirPath, verbose)
	}

	fmt.Println("Renaming completed.")
}
