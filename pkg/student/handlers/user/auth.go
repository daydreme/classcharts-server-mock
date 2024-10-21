package user

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type hasDOBResponse struct {
	HasDOB bool `json:"has_dob"`
}

func HasDOBHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	//data := responses.SuccesfulResponse[hasDOBResponse, []responses.Object]{
	//	Data: hasDOBResponse{
	//		HasDOB: strings.ToLower(code) == validCode,
	//	},
	//	Meta:    []responses.Object{},
	//	Success: 1,
	//}

	students := global.GetStudents()
	students = util.Filter(students, func(student global.StudentDB) bool {
		return strings.ToLower(code) == strings.ToLower(student.Code)
	})

	response := responses.NewSuccessfulResponse(hasDOBResponse{
		HasDOB: len(students) > 0,
	})

	response.Write(w)
}

func CheckPupilCodeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	students := global.GetStudents()
	students = util.Filter(students, func(student global.StudentDB) bool {
		return strings.ToLower(code) == strings.ToLower(student.Code)
	})

	response := responses.NewSuccessfulResponse(hasDOBResponse{
		HasDOB: len(students) > 0,
	})

	response.Write(w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	dob := r.FormValue("dob")
	// remember := r.FormValue("remember")
	recaptchaToken := r.FormValue("recaptcha-token")

	if recaptchaToken != "no-token-available" {
		// 2 options here. Return a 500 and:
		//panic("recaptchaToken is invalid")

		// or we can:
		//response := responses.NewErrorfulResponse("recaptchaToken is invalid.")
		//response.Write(w)

		// for now, we just print out a warning.
		fmt.Println("\033[33mWarning: recaptchaToken will most likely be required in the future. Please make sure to request login with 'recaptchaToken=no-token-available'.\033[0m")
	}

	students := global.GetStudents()
	students = util.Filter(students, func(student global.StudentDB) bool {
		return strings.ToLower(code) == strings.ToLower(student.Code)
	})

	if len(students) == 0 {
		response := responses.NewErrorfulResponse("The pupil code you have provided is incorrect. If you do not have your pupil code, or have forgotten it, please contact your school. Your school contact details can usually be found on your school's website.")
		response.Write(w)
		return
	}

	students = util.Filter(students, func(student global.StudentDB) bool {
		return strings.ToLower(dob) == strings.ToLower(student.DOB)
	})

	if len(students) == 0 {
		response := responses.NewErrorfulResponse("The pupil code you have provided is incorrect. If you do not have your pupil code, or have forgotten it, please contact your school. Your school contact details can usually be found on your school's website.")
		response.Write(w)
		return
	}

	student := students[0]

	// Not 100% parity here because we are returning the whole user object while CC only returns a subset for some reason
	data := student.ToStudentUser()

	globalSessionId := globalSessionId
	meta := userResponseMeta{
		SessionId: &globalSessionId,
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetCodeHandler(w http.ResponseWriter, r *http.Request) {
	dob := r.FormValue("date")

	students := global.GetStudents()
	students = util.Filter(students, func(student global.StudentDB) bool {
		return strings.ToLower(dob) == strings.ToLower(student.DOB)
	})

	response := responses.NewSuccessfulResponse(responses.Object{
		"code": students[0].Code,
	})
	response.Write(w)
}

func LogoutHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
