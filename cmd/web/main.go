package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct{
	logger *slog.Logger
}

func main(){
	addr := flag.String("addr", ":8080", "Server address")
	flag.Parse()

	fs := http.FileServer(http.Dir("./ui/static"))


	logger:= slog.New(slog.NewJSONHandler(os.Stdout,nil))

	app:= application{
		logger: logger,
	}


	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static",fs))
	mux.HandleFunc("GET /", app.home)


	logger.Info("Starting the server on ", "addr", *addr)

	s := &http.Server{
		Addr: *addr,
		Handler: mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	err:= s.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}