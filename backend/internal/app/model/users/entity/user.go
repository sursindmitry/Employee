package entity

type User struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	JobTitle       string `json:"job_title"`
	Email          string `json:"email"`
	EmailPassword  string `json:"email_password"`
	DomainName     string `json:"domain_name"`
	DomainPassword string `json:"domain_password"`
	Location       string `json:"location"`
	Phone          string `json:"phone"`
}
