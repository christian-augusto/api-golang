package models

type Student struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewStudent(id string, firstName string, lastName string) *Student {
	return &Student{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
