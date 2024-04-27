package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/theloosygoose/go-api/api"
	"github.com/theloosygoose/go-api/internal/tools"
)

var UnAutorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {
        var username string = r.URL.Query().Get("username")
        var token = r.Header.Get("Authorization")

        var err error

        if username == "" || token == "" {
            log.Error(UnAutorizedError)
            api.RequestErrorHandler(w, UnAutorizedError)
            return
        }

        var database *tools.DatabaseInterface
        database, err = tools.NewDatabase()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }

        var loginDetails *tools.LoginDetails
        loginDetails = (*database).GetUserLoginDetails(username)

        if (loginDetails == nil || (token != (*loginDetails).AuthToken)){
            log.Error(UnAutorizedError)
            api.RequestErrorHandler(w, UnAutorizedError)
            return
        }

        next.ServeHTTP(w, r)
        
	})
}
