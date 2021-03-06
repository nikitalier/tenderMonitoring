package application

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/nikitalier/tenderMonitoring/config"
	"github.com/nikitalier/tenderMonitoring/pkg/service"
)

//Application struct
type Application struct {
	serv   *http.Server
	svc    *service.Service
	logger zerolog.Logger
}

//Options struct
type Options struct {
	Svc    *service.Service
	Serv   config.ServerOpt
	Logger zerolog.Logger
}

//New - init app
func New(opt *Options) *Application {
	var allowedHeaders = handlers.AllowedHeaders(opt.Serv.AllowedHeaders)

	var exposedHeaders = handlers.ExposedHeaders(opt.Serv.ExposedHeaders)

	var allowedMethods = handlers.AllowedMethods(opt.Serv.AllowedMethods)

	var allowedCredentials = handlers.AllowCredentials()

	var originValidator handlers.OriginValidator = func(withAuth string) bool {
		return true
	}

	var allowedOriginValidator = handlers.AllowedOriginValidator(originValidator)

	app := &Application{
		svc: opt.Svc,
		serv: &http.Server{
			Addr: opt.Serv.Port,
		},
		logger: opt.Logger,
	}

	app.serv.Handler = handlers.CORS(
		allowedHeaders,
		exposedHeaders,
		allowedCredentials,
		allowedMethods,
		allowedOriginValidator,
	)(app.setupRoutes())

	app.logger.Info().Msg("App started on port" + opt.Serv.Port)

	return app
}

//Start app
func (app *Application) Start() {
	app.serv.ListenAndServe()
}

func (app *Application) setupRoutes() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/security/login", app.loginHandler).Methods("POST")
	r.HandleFunc("/users", app.findUserHandler).Methods("GET")
	r.HandleFunc("/users/all", app.getAllUsers).Methods("GET")

	r.HandleFunc("/keywords/all", app.getAllKeyWords).Methods("GET")
	r.HandleFunc("/keywords/remove", app.deleteKeyword).Methods("POST")
	r.HandleFunc("/keywords/add", app.addKeywordHandler).Methods("POST")

	r.HandleFunc("/tenders/all", app.getAllTenders).Methods("GET")
	r.HandleFunc("/tenders", app.getTenderHandler).Methods("GET")

	r.HandleFunc("/favorite", app.getFavorite).Methods("POST")
	r.HandleFunc("/favorite/update", app.updateFavorite).Methods("POST")

	r.HandleFunc("/comments/all", app.getAllComments).Methods("GET")
	r.HandleFunc("/comments/add", app.addNewComment).Methods("POST")

	r.HandleFunc("/tenderstatus/create", app.createTenderStatus).Methods("GET")
	r.HandleFunc("/tenderstatus/update", app.updateTenderStatus).Methods("POST")
	r.HandleFunc("/tenderstatus", app.getTenderStatus).Methods("GET")

	r.HandleFunc("/summary", app.getSummary).Methods("GET")

	return r
}
