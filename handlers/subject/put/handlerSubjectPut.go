package post

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	httperror "github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/services"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SubjectService interface {
	Update(subject *model.Subject) (*model.Subject, httperror.ApiError)
}

type HandlerSubjectPut struct {
	subjectService services.SubjectService
}

func NewHandlerSubjectPut(subjectService services.SubjectService) HandlerSubjectPut {
	return HandlerSubjectPut{
		subjectService: subjectService,
	}
}

func (h HandlerSubjectPut) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httperror.RenderInternalServerError(w, err)
		return
	}

	var subject model.Subject
	err = json.Unmarshal(body, &subject)
	if err != nil {
		httperror.RenderInternalServerError(w, err)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	subject.ID = id

	response, apiErr := h.subjectService.Update(&subject)
	if apiErr != nil {
		httperror.RenderFromApiError(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	render.JSON(w, r, response)
}
