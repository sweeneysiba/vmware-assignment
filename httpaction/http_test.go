package httpaction

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"vmware/forms"

	"github.com/joho/godotenv"
)

type assignmentResult struct {
	Data  forms.Assignments `json:"data"`
	Count int               `json:"count"`
}

func TestGet(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	httpaction := new(HTTPAction)
	resp, err := httpaction.Get(os.Getenv("URL1"))
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	var response assignmentResult
	err = json.Unmarshal(resp, &response)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	if len(response.Data) != 5 {
		t.Fatalf("Expected count %d but got %d", 5, len(response.Data))
	}
}

func TestGetCuncurrency(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	httpaction := new(HTTPAction)
	requests := []string{
		os.Getenv("URL1"), os.Getenv("URL2"), os.Getenv("URL3"),
	}
	chanl := make(chan ConcurrentResult)
	for _, val := range requests {
		go httpaction.GetCuncurrency(val, 5, chanl)
	}
	var result forms.Assignments
	for range requests {
		resp := <-chanl
		if resp.Error != nil {
			t.Fatalf("Expected no error but got %v", err)
		}
		var response assignmentResult
		err := json.Unmarshal(resp.Result, &response)
		if err != nil {
			t.Fatalf("Expected no error but got %v", err)
		}
		result = append(result, response.Data...)
	}
	if len(result) != 15 {
		t.Fatalf("Expected count %d but got %d", 15, len(result))
	}
}
