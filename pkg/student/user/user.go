package user

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/CommunityCharts/CCModels/student"
)

type userResponseMeta struct {
	Version   *string `json:"version,omitempty"`
	SessionId *string `json:"session_id,omitempty"`
}

const version = "20.0.1"

func StudentUserHandler(w http.ResponseWriter, r *http.Request) {
	var sessionId *string

	if r.FormValue("include_data") != "false" {
	}

	version := version

	meta := userResponseMeta{
		Version:   &version,
		SessionId: sessionId,
	}

	data := shared.Object{
		"user": student.NewUser(1, "Johnny Kyle", "https://placehold.co/320"),
	}

	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
