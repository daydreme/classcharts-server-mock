package data

import (
	"net/http"
	"strconv"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/CommunityCharts/CCModels/student"
)

func GetAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	s := strconv.Itoa(r.Context().Value("student").(student.DBStudentUser).Id)

	data := []shared.Announcement{
		{
			Id:          1,
			Title:       "Welcome to the new school year, student ID " + s,
			Description: "We are excited to welcome all students back to school for the 2020-2021 school year. We have a lot of exciting things planned for this year, and we can't wait to get started!",
		},
	}

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}
