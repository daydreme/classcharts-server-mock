package models

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/parent/models"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Announcement struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	SchoolName  string `json:"school_name"`
	TeacherName string `json:"teacher_name"`
	SchoolLogo  string `json:"school_logo"`

	Sticky responses.YesNoBool `json:"sticky"`
	State  string              `json:"state"` // viewed, new

	Timestamp string `json:"timestamp"`

	Attachments []AnnouncementAttachment `json:"attachments"`

	CommentVisibility string `json:"comment_visibility"`

	AllowComments    responses.YesNoBool `json:"allow_comments"`
	AllowReactions   responses.YesNoBool `json:"allow_reactions"`
	AllowConsent     responses.YesNoBool `json:"allow_consent"`
	PriorityPinned   responses.YesNoBool `json:"priority_pinned"`
	RequiresConsent  responses.YesNoBool `json:"requires_consent"`
	CanChangeConsent bool                `json:"can_change_consent"`

	Consent       Consent        `json:"consent"`
	PupilConsents []PupilConsent `json:"pupil_consents"`
}

type AnnouncementAttachment struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type Consent struct {
	Id           int                 `json:"id"`
	ConsentGiven responses.YesNoBool `json:"consent_given"`
	Comment      string              `json:"comment"`
	ParentName   string              `json:"parent_name"`
}

type PupilConsent struct {
	Pupil            responses.Object `json:"pupil"`
	CanChangeConsent bool             `json:"can_change_consent"`
	Consent          Consent          `json:"consent"`
}

func TryReadAnnouncement(name string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filepath.Join(cwd, "pkg", "models", "global", "announcements", name+".html"))
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func NewMockAnnouncements() []Announcement {
	now := time.Now()

	return []Announcement{
		{
			Id:          1012,
			Title:       "Year 10/11 'An Inspector Calls' Trip at Mayflower Theatre",
			Description: TryReadAnnouncement("inspector_calls_trip"),

			SchoolName:  "Primmit Secondary School",
			TeacherName: "Miss S Prile",
			SchoolLogo:  "https://via.placeholder.com/480",

			Sticky: responses.No,
			State:  "viewed",

			Timestamp: now.AddDate(0, 0, -2).Format(time.RFC3339),

			Attachments: []AnnouncementAttachment{{11, "PermissionSlip.pdf", "https://example.com/inspector_calls_trip.pdf"}},

			CommentVisibility: "public",

			AllowComments:    responses.Yes,
			AllowReactions:   responses.Yes,
			AllowConsent:     responses.Yes,
			PriorityPinned:   responses.No,
			RequiresConsent:  responses.Yes,
			CanChangeConsent: true,

			Consent: Consent{
				95,
				responses.Yes,
				"I give my consent for my child to attend the trip.",
				models.NewMockUser().Name,
			},
			PupilConsents: []PupilConsent{
				{
					responses.Object{
						"id":         1,
						"first_name": "John",
						"last_name":  "Doe",
					},
					true,
					Consent{
						95,
						responses.Yes,
						"I give my consent for John to attend the trip.",
						models.NewMockUser().Name,
					},
				},
				{
					responses.Object{
						"id":         2,
						"first_name": "Jeff",
						"last_name":  "Doo",
					},
					true,
					Consent{
						96,
						responses.Yes,
						"I give my consent for Jeff to attend the trip.",
						models.NewMockUser().Name,
					},
				},
			},
		},
	}
}
