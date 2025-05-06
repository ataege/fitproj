package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetHealth returns the health status of the API
func GetHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "healthy",
	})
}

// GetHello returns a simple hello message
func GetHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

// UploadFile handles file uploads and saves them to the uploads directory
func UploadFile(c *fiber.Ctx) error {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error retrieving file from request",
			"error":   err.Error(),
		})
	}

	// Create uploads directory if it doesn't exist
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating upload directory",
			"error":   err.Error(),
		})
	}

	// Generate a unique filename to prevent overwriting existing files
	timestamp := time.Now().Unix()
	uniqueFilename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
	filePath := filepath.Join(uploadDir, uniqueFilename)

	// Save the file
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error saving file",
			"error":   err.Error(),
		})
	}

	// Return the file path to the client
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "File uploaded successfully",
		"data": fiber.Map{
			"filename": uniqueFilename,
			"size":     file.Size,
			"path":     filePath,
		},
	})
}

// UploadMultipleFiles handles multiple file uploads
func UploadMultipleFiles(c *fiber.Ctx) error {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error parsing multipart form",
			"error":   err.Error(),
		})
	}

	// Get all files from the form
	files := form.File["files"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "No files were provided",
		})
	}

	// Create uploads directory if it doesn't exist
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating upload directory",
			"error":   err.Error(),
		})
	}

	uploadedFiles := make([]fiber.Map, 0, len(files))

	// Save each file
	for _, file := range files {
		// Generate a unique filename
		timestamp := time.Now().UnixNano()
		uniqueFilename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
		filePath := filepath.Join(uploadDir, uniqueFilename)

		// Save the file
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Error saving file",
				"error":   err.Error(),
			})
		}

		// Add file info to the response
		uploadedFiles = append(uploadedFiles, fiber.Map{
			"filename": uniqueFilename,
			"size":     file.Size,
			"path":     filePath,
		})
	}

	// Return information about all uploaded files
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("%d files uploaded successfully", len(files)),
		"data":    uploadedFiles,
	})
}

// ListFiles returns a list of all uploaded files
func ListFiles(c *fiber.Ctx) error {
	// Set the directory to list files from
	uploadDir := "./uploads"

	// Check if directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "No uploads directory found",
			"data":    []string{},
		})
	}

	// Read the directory
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error reading uploads directory",
			"error":   err.Error(),
		})
	}

	// Create a slice to store file information
	filesList := make([]fiber.Map, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Get file info
		fileInfo, err := file.Info()
		if err != nil {
			continue // Skip files that can't be read
		}

		filesList = append(filesList, fiber.Map{
			"filename": file.Name(),
			"size":     fileInfo.Size(),
			"modified": fileInfo.ModTime(),
		})
	}

	// Return the list of files
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Found %d files", len(filesList)),
		"data":    filesList,
	})
}

// DownloadFile allows downloading a specific file by filename
func DownloadFile(c *fiber.Ctx) error {
	// Get the filename from the request params
	filename := c.Params("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Filename is required",
		})
	}

	// Set the path to the file
	filePath := filepath.Join("./uploads", filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "File not found",
		})
	}

	// Return the file
	return c.Download(filePath)
}
