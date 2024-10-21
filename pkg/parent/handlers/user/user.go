package user

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/parent/models"
	"net/http"
)

type UserResponseMeta struct {
	SessionId *string `json:"session_id,omitempty"`
}

const globalSessionId = "test_parent"

func ParentUserHandler(w http.ResponseWriter, r *http.Request) {
	var sessionId *string

	if r.FormValue("include_data") != "false" {
		globalSessionId := globalSessionId
		sessionId = &globalSessionId
	}

	meta := UserResponseMeta{
		SessionId: sessionId,
	}

	data := responses.Object{
		"user": models.NewMockUser(),
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPupilsHandler(w http.ResponseWriter, r *http.Request) {
	response := responses.NewSuccessfulResponse(models.NewMockPupils())
	response.Write(w)
}
