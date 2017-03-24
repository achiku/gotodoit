package gotodoit

import (
	"log"
	"net/http"
	"os"

	"github.com/achiku/mux"
	"github.com/justinas/alice"
	"github.com/rs/xlog"
)

type ctxKeyType int

const (
	ctxKeyAuth ctxKeyType = iota
)

// RunServer run server
func RunServer(cfgPath string) {
	app, err := NewApp(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	host, _ := os.Hostname()
	logConf := xlog.Config{
		Fields: xlog.F{
			"role": "gotogoit",
			"host": host,
		},
		Output: xlog.NewOutputChannel(xlog.NewConsoleOutput()),
	}

	baseChain := alice.New(
		xlog.NewHandler(logConf),
		xlog.MethodHandler("method"),
		xlog.URLHandler("url"),
		xlog.UserAgentHandler("user_agent"),
		xlog.RefererHandler("referer"),
		xlog.RequestIDHandler("req_id", "Request-Id"),
		accessLoggingMiddleware,
	)
	jsonChain := baseChain.Append(
		jsonResponseMiddleware,
	)
	apiChain := jsonChain.Append(
		apiAuthMiddleware(app),
	)

	router := mux.NewRouter()
	r := router.PathPrefix("/v1").Subrouter()

	// healthcheck
	r.Methods("GET").Path("/healthcheck").Handler(
		jsonChain.Then(AppHandler{h: app.Healthcheck}))
	r.Methods("GET").Path("/healthcheck/nolog").Handler(
		AppHandler{h: app.Healthcheck})

	// user
	// r.Methods("POST").Path("/users").Handler(
	// 	apiChain.Then(AppHandler{h: app.CreateUser}))
	// r.Methods("POST").Path("/login").Handler(
	// 	apiChain.Then(AppHandler{h: app.LoginUser}))
	r.Methods("GET").Path("/users/me").Handler(
		apiChain.Then(AppHandler{h: app.GetUserDetail}))

	// todo
	r.Methods("GET").Path("/todos").Handler(
		apiChain.Then(AppHandler{h: app.GetTodos}))
	r.Methods("GET").Path("/todos/{todoID}").Handler(
		apiChain.Then(AppHandler{h: app.GetTodoByID}))

	if err := http.ListenAndServe(":"+app.Config.ServerPort, router); err != nil {
		log.Fatal(err)
	}
}
