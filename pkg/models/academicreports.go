package models

type PreviewAcademicReport struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type AcademicReportLessonGrade struct {
	FieldName string `json:"field_name"`
	Grade     string `json:"grade"`
}

type AcademicReportLessonTeacher struct {
	TeacherName string `json:"teacher_name"`
	Id          int    `json:"id"`
	Text        string `json:"text"`
}

type AcademicReportLesson struct {
	AttendancePercentage *string                       `json:"attendance_percentage"`
	LessonName           string                        `json:"lesson_name"`
	Teachers             []AcademicReportLessonTeacher `json:"teachers"`
	Grades               []AcademicReportLessonGrade   `json:"grades"`
}

type AcademicReport struct {
	Id                          int                    `json:"id"`
	Name                        string                 `json:"name"`
	StartDate                   string                 `json:"start_date"`
	EndDate                     string                 `json:"end_date"`
	OverallAttendancePercentage *string                `json:"overall_attendance_percentage"`
	AttendancePercentage        *string                `json:"attendance_percentage"`
	Lessons                     []AcademicReportLesson `json:"lessons"`
}

func NewMockPreviewAcademicReports() []PreviewAcademicReport {
	return []PreviewAcademicReport{
		// TODO: fix the date here
		{Id: 1, Name: "Midterm Report", StartDate: "2023-01-01", EndDate: "2023-01-31"},
		{Id: 2, Name: "Quarterly Report", StartDate: "2023-04-01", EndDate: "2023-06-30"},
		{Id: 3, Name: "Final Report", StartDate: "2023-09-01", EndDate: "2023-12-31"},
	}
}

func NewMockAcademicReports() []AcademicReport {
	attendance := "95"
	return []AcademicReport{
		{
			Id:                          1,
			Name:                        "Midterm Report",
			StartDate:                   "2023-01-01",
			EndDate:                     "2023-01-31",
			OverallAttendancePercentage: &attendance,
			AttendancePercentage:        &attendance,
			Lessons: []AcademicReportLesson{
				{
					AttendancePercentage: &attendance,
					LessonName:           "Mathematics",
					Teachers: []AcademicReportLessonTeacher{
						{TeacherName: "Jane Smith", Id: 101, Text: "Excellent performance"},
						{TeacherName: "John Brown", Id: 102, Text: "Needs improvement in calculus"},
					},
					Grades: []AcademicReportLessonGrade{
						{FieldName: "Homework", Grade: "A"},
						{FieldName: "Midterm Exam", Grade: "B+"},
					},
				},
				{
					LessonName: "Science",
					Teachers: []AcademicReportLessonTeacher{
						{TeacherName: "Albert Johnson", Id: 103, Text: "Great interest in physics"},
					},
					Grades: []AcademicReportLessonGrade{
						{FieldName: "Lab Work", Grade: "A-"},
						{FieldName: "Final Exam", Grade: "A"},
					},
				},
			},
		},
		{
			Id:                          2,
			Name:                        "Quarterly Report",
			StartDate:                   "2023-04-01",
			EndDate:                     "2023-06-30",
			OverallAttendancePercentage: &attendance,
			AttendancePercentage:        &attendance,
			Lessons: []AcademicReportLesson{
				{
					AttendancePercentage: &attendance,
					LessonName:           "History",
					Teachers: []AcademicReportLessonTeacher{
						{TeacherName: "Emily White", Id: 104, Text: "Shows excellent understanding of historical events"},
					},
					Grades: []AcademicReportLessonGrade{
						{FieldName: "Essay", Grade: "A"},
						{FieldName: "Oral Presentation", Grade: "B+"},
					},
				},
			},
		},
		{
			Id:                          3,
			Name:                        "Final Report",
			StartDate:                   "2023-09-01",
			EndDate:                     "2023-12-31",
			OverallAttendancePercentage: &attendance,
			AttendancePercentage:        &attendance,
			Lessons: []AcademicReportLesson{
				{
					AttendancePercentage: &attendance,
					LessonName:           "Computer Science",
					Teachers: []AcademicReportLessonTeacher{
						{TeacherName: "Thomas Wayne", Id: 105, Text: "Excellent coding skills"},
					},
					Grades: []AcademicReportLessonGrade{
						{FieldName: "Project", Grade: "A+"},
						{FieldName: "Final Exam", Grade: "A"},
					},
				},
			},
		},
	}
}
