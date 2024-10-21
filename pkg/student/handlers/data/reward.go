package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global/models"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/util"
	"net/http"
)

func GetRewardHandler(w http.ResponseWriter, r *http.Request) {
	score, err := util.RandomInt(150, 290)
	if err != nil {
		panic(err)
	}

	meta := models.GetRewardMeta{
		PupilScoreBalance: score,
	}

	data := models.NewMockRewards(meta)
	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	data := models.NewMockPurchased()
	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
