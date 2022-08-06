package get

import (
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	httperror "github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/services"
	"net/http"
)

type SubjectService interface {
	ListAll() ([]model.Subject, httperror.ApiError)
	GetByID(id string, loadParticipants bool) (model.Subject, httperror.ApiError)
}

type HandlerSubjectGet struct {
	subjectService services.SubjectService
}

func NewHandlerSubjectGet(subjectService services.SubjectService) HandlerSubjectGet {
	return HandlerSubjectGet{
		subjectService: subjectService,
	}
}

func (h HandlerSubjectGet) ListAll(w http.ResponseWriter, r *http.Request) {
	subjects, err := h.subjectService.ListAll()
	if err != nil {
		httperror.RenderFromApiError(w, err)
		return
	}
	render.JSON(w, r, subjects)
}

func (h HandlerSubjectGet) Get(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	loadParticipant := r.URL.Query().Get("load_participant") == "true"

	subjects, err := h.subjectService.Get(idParam, loadParticipant)
	if err != nil {
		httperror.RenderFromApiError(w, err)
		return
	}
	render.JSON(w, r, subjects)
}