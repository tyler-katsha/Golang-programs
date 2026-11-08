package api

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/tyler-katsha/Golang-programs/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.setReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	
	 ____  _  _  __    ____  ____ 
(_  _)( \/ )(  )  ( ___)(  _ \
  )(   \  /  )(__  )__)  )   /
 (__)  (__) (____)(____)(_)\_)

 
	`)

}