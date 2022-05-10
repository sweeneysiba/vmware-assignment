package controllers

import (
	"net/http"
	"vmware/forms"
	"vmware/models"

	"github.com/gin-gonic/gin"
)

type AssignmentController struct{}

var assignmentModel = new(models.AssignmentModel)

//Create ...
func (ctrl AssignmentController) Get(c *gin.Context) {

	var form forms.GetAssignmentForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "Invalid Request", "message": validationErr.Error()})
		return
	}
	result, err := assignmentModel.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Assignment could not be fetched", "error": err.Error()})
		return
	}
	result.Data = assignmentModel.LimitGetAssignmentResult(result.Data, int(form.Limit))
	result.Data, err = assignmentModel.SortGetAssignmentResult(result.Data, form.Sort)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Assignment could not be Sort", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Assignment fetched Successfully", "data": result.Data, "count": result.Count})
}
