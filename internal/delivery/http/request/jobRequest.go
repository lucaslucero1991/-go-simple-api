package request

type JobRequest struct {
	Name    string   `json:"name"`
	Country string   `json:"country"`
	Salary  int      `json:"salary"`
	Skills  []string `json:"skills"`
}
