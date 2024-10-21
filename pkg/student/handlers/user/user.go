package user

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/student/models"
	"net/http"
)

type userResponseMeta struct {
	Version   *string `json:"version,omitempty"`
	SessionId *string `json:"session_id,omitempty"`
}

const version = "20.0.1"
const globalSessionId = "test"

func StudentUserHandler(w http.ResponseWriter, r *http.Request) {
	var sessionId *string

	if r.FormValue("include_data") != "false" {
		globalSessionId := globalSessionId
		sessionId = &globalSessionId
	}

	version := version

	meta := userResponseMeta{
		Version:   &version,
		SessionId: sessionId,
	}

	data := responses.Object{
		"user": models.NewUser(),
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
