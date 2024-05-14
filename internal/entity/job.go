package entity

type Job struct {
	ID      int
	Name    string
	Country string
	Salary  int
	Skills  string
}

func NewJob(name string, country string, salary int, skills string) *Job {
	return &Job{
		Name:    name,
		Country: country,
		Salary:  salary,
		Skills:  skills,
	}
}
