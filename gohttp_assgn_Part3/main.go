package main

import (
	//"fmt"
	//"domain"
	"time"
	dbrepo "dbrepo/userrepo"
	mongoutils "utils"
	"net/http"
	"github.com/gorilla/mux"
	logger "log"
	"os"
	"delivery/restapplication/usercrudhandler"
	handlerlib "delivery/restapplication/packages/httphandlers"
)


func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	
	
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "hema"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	usersvc := dbrepo.NewService(repoaccess)
	
	hndlr := usercrudhandler.NewRestCrudHandler(usersvc)

	pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/ping/", pingHandler)
	h.Handle("/restaurantservice/restaurant/{id}", hndlr)
	h.Handle("/restaurantservice/restaurant/", hndlr)
	logger.Println("Resource Setup Done.")
	logger.Fatal(http.ListenAndServe(":8080", h))

}	
