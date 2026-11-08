package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/tyler-katsha/Golang-programs/internal/handlers"
)

func main(){
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	
	 ____  _  _  __    ____  ____ 
(_  _)( \/ )(  )  ( ___)(  _ \
  )(   \  /  )(__  )__)  )   /
 (__)  (__) (____)(____)(_)\_)

 
	`)

	err := http.ListenAndServe("localhost:8000",r)

	if err != nil {
		log.Error(err)
	}
}