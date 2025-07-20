package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nchikkam/context-findr-be/model"
	"github.com/nchikkam/context-findr-be/utils/classifiers"
	utils "github.com/nchikkam/context-findr-be/utils/infrastructure"
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
	results := classifiers.ExtractTextContext(fileName, query)

	_, err_exists := results["error"]
	if err_exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("search encoutered some error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"matches": results,
	})
}
