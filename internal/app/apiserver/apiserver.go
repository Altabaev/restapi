package apiserver

import (
	"database/sql"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionsKey))

	s := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
