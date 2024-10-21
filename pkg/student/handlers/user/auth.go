package user

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/student/models"
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

	hasDOB := false
	students := global.GetStudents()
	for _, studentDB := range students {
		if strings.ToLower(code) == strings.ToLower(studentDB.Code) {
			hasDOB = true
			break
		}
	}

	response := responses.NewSuccessfulResponse(hasDOBResponse{
		HasDOB: hasDOB,
	})

	response.Write(w)
}

func CheckPupilCodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	hasDOB := false
	students := global.GetStudents()
	for _, studentDB := range students {
		if strings.ToLower(code) == strings.ToLower(studentDB.Code) {
			hasDOB = true
			break
		}
	}

	response := responses.NewSuccessfulResponse(hasDOBResponse{
		HasDOB: hasDOB,
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
	var filteredStudents []global.StudentDB
	for _, studentDB := range students {
		if strings.ToLower(code) == strings.ToLower(studentDB.Code) {
			filteredStudents = append(filteredStudents, studentDB)
		}
	}

	if len(filteredStudents) == 0 {
		response := responses.NewErrorfulResponse("The pupil code you have provided is incorrect. If you do not have your pupil code, or have forgotten it, please contact your school. Your school contact details can usually be found on your school's website.")
		response.Write(w)
		return
	}

	validDOB := false
	for _, studentDB := range filteredStudents {
		if dob == studentDB.DOB {
			validDOB = true
			break
		}
	}

	if !validDOB {
		response := responses.NewErrorfulResponse("The date of birth you have provided is incorrect")
		response.Write(w)
		return
	}

	// Not 100% parity here because we are returning the whole user object while CC only returns a subset for some reason
	data := models.NewMockUser()

	globalSessionId := globalSessionId
	meta := UserResponseMeta{
		SessionId: &globalSessionId,
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetCodeHandler(w http.ResponseWriter, r *http.Request) {
	dob := r.FormValue("date")

	students := global.GetStudents()
	var studentCode string
	validDOB := false
	for _, studentDB := range students {
		if dob == studentDB.DOB {
			validDOB = true
			studentCode = studentDB.Code
			break
		}
	}

	if !validDOB {
		response := responses.NewErrorfulResponse("The date you provided is invalid.")
		response.Write(w)
		return
	}

	response := responses.NewSuccessfulResponse(responses.Object{
		"code": studentCode,
	})
	response.Write(w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
