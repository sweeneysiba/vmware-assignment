package models

import (
	"encoding/json"
	"errors"
	"os"
	"vmware/forms"
	"vmware/httpaction"
)

type AssignmentModel struct{}

type AssignmentResult struct {
	Data  forms.Assignments `json:"data"`
	Count int               `json:"count"`
}

var httpAction = new(httpaction.HTTPAction)

func (m AssignmentModel) Get() (AssignmentResult, error) {
	var result AssignmentResult
	URLs := []string{
		os.Getenv("URL1"),
		os.Getenv("URL2"),
		os.Getenv("URL3"),
	}
	ch := make(chan httpaction.ConcurrentResult)
	for _, url := range URLs {
		go httpAction.GetCuncurrency(url, 5, ch) //retry for max 5 times
	}
	for range URLs {
		resp := <-ch
		if resp.Error != nil {
			return result, resp.Error
		}
		var response AssignmentResult
		err := json.Unmarshal(resp.Result, &response)
		if err != nil {
			return result, err
		}
		result.Data = append(result.Data, response.Data...)
	}
	result.Count = len(result.Data)
	return result, nil
}

//SortGetAssignmentResult ...
func (m AssignmentModel) SortGetAssignmentResult(assignments forms.Assignments, sortKey string) (forms.Assignments, error) {
	var n = len(assignments)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			switch sortKey {
			case "url":
				if assignments[j-1].URL > assignments[j].URL {
					assignments[j-1], assignments[j] = assignments[j], assignments[j-1]
				}
			case "views":
				if assignments[j-1].Views > assignments[j].Views {
					assignments[j-1], assignments[j] = assignments[j], assignments[j-1]
				}
			case "relevanceScore":
				if assignments[j-1].RelevanceScore > assignments[j].RelevanceScore {
					assignments[j-1], assignments[j] = assignments[j], assignments[j-1]
				}
			default:
				return assignments, errors.New("wrong sort key")
			}
			j = j - 1
		}
	}

	return assignments, nil
}

//LimitGetAssignmentResult ...
func (m AssignmentModel) LimitGetAssignmentResult(assignments forms.Assignments, limit int) forms.Assignments {
	if len(assignments) > limit {
		assignments = assignments[:limit]
	}
	return assignments
}
