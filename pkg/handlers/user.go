package handlers

import (
	"net/http"

	"github.com/daydreme/classcharts-server-mock/pkg/models"
	"github.com/daydreme/classcharts-server-mock/pkg/responses"
)

type userResponseMeta struct {
	Version   *string `json:"version,omitempty"`
	SessionId *string `json:"session_id,omitempty"`
}

const version = "20.0.1"
const globalSessionId = "test"

func UserHandler(w http.ResponseWriter, r *http.Request) {
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

	data := map[string]interface{}{
		"user": models.NewMockUser(),
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
