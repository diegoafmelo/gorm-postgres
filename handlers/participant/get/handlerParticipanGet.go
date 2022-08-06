package get

import (
	"github.com/go-chi/render"
	httperror "github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/services"
	"net/http"
)

type ParticipantService interface {
	ListAll(isJoin bool, isActive bool) ([]model.Participant, httperror.ApiError)
}

type HandlerParticipantGet struct {
	participantService services.ParticipantService
}

func NewHandlerParticipantGet(participantService services.ParticipantService) HandlerParticipantGet {
	return HandlerParticipantGet{
		participantService: participantService,
	}
}

func (h HandlerParticipantGet) ListAll(w http.ResponseWriter, r *http.Request) {
	isJoin := r.URL.Query().Get("is_join") == "true"
	isActive := r.URL.Query().Get("is_active") == "true"

	subjects, err := h.participantService.ListAll(isJoin, isActive)
	if err != nil {
		httperror.RenderFromApiError(w, err)
		return
	}
	render.JSON(w, r, subjects)
}