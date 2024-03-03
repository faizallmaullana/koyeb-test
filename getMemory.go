package main

import (
	"bufio"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func GetMemory(c *gin.Context) {
	cmd := exec.Command("df", "-h")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Error creating StdoutPipe",
		})
		return
	}
	if err := cmd.Start(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Error starting command",
		})
		return
	}

	scanner := bufio.NewScanner(stdout)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Error reading standard input",
		})
	}

	if err := cmd.Wait(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Error waiting for command",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"disk_usage": output,
	})
}
