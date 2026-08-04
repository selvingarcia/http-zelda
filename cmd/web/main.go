package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/selvingarcia/http-zelda/internal/env"
)

//Structs

type application struct{
	logger *slog.Logger
}

type config struct{
	addr string
	apiURL string
}

func main(){

	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	// struct config 
	cfg := config{
		addr: env.GetString("ADRR", ":8080"),
	}

	//struct application

	app:= application{
		logger: logger,
	}

	//logger starting information on server and db status

	logger.Info("Starting the server on ", "addr", cfg.addr)

	//server configuration

	s := &http.Server{
		Addr: cfg.addr,
		Handler: app.routes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	err:= s.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}

// eliminate static path 
func neuter(file http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/"){
			http.NotFound(w,r)
			return
		}

		file.ServeHTTP(w,r)
	})
}