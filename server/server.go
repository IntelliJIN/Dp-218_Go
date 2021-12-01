package server

import (
	"Dp-218_Go/configs"
	"Dp-218_Go/pgdb"
	"Dp-218_Go/pkg/handlers"
	"Dp-218_Go/pkg/repository"
	"Dp-218_Go/pkg/services"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func Run()error{
	cfg := configs.Get()
	pgDB :=  pgdb.Dial(cfg)

	fileRepository := repository.NewFileRepository(pgDB)
	fileService := services.NewFileService(fileRepository)
	fileHandler := handlers.NewFileHandler(fileService)

	scooterRepository := repository.NewScooterRepository(pgDB)
	scooterService := services.NewScooterService(scooterRepository)
	scooterHandler := handlers.NewScooterHandler(scooterService)

	scooterModelRepository := repository.NewScooterModelRepository(pgDB)
	scooterModelService := services.NewScooterModelService(scooterModelRepository)
	scooterModelHandler := handlers.NewScooterModelHandler(scooterModelService)


	r := mux.NewRouter()
	r.HandleFunc("/upload", fileHandler.UploadFile).Methods("POST")
//	r.HandleFunc("/download", fileHandler.DownloadFile).Methods("GET")

	r.HandleFunc("/createScooter", scooterHandler.CreateScooter).Methods("POST")
	r.HandleFunc("/getScooters", scooterHandler.GetScooters).Methods("GET")
	r.HandleFunc("/getScooter/{id}", scooterHandler.GetScooterById).Methods("GET")
	r.HandleFunc("/getScooter/{modelName}", scooterHandler.GetScooterByModelName).Methods("GET")
	r.HandleFunc("/updateSerial/{id}", scooterHandler.UpdateScooterSerial).Methods("PUT")
	r.HandleFunc("/deleteScooter/{id}", scooterHandler.DeleteScooter).Methods("DELETE")
	r.HandleFunc("/getScooterByModelId/{id}", scooterHandler.GetScooterByModelId).Methods("GET")

	r.HandleFunc("/createModel", scooterModelHandler.CreateScooterModel).Methods("POST")
	r.HandleFunc("/getModels", scooterModelHandler.GetScooterModels).Methods("GET")
	r.HandleFunc("/getModel/{id}", scooterModelHandler.GetScooterModelByID).Methods("GET")
	r.HandleFunc("/editModel/{id}", scooterModelHandler.EditScooterModel).Methods("PUT")
	r.HandleFunc("/deleteModel/{id}", scooterModelHandler.DeleteScooterModel).Methods("DELETE")

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", r)
	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down")
	os.Exit(0)
	return nil

}
