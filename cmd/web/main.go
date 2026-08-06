package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/selvingarcia/http-zelda/internal/env"
	"github.com/selvingarcia/http-zelda/internal/models"
)

//Structs

type application struct{
	logger *slog.Logger
	addr string
	db *models.StatusModel
}



func main(){

	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	pool, err := openDB()
	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}
	//struct application

	app:= application{
		logger: logger,
		addr: env.GetString("ADRR", ":8080"),
		db: &models.StatusModel{DB:pool},
	}

	defer pool.Close()
	logger.Info("Database connection succesfull on ", "addr", env.GetString("DB_PORT", ""))

	//logger starting information on server and db status

	logger.Info("Starting the server on ", "addr", app.addr)

	//server configuration

	s := &http.Server{
		Addr: app.addr,
		Handler: app.routes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	err= s.ListenAndServe()

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

func openDB()(*sql.DB, error){
	cfg:= pq.Config{
		Host:           env.GetString("DB_HOST", "localhost"),
		Port:           uint16(env.GetInt("DB_PORT", "5432")),
		User:           env.GetString("DB_USER", "pqgo"),
		Password:       env.GetString("DB_PASSWORD", ""),
		Database:        env.GetString("DB_NAME", "httpzelda"),
		SSLMode:        pq.SSLMode(env.GetString("DB_SSLMODE", "disable")),
		ConnectTimeout: env.GetTime("DB_CONNECT_TIMEOUT", "5s"),

	}

	c , err := pq.NewConnectorConfig(cfg) 
	if err!=nil{
		return nil, err
	}
	
	db:= sql.OpenDB(c)

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}