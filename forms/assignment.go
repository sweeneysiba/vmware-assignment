package forms

//AssignmentForm ...
type AssignmentForm struct{}

//GetAssignmentForm ...
type GetAssignmentForm struct {
	Sort  string `form:"sort" json:"sort" binding:"required"`
	Limit int64  `form:"limit" json:"limit" binding:"required"`
}

type Assignment struct {
	URL            string  `form:"url" json:"url" binding:"required"`
	Views          int64   `form:"views" json:"views" binding:"required"`
	RelevanceScore float64 `form:"relevance_score" json:"relevanceScore" binding:"required"`
}

type Assignments []Assignment

//Sort ...
func (f AssignmentForm) Sort(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the sort key"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//Limit ...
func (f AssignmentForm) Limit(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the limit value"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}
