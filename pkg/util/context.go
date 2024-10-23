package util

import (
	"github.com/CommunityCharts/CCModels/student"
	"net/http"
)

func GetStudent(r *http.Request) student.DBStudentUser {
	return r.Context().Value("student").(student.DBStudentUser)
}
