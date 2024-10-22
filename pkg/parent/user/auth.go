package user

import (
	"fmt"
	"net/http"

	"github.com/CommunityCharts/CCModels/parent"
	"github.com/CommunityCharts/CCModels/shared"
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
		response := shared.NewErrorfulResponse("Email address or password provided is incorrect")
		response.Write(w)
		return
	}

	data := parent.NewParentUser(1, "Jamie Doe", "james@example.com", "en", true)

	globalSessionId := globalSessionId
	meta := UserResponseMeta{
		SessionId: &globalSessionId,
	}

	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func LogoutHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
