package data

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/daydreme/classcharts-server-mock/pkg/util"
)

func GetRewardHandler(w http.ResponseWriter, _ *http.Request) {
	score, err := util.RandomInt(150, 290)
	if err != nil {
		panic(err)
	}

	meta := shared.GetRewardMeta{
		PupilScoreBalance: score,
	}

	data := shared.NewMockRewards(meta)
	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPurchaseHandler(w http.ResponseWriter, _ *http.Request) {
	data := shared.NewMockPurchased()
	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}
