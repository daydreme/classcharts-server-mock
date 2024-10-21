package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"net/http"
	"strconv"
)

func AddConcernHandler(w http.ResponseWriter, r *http.Request) {
	studentId, err := strconv.Atoi(r.FormValue("pupil_id"))
	if err != nil {
		panic(err)
	}

	concern := r.FormValue("text")

	student := global.GetStudentByID(studentId)
	student.Concerns = append(student.Concerns, concern)

	global.UpdateStudent(student)

	response := responses.NewSuccessfulResponse(responses.Object{})
	response.Write(w)
}
