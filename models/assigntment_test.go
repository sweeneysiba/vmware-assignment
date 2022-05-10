package models

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var assignmentModel = new(AssignmentModel)

func TestGet(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	assignmentResult, err := assignmentModel.Get()
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	if assignmentResult.Count != 15 {
		t.Fatalf("Expected no count %d but got %d", 15, assignmentResult.Count)
	}
	limitData := assignmentModel.LimitGetAssignmentResult(assignmentResult.Data, 10)
	if len(limitData) != 10 {
		t.Fatalf("Expected total no of assignment %d but got %d", 10, len(limitData))
	}
	sortUrlData, err := assignmentModel.SortGetAssignmentResult(assignmentResult.Data, "url")
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if sortUrlData[0].URL != "www.example.com/abc1" {
		t.Fatalf("Expected first url as %s but got %s", "www.example.com/abc1", sortUrlData[0].URL)
	}
	sortViewsData, err := assignmentModel.SortGetAssignmentResult(assignmentResult.Data, "views")
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if sortViewsData[0].Views != 1000 {
		t.Fatalf("Expected first view as %d but got %d", 1000, sortViewsData[0].Views)
	}
	sortRelevanceScoreData, err := assignmentModel.SortGetAssignmentResult(assignmentResult.Data, "relevanceScore")
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if sortRelevanceScoreData[0].RelevanceScore != 0.1 {
		t.Fatalf("Expected first url as %f but got %f", 0.1, sortRelevanceScoreData[0].RelevanceScore)
	}
}
