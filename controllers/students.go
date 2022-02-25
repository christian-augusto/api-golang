package controllers

import (
	"api-golang/models"
	"encoding/json"
	"net/http"
)

var students []*models.Student

// GET /students action
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	loadStudents()

	studentsBytes, _ := json.Marshal(students)

	w.Write(studentsBytes)
}

func loadStudents() {
	if students != nil {
		return
	}

	students = make([]*models.Student, 2)

	students[0] = models.NewStudent("1", "Nicole", "Louise Assis")
	students[1] = models.NewStudent("2", "Gustavo Márcio", "Henrique Mendes")
}
