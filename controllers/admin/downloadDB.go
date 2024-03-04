package admin

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func DownloadDB(c *gin.Context) {
	// Execute the command to get the current directory
	output, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Println("Error:", err)
		c.String(http.StatusInternalServerError, "Failed to get current directory")
		return
	}

	// Convert output bytes to string and trim any leading/trailing whitespace
	currentDir := strings.TrimSpace(string(output))

	// Print the current directory
	fmt.Println("Current directory:", currentDir)

	// Construct the file path
	filePath := fmt.Sprintf("%s/zpd.db", currentDir)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer file.Close()

	// Set the appropriate headers for the file download
	c.Header("Content-Disposition", "attachment; filename=zpd.db")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")

	// Stream the file to the response writer
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to download file")
		return
	}
}
