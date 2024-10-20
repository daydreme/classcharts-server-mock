package global

import "time"

type ActivityStyle struct {
	BorderColour *string `json:"border_color"`
	CustomClass  *string `json:"custom_class"`
}

type Activity struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Polarity string `json:"polarity"`
	Reason   string `json:"reason"`
	Score    int    `json:"score"`

	Timestamp string `json:"timestamp"`

	Style ActivityStyle `json:"style"`

	LessonName  *string `json:"lesson_name"`
	TeacherName *string `json:"teacher_name"`
	RoomName    *string `json:"room_name"`

	Note *string `json:"note"`

	DetentionDate     *string `json:"detention_date"`
	DetentionTime     *string `json:"detention_time"`
	DetentionLocation *string `json:"detention_location"`
	DetentionType     *string `json:"detention_type"`
}

// FAKE INFO:
//	Science - Miss I Aritus (10C/Sc1)
//	Maths - Mr M Smith (10C/Ma1)
//	Computer Science - Miss C Yite (10C/Cs1)
//	English - Miss S Prile (10C/En1)

//	General/Attendance/Other - Mr T Pott - (No class, no room)
//	Year Leaders Office - Detention Location

func NewMockActivities() []Activity {
	now := time.Now()

	stringPtr := func(s string) *string {
		return &s
	}

	return []Activity{
		{
			Id:       1,
			Type:     "behaviour",
			Polarity: "positive",
			Reason:   "Homework completion",
			Score:    5,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  stringPtr("10C/Ma1"),
			TeacherName: stringPtr("Mr M Smith"),
			RoomName:    stringPtr("X18"),

			Note: stringPtr("Well done on completing your Maths homework on time for the past 3 years without fail."),

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
		{
			Id:       2,
			Type:     "behaviour",
			Polarity: "positive",
			Reason:   "Outstanding project",
			Score:    2,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  stringPtr("10C/Sc1"),
			TeacherName: stringPtr("Miss I Aritus"),
			RoomName:    stringPtr("S12"),

			Note: stringPtr("John showcased an amazing science project today. It definitely blew us away! 👍🏻"),

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
		{
			Id:       3,
			Type:     "behaviour",
			Polarity: "positive",
			Reason:   "Excellent participation",
			Score:    3,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  stringPtr("10C/Cs1"),
			TeacherName: stringPtr("Miss C Yite"),
			RoomName:    stringPtr("X5"),

			Note: stringPtr("Amazing participation from all students today. Well done."),

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
		{
			Id:       4,
			Type:     "behaviour",
			Polarity: "positive",
			Reason:   "Perfect attendance",
			Score:    4,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  nil,
			TeacherName: stringPtr("Mr T Pott"),
			RoomName:    nil,

			Note: stringPtr("Well done, John, for coming in on-time every single day since Year 7. " +
				"We've all seen you struggle throughout times where you've been seriously ill yet still have had the confidence and passion to come into school. " +
				"If only more of the kids were like you! Keep it up!"),

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
		{
			Id:       5,
			Type:     "behaviour",
			Polarity: "negative",
			Reason:   "Disruption",
			Score:    -3,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  stringPtr("10C/En1"),
			TeacherName: stringPtr("Miss S Prile"),
			RoomName:    stringPtr("E16"),

			Note: stringPtr("John was being extremely disruptive today by shouting out."),

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
		{
			Id:       6,
			Type:     "detention",
			Polarity: "negative",
			Reason:   "Lateness",
			Score:    -2,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  nil,
			},

			LessonName:  nil,
			TeacherName: stringPtr("Mr T Pott"),
			RoomName:    nil,

			Note: stringPtr("John was 3.145 seconds late to school today."),

			DetentionDate:     stringPtr(now.Add(time.Hour * 48).Format("2006-01-02")),
			DetentionTime:     stringPtr("10:20"),
			DetentionLocation: stringPtr("Year Leaders Office"),
			DetentionType:     stringPtr("Lunchtime DT"),
		},
		{
			Id:       7,
			Type:     "event",
			Polarity: "blank",
			Reason:   "Late to school",
			Score:    0,

			Timestamp: now.Add(-time.Hour * 2).Format("2006-01-02 15:04:05"),

			Style: ActivityStyle{
				BorderColour: nil,
				CustomClass:  stringPtr("colour-purple"),
			},

			LessonName:  nil,
			TeacherName: stringPtr("Mr T Pott"),
			RoomName:    nil,

			Note: nil,

			DetentionDate:     nil,
			DetentionTime:     nil,
			DetentionLocation: nil,
			DetentionType:     nil,
		},
	}
}
