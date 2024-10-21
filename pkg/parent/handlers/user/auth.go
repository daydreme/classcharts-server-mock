package user

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/parent/models"
	"net/http"
)

const validEmail = "example"
const validPassword = ""

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
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

	if email != validEmail || password != validPassword {
		response := responses.NewErrorfulResponse("Email address or password provided is incorrect")
		response.Write(w)
		return
	}

	data := models.NewMockUser()

	globalSessionId := globalSessionId
	meta := UserResponseMeta{
		SessionId: &globalSessionId,
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
