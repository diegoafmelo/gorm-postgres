package post

import (
	"encoding/json"
	"github.com/go-chi/render"
	httperror "github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/services"
	"io/ioutil"
	"net/http"
)

type SubjectService interface {
	Save(subject *model.Subject) (*model.Subject, httperror.ApiError)
}

type HandlerSubjectPost struct {
	subjectService services.SubjectService
}

func NewHandlerSubjectPost(subjectService services.SubjectService) HandlerSubjectPost {
	return HandlerSubjectPost{
		subjectService: subjectService,
	}
}

func (h HandlerSubjectPost) Save(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httperror.RenderInternalServerError(w, err)
		return
	}

	var request model.Subject
	err = json.Unmarshal(body, &request)
	if err != nil {
		httperror.RenderInternalServerError(w, err)
		return
	}

	response, apiErr := h.subjectService.Save(&request)
	if apiErr != nil {
		httperror.RenderFromApiError(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	render.JSON(w, r, response)
}
