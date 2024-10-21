package models

type PreviewOnReportCard struct {
	Id int `json:"id"`
	// Moment Date Format: DD/MM/YYYY
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

type OnReportCardComment struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	// Moment Date Format: DD/MM/YYYY HH:mm
	Date    string `json:"date"`
	CanEdit bool   `json:"can_edit"`
}

type OnReportCardTargetPupilReportCard struct {
	Id int `json:"id"`
	// === 'ranking' ??
	RequiredTeacherAction *string `json:"required_teacher_action"`
}

type OnReportCardTargetPeriod struct {
	Id     int  `json:"id"`
	MarkId *int `json:"mark_id"`
	// YES/NO
	Met                *string `json:"met"`
	Label              string  `json:"label"`
	CommentDescription string  `json:"comment_description"`
	// Moment Date Format: DD/MM/YYYY HH:mm
	CommentCreated string  `json:"comment_created"`
	MarkRanking    *string `json:"mark_ranking"`
	TeacherName    *string `json:"teacher_name"`
}

type OnReportCardSummaryComment struct {
	Id           int    `json:"id"`
	ReportCardId int    `json:"report_card_id"`
	Text         string `json:"text"`
	TeacherName  string `json:"teacher_name"`
	// Moment Date Format: DD/MM/YYYY HH:mm
	Date string `json:"date"`
}

type OnReportCardTarget struct {
	Id              int                                `json:"id"`
	PupilReportCard *OnReportCardTargetPupilReportCard `json:"pupil_report_card"`
	Description     string                             `json:"description"`
}

type PreviewOnReportCardTarget struct {
	Id int `json:"id"`
}

type OnReportCard struct {
	Id int `json:"id"`
	// Moment Date Format: ddd DD/MM/YYYY
	StartDate      string                      `json:"start_date"`
	EndDate        string                      `json:"end_date"`
	Reason         string                      `json:"reason"`
	Level          *string                     `json:"level"`
	Teacher        *string                     `json:"teacher"`
	Description    string                      `json:"description"`
	ParentComments *[]OnReportCardComment      `json:"parent_comments"`
	Targets        []PreviewOnReportCardTarget `json:"targets"`
}

func NewMockPreviewOnReportCards() []PreviewOnReportCard {
	return []PreviewOnReportCard{
		{
			Id:          1,
			StartDate:   "01/06/2023",
			EndDate:     "03/06/2023",
			Reason:      "Lateness",
			Description: "Arrived 30 minutes late",
		},
		{
			Id:          2,
			StartDate:   "05/06/2023",
			EndDate:     "06/06/2023",
			Reason:      "On Call",
			Description: "Called in during class",
		},
		{
			Id:          3,
			StartDate:   "10/06/2023",
			EndDate:     "12/06/2023",
			Reason:      "Behavior",
			Description: "Disruptive behavior in class",
		},
	}
}

func NewMockOnReportCards() []OnReportCard {
	level := "SLT"
	teacher := "Mrs. Smith"
	comments := []OnReportCardComment{
		{
			Id:          1,
			Description: "My child has not done this :<(",
			Date:        "01/06/2023 09:00",
			CanEdit:     true,
		},
	}

	return []OnReportCard{
		{
			Id:             1,
			StartDate:      "2023-06-01",
			EndDate:        "2023-06-03",
			Reason:         "Lateness",
			Level:          &level,
			Teacher:        &teacher,
			Description:    "Student arrived 30 minutes late.",
			ParentComments: &comments,
			Targets: []PreviewOnReportCardTarget{
				{
					Id: 1,
				},
			},
		},
		{
			Id:             2,
			StartDate:      "05/06/2023",
			EndDate:        "06/06/2023",
			Reason:         "On Call",
			Level:          &level,
			Teacher:        &teacher,
			Description:    "Student was called during class.",
			ParentComments: &comments,
		},
		{
			Id:             3,
			StartDate:      "10/06/2023",
			EndDate:        "12/06/2023",
			Reason:         "Behavior",
			Level:          &level,
			Teacher:        &teacher,
			Description:    "Disruptive behavior in class.",
			ParentComments: &comments,
		},
	}
}

func NewMockOnReportCardTargetPeriods() []OnReportCardTargetPeriod {
	markId1 := 1
	markId2 := 2

	no := "no"
	yes := "yes"
	ranking := "ranking"
	teacherName := "Mrs. Smith"

	return []OnReportCardTargetPeriod{
		{
			Id:                 1,
			MarkId:             &markId1,
			Met:                &yes,
			Label:              "P1",
			CommentDescription: "Some comment",
			CommentCreated:     "01/06/2023 13:00",
			TeacherName:        &teacherName,
		},
		{
			Id:                 2,
			MarkId:             &markId2,
			Met:                &no,
			Label:              "P2",
			CommentDescription: "Another Some comment",
			CommentCreated:     "01/06/2023 12:00",
			TeacherName:        &teacherName,
		},
		{
			Id:                 3,
			Label:              "P3",
			CommentDescription: "Another Another Some comment",
			CommentCreated:     "01/06/2023",
			MarkRanking:        &ranking,
		},
	}
}

func NewMockOnReportCardTargets() []OnReportCardTarget {
	// requiredTeacherAction := "ranking"
	pupilReportCard := OnReportCardTargetPupilReportCard{
		Id:                    1,
		RequiredTeacherAction: nil,
	}

	return []OnReportCardTarget{
		{
			Id:              1,
			PupilReportCard: &pupilReportCard,
			Description:     "Your target is to not be late.",
		},
	}
}

func NewMockOnReportCardSummaryComments() []OnReportCardSummaryComment {
	return []OnReportCardSummaryComment{
		{
			Id:           1,
			ReportCardId: 1,
			TeacherName:  "Mrs. Smith",
			Text:         "Late to lesson in English",
			Date:         "01/06/2023 9:00",
		},
		{
			Id:           2,
			ReportCardId: 1,
			TeacherName:  "Mrs. Smith",
			Text:         "Late to lessson in Maths",
			Date:         "02/06/2023 10:00",
		},
	}
}
