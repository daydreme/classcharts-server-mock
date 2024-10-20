package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/models/global"
	"github.com/daydreme/classcharts-server-mock/pkg/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/util"
	"net/http"
)

func GetRewardHandler(w http.ResponseWriter, r *http.Request) {
	score, err := util.RandomInt(150, 290)
	if err != nil {
		panic(err)
	}

	meta := global.GetRewardMeta{
		PupilScoreBalance: score,
	}

	data := global.NewMockRewards(meta)
	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

func GetPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	data := global.NewMockPurchased()
	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
