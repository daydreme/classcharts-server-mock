package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/daydreme/classcharts-server-mock/pkg/models"
	"github.com/daydreme/classcharts-server-mock/pkg/responses"
)

const validCode = "test"
const validDOB = "2024-01-01"

type hasDOBResponse struct {
	HasDOB bool `json:"has_dob"`
}

func HasDOBHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	data := responses.SuccesfulResponse[hasDOBResponse, []responses.Object]{
		Data: hasDOBResponse{
			HasDOB: strings.ToLower(code) == validCode,
		},
		Meta:    []responses.Object{},
		Success: 1,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	// remember := r.FormValue("remember")
	recaptchaToken := r.FormValue("recaptcha-token")

	if recaptchaToken != "no-token-available" {
		panic("recaptchaToken is invalid")
	}

	if strings.ToLower(code) != validCode {
		response := responses.NewErrorfulResponse("The pupil code you have provided is incorrect If you do not have your pupil code, or have forgotten it, please contact your school. Your school contact details can usually be found on your school's website.")
		response.Write(w)
		return
	}

	dob := r.FormValue("dob")

	if dob != validDOB {
		response := responses.NewErrorfulResponse("The date of birth you have provided is incorrect")
		response.Write(w)
		return
	}

	// Not 100% parity here because we are returning the whole user object while CC only returns a subset for some reason
	data := models.NewMockUser()

	globalSessionId := globalSessionId
	meta := userResponseMeta{
		SessionId: &globalSessionId,
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetCodeHandler(w http.ResponseWriter, r *http.Request) {
	dob := r.FormValue("date")

	if dob != validDOB {
		response := responses.NewErrorfulResponse("The date you provided is invalid.")
		response.Write(w)
		return
	}

	response := responses.NewSuccessfulResponse(map[string]interface{}{
		"code": validCode,
	})
	response.Write(w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
