package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nchikkam/context-findr-be/model"
	"github.com/nchikkam/context-findr-be/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type SearchInput struct {
	Input string
}

// @Summary		Search context for give word
// @Description	search in file
// @Param 		q    query     string  false  "search by word"
// @Produce		json
// @Security 	BearerAuth
// @Success		200
// @Router		/api/v1/search [get]
func Search(c *gin.Context) {

	query := c.Query("q")

	if len(query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Search input"})
		return
	}

	email := c.GetString("userEmail")

	var userFile model.File
	err := getFileCollection().FindOne(c, bson.M{"email": email}).Decode(&userFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("files collection err : %v", err)})
		return
	}

	fileName := utils.Store + userFile.Name

	// todo: create go module for grepping
	cmd := exec.Command("grep", "-n", query, fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("search encoutered some error: %v", err)})
		return
	}

	results := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		var snippet string = scanner.Text()
		match := strings.Split(snippet, ":")

		number := match[0]
		line := strings.Join(match[1:], "")

		results[number] = line
	}

	c.JSON(http.StatusOK, gin.H{
		"matches": results,
	})
}
