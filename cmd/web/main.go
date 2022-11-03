package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ryutaro-asada/go-practice/pkg/config"
	"github.com/ryutaro-asada/go-practice/pkg/handlers"
	"github.com/ryutaro-asada/go-practice/pkg/render"
)

const portNumber = ":8080"

// app is config
var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Sesstion = session

	// cache create
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not get chache")
	}

	// give config chache
	app.TemplateCache = tc
	app.UseCache = true

	// create repo that is data pool used by handler
	// give pointer for changed in func
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// give pointer for changed in func
	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("port number is ", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
