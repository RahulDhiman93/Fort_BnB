package main

import (
	"bookingApp/internal/config"
	"bookingApp/internal/driver"
	"bookingApp/internal/handlers"
	"bookingApp/internal/helpers"
	"bookingApp/internal/models"
	"bookingApp/internal/render"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println("Application listening on ", portNum)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	//Session date types (what I am to put)
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//change to true for Production
	app.InProduction = false

	//Info Log setup
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	//Error Log setup
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//Connect to DB
	log.Println("<<<-- Connecting to DB")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=HotelBookings user=rahuldhiman password=")
	if err != nil {
		log.Fatal("Cannot connect to DB, dying!!!...")
		return nil, err
	}
	log.Println("Connected to DB -->>>")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return db, err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
