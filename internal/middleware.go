package middleware

import (
	"errors"
	"go/token"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tyler-katsha/Golang-programs/api"
	"github.com/tyler-katsha/Golang-programs/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.ResponseWriter(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().GET("username")
		var token = r.Header.GET("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if(loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)
	})
}