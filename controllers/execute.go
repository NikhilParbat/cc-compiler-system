package controllers

import (
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/NikhilParbat/CC-Compiler-Go/models"

	"github.com/gin-gonic/gin"
)

func ExecuteCode(c *gin.Context) {
	var codeReq models.CodeRequest
	if err := c.BindJSON(&codeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cmd *exec.Cmd
	switch codeReq.Language {
	case "js":
		cmd = exec.Command("node", "-e", codeReq.Code)
	case "py":
		cmd = exec.Command("python", "-")
		cmd.Stdin = strings.NewReader(codeReq.Code)
	case "rb":
		cmd = exec.Command("ruby", "-e", codeReq.Code)
	case "php":
		cmd = exec.Command("php", "-r", codeReq.Code)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported language"})
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.CodeResponse{
		Output: string(output),
		Error:  "",
	}
	c.JSON(http.StatusOK, response)

	// Remove the temporary executable file if it exists
	if _, err := os.Stat("./temp"); err == nil {
		os.Remove("./temp")
	}
}
