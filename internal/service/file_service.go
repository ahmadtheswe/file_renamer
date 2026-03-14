package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileService interface {
	BuildNewFileName(prefix string, totalFiles int, correctedPrefixCount int, index int, extension string) string
	CheckFileExtension(fileExtension string) string
	CountFileWithCorrectPrefix(files []os.DirEntry, prefix string) int
	FilterFilesByExtensionAndPrefix(entries []os.DirEntry, extension string, prefix string) []string
	RenameFiles(fileName string, newFileName string, dirPath string, isVerbose bool) error
}

func NewFileService() FileService {
	return &fileServiceImpl{}
}

func (s *fileServiceImpl) BuildNewFileName(prefix string, totalFiles int, correctedPrefixCount int, index int, extension string) string {
	return prefix + "_" + addZeroPrefix(totalFiles+correctedPrefixCount, index) + extension
}

func addZeroPrefix(totalFiles int, index int) string {
	digits := len(fmt.Sprintf("%d", totalFiles))
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, index)
}

func (s *fileServiceImpl) CheckFileExtension(fileExtension string) string {
	if fileExtension == "" {
		return fileExtension
	}
	if len(fileExtension) > 0 && fileExtension[0] == '.' {
		return fileExtension
	}
	return "." + fileExtension
}

func (s *fileServiceImpl) FilterFilesByExtensionAndPrefix(entries []os.DirEntry, extension string, prefix string) []string {
	var filtered []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) == extension {
			filtered = append(filtered, entry.Name())
		}

		if extension == "" {
			filtered = append(filtered, entry.Name())
		}
	}

	// Remove files that already have the prefix
	if prefix != "" {
		result := filtered[:0]
		for _, file := range filtered {
			if !strings.HasPrefix(file, prefix) {
				result = append(result, file)
			}
		}
		filtered = result
	}
	return filtered
}

type fileServiceImpl struct{}

func (s *fileServiceImpl) CountFileWithCorrectPrefix(files []os.DirEntry, prefix string) int {
	count := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			count++
		}
	}
	return count
}

func (s *fileServiceImpl) RenameFiles(fileName string, newFileName string, dirPath string, isVerbose bool) error {
	if isVerbose {
		fmt.Printf("Renaming '%s' to '%s'\n", fileName, newFileName)
	}
	err := os.Rename(filepath.Join(dirPath, fileName), filepath.Join(dirPath, newFileName))
	if err != nil {
		fmt.Printf("Error renaming file '%s': %v\n", fileName, err)
	}
	return err
}
