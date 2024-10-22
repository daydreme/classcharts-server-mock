package user

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/parent"
	"github.com/CommunityCharts/CCModels/school"
	"github.com/CommunityCharts/CCModels/shared"
	"github.com/CommunityCharts/CCModels/student"
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

	data := shared.Object{
		"user": parent.NewParentUser(1, "Jamie Doe", "james@example.com", "en", true),
	}

	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPupilsHandler(w http.ResponseWriter, _ *http.Request) {
	response := shared.NewSuccessfulResponse(parent.NewPupilsFromStudentsAndSchool([]student.StudentUser{
		student.NewUser(1, "Jeremy Kyle", "https://placehold.co/340"),
	}, school.NewSchool(1, "Test School", "en")))
	response.Write(w)
}
