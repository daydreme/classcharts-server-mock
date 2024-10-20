package parent

import (
	"github.com/daydreme/classcharts-server-mock/pkg/handlers/db"
	"github.com/daydreme/classcharts-server-mock/pkg/models/student"
)

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Language        string `json:"language"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

type Pupil struct {
	student.User

	SchoolName string `json:"school_name"`
	SchoolLogo string `json:"school_logo"`

	Timezone string `json:"timezone"`

	DisplayCovidTests   bool `json:"display_covid_tests"`
	CanRecordCovidTests bool `json:"can_record_covid_tests"`

	DetentionYesCount      int `json:"detention_yes_count"`
	DetentionNoCount       int `json:"detention_no_count"`
	DetentionPendingCount  int `json:"detention_pending_count"`
	DetentionUpscaledCount int `json:"detention_upscaled_count"`

	HomeworkTodoCount         int `json:"homework_todo_count"`
	HomeworkLateCount         int `json:"homework_late_count"`
	HomeworkNotCompletedCount int `json:"homework_not_completed_count"`
	HomeworkExcusedCount      int `json:"homework_excused_count"`
	HomeworkCompletedCount    int `json:"homework_completed_count"`
	HomeworkSubmittedCount    int `json:"homework_submitted_count"`
}

func NewMockUser() User {
	return User{
		Id:              1,
		Name:            "Jane Doe",
		Email:           "jane@example.com",
		Language:        "en",
		IsEmailVerified: true,
	}
}

func NewMockPupils() []Pupil {
	var pupils []Pupil

	students := db.GetStudents()

	for _, studentDB := range students {
		pupils = append(pupils, Pupil{
			User: student.NewMockUserFromStudentDB(studentDB),

			SchoolName: "Primmit Secondary School",
			SchoolLogo: "https://via.placeholder.com/480",

			Timezone: "Europe/London",

			DisplayCovidTests:   true,
			CanRecordCovidTests: true,

			DetentionYesCount:      2,
			DetentionNoCount:       1,
			DetentionPendingCount:  4,
			DetentionUpscaledCount: 3,

			HomeworkTodoCount:         5,
			HomeworkLateCount:         1,
			HomeworkNotCompletedCount: 2,
			HomeworkExcusedCount:      1,
			HomeworkCompletedCount:    3,
			HomeworkSubmittedCount:    4,
		})
	}

	return pupils
}
