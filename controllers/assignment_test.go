package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vmware/models"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Body     string
	Response models.AssignmentResult
	Error    ErrorResponse
	Code     int
}

type ErrorResponse struct {
	Error   string
	Message string
}

var requests = []Request{
	{
		Body: `{"sort":"url","limit":15}`,
		Response: models.AssignmentResult{
			Count: 15,
		},
		Code: 200,
	},
	{
		Body: `{"sort":"views","limit":5}`,
		Response: models.AssignmentResult{
			Count: 15,
		},
		Code: 200,
	},
	{
		Body: `{"sort":"relevanceScore","limit":5}`,
		Response: models.AssignmentResult{
			Count: 15,
		},
		Code: 200,
	},
	{
		Body: `{"sort":"wrongKey","limit":5}`,
		Error: ErrorResponse{
			Error:   "wrong sort key",
			Message: "Assignment could not be Sort",
		},
		Code: 406,
	},
	{Body: `{"sort":"url"}`, Error: ErrorResponse{
		Error:   "Invalid Request",
		Message: "Key: 'GetAssignmentForm.Limit' Error:Field validation for 'Limit' failed on the 'required' tag",
	},
		Code: 406,
	},
	{Body: `{"limit":5}`,
		Error: ErrorResponse{
			Error:   "Invalid Request",
			Message: "Key: 'GetAssignmentForm.Sort' Error:Field validation for 'Sort' failed on the 'required' tag",
		},
		Code: 406,
	},
}

func TestControllerGetAll(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	assignment := new(AssignmentController)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/v1/assignemt/get", assignment.Get)

	for _, request := range requests {
		loginPayload := strings.NewReader(request.Body)
		req, err := http.NewRequest(http.MethodPost, "/v1/assignemt/get", loginPayload)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		if w.Code != request.Code {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
		if request.Code != 200 {
			var ErrorResult ErrorResponse
			err = json.Unmarshal(body, &ErrorResult)
			if ErrorResult.Message != request.Error.Message {
				t.Fatalf("Expected to get Message %s but instead got %s\n", ErrorResult.Message, request.Error.Message)
			}
			if ErrorResult.Error != request.Error.Error {
				t.Fatalf("Expected to get Error %s but instead got %s\n", ErrorResult.Message, request.Error.Message)
			}
		} else {
			var result models.AssignmentResult
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if result.Count != 15 {
				t.Fatalf("Expected to get count %d but instead got %d\n", result.Count, 15)
			}
		}
	}
}
