package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global/models"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"net/http"
)

func GetAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	data := models.NewMockAnnouncements()
	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
