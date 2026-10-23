package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lib/pq"
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
	defer pool.Close()

	if err := initSchema(pool, env.GetString("CODES_CSV", "codes.csv")); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	addr := env.GetString("ADDR", ":8080")
	// Web Port 
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	//struct application

	app:= application{
		logger: logger,
		addr: addr,
		db: &models.StatusModel{DB:pool},
	}

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



func openDB()(*sql.DB, error){

	if url := os.Getenv("DATABASE_URL"); url != "" {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(env.GetInt("DB_MAX_OPEN_CONNS", "30"))
	db.SetMaxIdleConns(env.GetInt("DB_MAX_IDLE_CONNS", "30"))
	db.SetConnMaxIdleTime(env.GetTime("DB_MAX_IDLE_TIME", "15m"))

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
	}

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

	db.SetMaxOpenConns(env.GetInt("DB_MAX_OPEN_CONNS", "30"))
	db.SetMaxIdleConns(env.GetInt("DB_MAX_IDLE_CONNS", "30"))
	db.SetConnMaxIdleTime(env.GetTime("DB_MAX_IDLE_TIME", "15m"))

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}