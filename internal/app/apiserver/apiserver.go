package apiserver

import (
	"github.com/Altabaev/Go-Rest-Api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}

	a.configureRouter()

	if err := a.configureStore(); err != nil {
		return err
	}

	a.logger.Info("Starting API server")

	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

func (a *APIServer) configureRouter() {
	a.router.HandleFunc("/hello", a.HandleHello())
}

func (a *APIServer) HandleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello")
		if err != nil {
			return
		}
	}
}

func (a *APIServer) configureStore() error {
	st := store.New(a.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	a.store = st

	return nil
}
