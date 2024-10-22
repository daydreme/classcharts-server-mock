package data

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
)

func GetAnnouncementHandler(w http.ResponseWriter, _ *http.Request) {
	data := []shared.Announcement{
		{
			Id:          1,
			Title:       "Welcome to the new school year!",
			Description: "We are excited to welcome all students back to school for the 2020-2021 school year. We have a lot of exciting things planned for this year, and we can't wait to get started!",
		},
	}

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}
