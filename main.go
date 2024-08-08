package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	inputFolder := "./original"    // Replace with the path to the folder containing original images.
	outputFolder := "./compressed" // Output folder where compressed images will be saved.

	err := os.Mkdir(outputFolder, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating output folder:", err)
		return
	}

	files, err := ioutil.ReadDir(inputFolder)
	if err != nil {
		fmt.Println("Error reading input folder:", err)
		return
	}

	totalOriginalSize := int64(0)
	totalCompressedSize := int64(0)

	for _, file := range files {
		if isImageFile(file) {
			originalPath := filepath.Join(inputFolder, file.Name())
			compressedPath := filepath.Join(outputFolder, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))+".jpg")

			err := convertToJPEG(originalPath, compressedPath)
			if err != nil {
				fmt.Printf("Error converting %s to JPEG: %s\n", file.Name(), err)
				continue
			}

			originalSize, _ := getImageSize(originalPath)
			compressedSize, _ := getImageSize(compressedPath)

			totalOriginalSize += originalSize
			totalCompressedSize += compressedSize

			sizeDifference := (originalSize - compressedSize) / (1024 * 1024) // Convert bytes to megabytes
			fmt.Printf("Converted %s: %d megabytes smaller\n", file.Name(), sizeDifference)
		}
	}

	totalOriginalSizeMB := totalOriginalSize / (1024 * 1024)
	totalCompressedSizeMB := totalCompressedSize / (1024 * 1024)
	totalDifferenceMB := (totalOriginalSize - totalCompressedSize) / (1024 * 1024)

	fmt.Printf("Total original size: %d megabytes\n", totalOriginalSizeMB)
	fmt.Printf("Total compressed size: %d megabytes\n", totalCompressedSizeMB)
	fmt.Printf("Total storage saved: %d megabytes\n", totalDifferenceMB)
}

func isImageFile(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	ext := filepath.Ext(info.Name())
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func convertToJPEG(originalPath, compressedPath string) error {
	// Replace this path with the full path to the vips.exe executable
	vipsPath := ".\\libvips\\bin\\vips.exe"

	cmd := exec.Command(vipsPath, "copy", originalPath, compressedPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("vips command failed: %s\nOutput: %s", err, string(output))
	}

	return nil
}

func getImageSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
