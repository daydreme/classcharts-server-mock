package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/models/global"
	"github.com/daydreme/classcharts-server-mock/pkg/models/responses"
	"net/http"
)

func GetAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	data := global.NewMockAnnouncements()
	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
