package gotodoit

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"time"

	"github.com/achiku/gotodoit/iapi"
	"github.com/achiku/gotodoit/model"
	"github.com/achiku/gotodoit/service"
	"github.com/rs/xlog"
)

var tokenRegex = regexp.MustCompile(`^(?i)bearer (\w+)$`)

func jsonResponseMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func recoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// AuthModel auth info
type AuthModel struct {
	User  *model.TodoUser
	Token string
}

func getAuthData(ctx context.Context) *AuthModel {
	auth, ok := ctx.Value(ctxKeyAuth).(AuthModel)
	if !ok {
		return &AuthModel{
			Token: "",
			User: &model.TodoUser{
				Email:    "anonymous",
				Status:   "inactive",
				UUID:     "noid",
				Username: "anonymous",
			},
		}
	}
	return &auth
}

func apiAuthMiddleware(app *App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := xlog.FromRequest(r)
			token := r.Header.Get("Authorization")
			matches := tokenRegex.FindStringSubmatch(token)
			encoder := json.NewEncoder(w)
			if len(matches) != 2 {
				e := iapi.Error{
					Code:        iapi.Unauthorized,
					Description: "token is not authorized"}
				w.WriteHeader(http.StatusUnauthorized)
				encoder.Encode(e)
				return
			}
			tkn := matches[1]
			log.Println(tkn)
			u, found, err := service.GetUserByAccessToken(app.DB, tkn)
			if err != nil {
				logger.Error(err)
				e := iapi.Error{
					Code:        iapi.InternalServerError,
					Description: "something went wrong"}
				w.WriteHeader(http.StatusInternalServerError)
				encoder.Encode(e)
				return
			}
			if !found {
				e := iapi.Error{Code: iapi.Unauthorized, Description: "token is not authorized"}
				w.WriteHeader(http.StatusUnauthorized)
				encoder.Encode(e)
				return
			}
			auth := AuthModel{
				Token: tkn,
				User:  u,
			}
			ctx := context.WithValue(r.Context(), ctxKeyAuth, auth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func accessLoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger := xlog.FromRequest(r)
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		logger.Infof(
			"access-log",
			xlog.F{
				"req_time":      t2.Sub(t1),
				"req_time_nsec": t2.Sub(t1).Nanoseconds(),
			},
		)
	}
	return http.HandlerFunc(fn)
}
