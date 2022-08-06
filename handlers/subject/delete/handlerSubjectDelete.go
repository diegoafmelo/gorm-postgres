package post

import (
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	httperror "github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/services"
	"net/http"
)

type SubjectService interface {
	Delete(id string) httperror.ApiError
}

type HandlerSubjectDelete struct {
	subjectService services.SubjectService
}

func NewHandlerSubjectDelete(subjectService services.SubjectService) HandlerSubjectDelete {
	return HandlerSubjectDelete{
		subjectService: subjectService,
	}
}

func (h HandlerSubjectDelete) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	apiErr := h.subjectService.Delete(idParam)
	if apiErr != nil {
		httperror.RenderFromApiError(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	render.JSON(w, r, nil)
}
