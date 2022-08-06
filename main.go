package main

import (
	"github.com/gorilla/mux"
	"github.com/mercadolibre/dr-gorm/database"
	get2 "github.com/mercadolibre/dr-gorm/handlers/participant/get"
	delete "github.com/mercadolibre/dr-gorm/handlers/subject/delete"
	"github.com/mercadolibre/dr-gorm/handlers/subject/get"
	"github.com/mercadolibre/dr-gorm/handlers/subject/post"
	put "github.com/mercadolibre/dr-gorm/handlers/subject/put"
	"github.com/mercadolibre/dr-gorm/repository"
	"github.com/mercadolibre/dr-gorm/services"
	"log"
	"net/http"
)

func main() {
	databaseGorm := database.NewDataBaseConnection(NewDataBase())
	subjectRepository := repository.NewSubjectRepository(databaseGorm)
	subjectService := services.NewSubjectService(subjectRepository)

	participantRepository := repository.NewParticipantRepository(databaseGorm)
	participantService := services.NewParticipantService(participantRepository)

	handlerSubjectGet := get.NewHandlerSubjectGet(subjectService)
	handlerSubjectPost := post.NewHandlerSubjectPost(subjectService)
	handlerSubjectPut := put.NewHandlerSubjectPut(subjectService)
	handlerSubjectDelete := delete.NewHandlerSubjectDelete(subjectService)

	handlerParticipantGet := get2.NewHandlerParticipantGet(participantService)


	router := mux.NewRouter()
	router.HandleFunc("/participants", handlerParticipantGet.ListAll).Methods(http.MethodGet)

	router.HandleFunc("/subjects", handlerSubjectGet.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/subjects/{id}", handlerSubjectGet.Get).Methods(http.MethodGet)
	router.HandleFunc("/subjects", handlerSubjectPost.Save).Methods(http.MethodPost)
	router.HandleFunc("/subjects/{id}", handlerSubjectPut.Update).Methods(http.MethodPut)
	router.HandleFunc("/subjects/{id}", handlerSubjectDelete.Delete).Methods(http.MethodDelete)

	log.Println("API DR-GORM is running!")

	http.ListenAndServe(":9090", router)
}

func NewDataBase() database.Database {
	return database.Database{
		Username: "postgres",
		Password: "postgres",
		Name:     "dr-gorm",
	}
}
