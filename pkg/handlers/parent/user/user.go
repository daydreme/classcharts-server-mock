package user

import (
	"github.com/daydreme/classcharts-server-mock/pkg/models/parent"
	"github.com/daydreme/classcharts-server-mock/pkg/models/responses"
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
		"user": parent.NewMockUser(),
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPupilsHandler(w http.ResponseWriter, r *http.Request) {
	response := responses.NewSuccessfulResponse(parent.NewMockPupils())
	response.Write(w)
}
