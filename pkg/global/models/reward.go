package models

import "math"

type Reward struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	Price int `json:"price"`

	StockControl bool `json:"stock_control"`
	Stock        int  `json:"stock"`

	CanPurchase            bool   `json:"can_purchase"`
	UnableToPurchaseReason string `json:"unable_to_purchase_reason"`

	OncePerPupil bool `json:"once_per_pupil"`

	Purchased      bool `json:"purchased"`
	PurchasedCount int  `json:"purchased_count"`

	PriceBalanceDifference int `json:"price_balance_difference"`
}

type GetRewardMeta struct {
	PupilScoreBalance int `json:"pupil_score_balance"`
}

type Purchased struct {
	SinglePurchase bool `json:"single_purchase"`
	OrderId        int  `json:"order_id"`
	Balance        int  `json:"balance"`
}

func NewMockRewards(meta GetRewardMeta) []Reward {
	return []Reward{
		{
			Id:          1,
			Name:        "Sweet Treat",
			Description: "Feeling a bit hungry? Why not treat yourself to a sweet treat! Choose from a selection of chocolates, sweets, and other treats.",
			Photo:       "https://via.placeholder.com/150",

			Price: 250,

			StockControl: true,
			Stock:        100,

			CanPurchase:            meta.PupilScoreBalance >= 250,
			UnableToPurchaseReason: "Insufficient balance",

			OncePerPupil: false,

			Purchased:      false,
			PurchasedCount: 0,

			PriceBalanceDifference: int(math.Max(float64(250-meta.PupilScoreBalance), 0)),
		},
	}
}

func NewMockPurchased() Purchased {
	return Purchased{
		SinglePurchase: true,
		OrderId:        1,
		Balance:        100,
	}
}
