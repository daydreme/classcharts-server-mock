package data

import (
	"net/http"
	"strconv"

	"github.com/CommunityCharts/CCModels/shared"
	db "github.com/daydreme/classcharts-server-mock/pkg"
)

func AddConcernHandler(w http.ResponseWriter, r *http.Request) {
	studentId, err := strconv.Atoi(r.FormValue("pupil_id"))
	if err != nil {
		panic(err)
	}

	concern := r.FormValue("text")

	student := db.GetStudentByID(studentId)
	student.Concerns = append(student.Concerns, concern)

	db.UpdateStudent(student)

	response := shared.NewSuccessfulResponse(shared.Object{})
	response.Write(w)
}
