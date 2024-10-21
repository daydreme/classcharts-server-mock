package models

type Lesson struct {
	TeacherName         string `json:"teacher_name"`
	TeacherID           string `json:"teacher_id"`
	LessonName          string `json:"lesson_name"`
	SubjectName         string `json:"subject_name"`
	IsAlternativeLesson bool   `json:"is_alternative_lesson"`
	IsBreak             *bool  `json:"is_break,omitempty"`
	PeriodName          string `json:"period_name"`
	PeriodNumber        string `json:"period_number"`
	RoomName            string `json:"room_name"`
	Date                string `json:"date"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	Key                 int    `json:"key"`
	NoteAbstract        string `json:"note_abstract"`
	Note                string `json:"note"`
	PupilNoteAbstract   string `json:"pupil_note_abstract"`
	PupilNote           string `json:"pupil_note"`
	PupilNoteRaw        string `json:"pupil_note_raw"`
}

type Period struct {
	Number    string `json:"number"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func NewMockPeriods() []Period {
	return []Period{
		{Number: "1", StartTime: "08:00", EndTime: "09:00"},
		{Number: "2", StartTime: "09:00", EndTime: "10:00"},
		{Number: "3", StartTime: "10:00", EndTime: "11:00"},
	}
}

func NewMockLessons() []Lesson {
	isBreakTrue := true
	isBreakFalse := false

	return []Lesson{
		{
			TeacherName:         "John Doe",
			TeacherID:           "T123",
			LessonName:          "Math",
			SubjectName:         "Mathematics",
			IsAlternativeLesson: false,
			IsBreak:             &isBreakFalse,
			PeriodName:          "Morning",
			PeriodNumber:        "1",
			RoomName:            "101",
			Date:                "2024-05-29",
			StartTime:           "2024-05-29T08:00:00+00:00",
			EndTime:             "2024-05-29T09:00:00+00:00",
			Key:                 1,
			NoteAbstract:        "Introduction to Algebra",
			Note:                "Detailed lesson plan on Algebra",
			PupilNoteAbstract:   `Small preview note...`,
			PupilNote:           "Big note...",
			PupilNoteRaw:        "Big note but raw?",
		},
		{
			TeacherName:         "Jane Smith",
			TeacherID:           "T124",
			LessonName:          "Science",
			SubjectName:         "Biology",
			IsAlternativeLesson: true,
			IsBreak:             &isBreakTrue,
			PeriodName:          "Afternoon",
			PeriodNumber:        "3",
			RoomName:            "102",
			Date:                "2024-05-29",
			StartTime:           "2024-05-29T09:00:00+00:00",
			EndTime:             "2024-05-29T10:00:00+00:00",
			Key:                 2,
			NoteAbstract:        "Introduction to Cells",
			Note:                "Detailed lesson plan on Cell Biology",
			PupilNoteAbstract:   "Cell basics",
			PupilNote:           "Students will learn the basics of Cell Biology",
			PupilNoteRaw:        "Learn about Cells",
		},
		{
			TeacherName:         "Michael Brown",
			TeacherID:           "T125",
			LessonName:          "History",
			SubjectName:         "World History",
			IsAlternativeLesson: false,
			IsBreak:             nil, // This will be omitted in JSON
			PeriodName:          "Evening",
			PeriodNumber:        "5",
			RoomName:            "103",
			Date:                "2024-05-29",
			StartTime:           "2024-05-29T10:00:00+00:00",
			EndTime:             "2024-05-29T11:00:00+00:00",
			Key:                 3,
			NoteAbstract:        "Ancient Civilizations",
			Note:                "Detailed lesson plan on Ancient Civilizations<script>alert(1)</script>",
			PupilNoteAbstract:   "Civilization basics",
			PupilNote:           "Students will learn about Ancient Civilizations",
			PupilNoteRaw:        "Learn about Ancient Civilizations",
		},
	}
}
