package data

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/CommunityCharts/CCModels/student"
	"github.com/CommunityCharts/CCServerMock/pkg/db"
)

func GetAnnouncementsHandler(w http.ResponseWriter, r *http.Request) {
	sch, err := db.GetSchoolByID(r.Context().Value("student").(student.DBStudentUser).SchoolId)
	if err != nil {
		panic(err)
	}

	response := shared.NewSuccessfulResponse(sch.Announcements)
	response.Write(w)
}

func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	sch, err := db.GetSchoolByID(r.Context().Value("student").(student.DBStudentUser).SchoolId)
	if err != nil {
		panic(err)
	}

	db.MakeAnnouncement(
		sch,
		r.FormValue("title"),
		r.FormValue("content"),
		r.FormValue("teacher"),
		shared.No,
		shared.AnnouncementAttachment{},
	)

	response := shared.NewSuccessfulResponse(shared.Object{})
	response.Write(w)
}
