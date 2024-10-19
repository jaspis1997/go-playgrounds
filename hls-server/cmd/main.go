package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handleIndex)
	router.POST("/upload", handleUpload)
	router.Static("/hls", "./hls/output")
	router.Static("/public", "./public")

	fmt.Println("Server is running on http://localhost:8080")
	router.Run("localhost:8080")
}

func handleIndex(c *gin.Context) {
	// Indexページの処理をここに追加
	c.File("./cmd/index.html")
}

func handleUpload(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "upload-*.mp4")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	go func() {
		if err := convertToHLS(tempFile.Name(), "output"); err != nil {
			fmt.Printf("Error converting to HLS: %v\n", err)
		}
	}()

	c.String(200, "File uploaded successfully")
}

func convertToHLS(inputPath, outputName string) error {
	outputDir := filepath.Join("hls", outputName)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, "playlist.m3u8")

	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-f", "hls",
		outputPath)

	return cmd.Run()
}
