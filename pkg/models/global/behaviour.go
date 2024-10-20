package global

import "time"

type BehaviourTimelineItem struct {
	End      string `json:"end"`
	Name     string `json:"name"`
	Negative int    `json:"negative"`
	Positive int    `json:"positive"`
	Start    string `json:"start"`
}

type Behaviour struct {
	NegativeReasons    map[string]int `json:"negative_reasons"`
	OtherNegative      []string       `json:"other_negative"`
	OtherNegativeCount map[string]int `json:"other_negative_count"`

	PositiveReasons    map[string]int `json:"positive_reasons"`
	OtherPositive      []string       `json:"other_positive"`
	OtherPositiveCount map[string]int `json:"other_positive_count"`

	Timeline []BehaviourTimelineItem `json:"timeline"`
}

const dateFormat = "2006-01-02"

func NewMockBehaviour() Behaviour {
	now := time.Now()

	return Behaviour{
		NegativeReasons: map[string]int{
			"Disruption":     -3,
			"Lateness":       -2,
			"Other negative": -5,
		},
		OtherNegative: []string{
			"Unfinished assignment",
			"Missing materials",
		},
		OtherNegativeCount: map[string]int{
			"Unfinished assignment": -3,
			"Missing materials":     -1,
		},

		PositiveReasons: map[string]int{
			"Homework completion":  5,
			"Outstanding project":  2,
			"Class contribution":   10,
			"Team collaboration":   4,
			"Exceptional behavior": 7,
			"Other positive":       3,
		},
		OtherPositive: []string{
			"Perfect attendance",
			"Excellent participation",
		},
		OtherPositiveCount: map[string]int{
			"Perfect attendance":      4,
			"Excellent participation": 3,
		},

		Timeline: []BehaviourTimelineItem{
			{
				Positive: 10,
				Negative: -1,
				Name:     now.AddDate(0, -7, 0).Format("1/2006"), // 7 months ago
				Start:    now.AddDate(0, -7, 0).Format(dateFormat),
				End:      now.AddDate(0, -7, 30).Format(dateFormat),
			},
			{
				Positive: 8,
				Negative: -2,
				Name:     now.AddDate(0, -6, 0).Format("1/2006"), // 6 months ago
				Start:    now.AddDate(0, -6, 0).Format(dateFormat),
				End:      now.AddDate(0, -6, 30).Format(dateFormat),
			},
			{
				Positive: 12,
				Negative: 0,
				Name:     now.AddDate(0, -5, 0).Format("1/2006"), // 5 months ago
				Start:    now.AddDate(0, -5, 0).Format(dateFormat),
				End:      now.AddDate(0, -5, 30).Format(dateFormat),
			},
			{
				Positive: 15,
				Negative: -3,
				Name:     now.AddDate(0, -4, 0).Format("1/2006"), // 4 months ago
				Start:    now.AddDate(0, -4, 0).Format(dateFormat),
				End:      now.AddDate(0, -4, 30).Format(dateFormat),
			},
			{
				Positive: 20,
				Negative: 0,
				Name:     now.AddDate(0, -3, 0).Format("1/2006"), // 3 months ago
				Start:    now.AddDate(0, -3, 0).Format(dateFormat),
				End:      now.AddDate(0, -3, 30).Format(dateFormat),
			},
			{
				Positive: 25,
				Negative: -2,
				Name:     now.AddDate(0, -2, 0).Format("1/2006"), // 2 months ago
				Start:    now.AddDate(0, -2, 0).Format(dateFormat),
				End:      now.AddDate(0, -2, 30).Format(dateFormat),
			},
			{
				Positive: 30,
				Negative: -1,
				Name:     now.AddDate(0, -1, 0).Format("1/2006"), // 1 month ago
				Start:    now.AddDate(0, -1, 0).Format(dateFormat),
				End:      now.AddDate(0, -1, 30).Format(dateFormat),
			},
			{
				Positive: 35,
				Negative: 0,
				Name:     now.Format("1/2006"), // current month
				Start:    now.Format(dateFormat),
				End:      now.AddDate(0, 0, 30).Format(dateFormat),
			},
		},
	}
}
