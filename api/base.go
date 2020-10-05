package api

import (
	"fmt"
	"go-line/config"
	"go-line/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	router *mux.Router
	db     *gorm.DB
}

func (a *App) InitAndServe(conf *config.Configuration) {
	//Open database
	a.db = database.InitMySQL(&conf.Database)

	a.router = mux.NewRouter()

	a.SetRoute()

	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	fmt.Println("API Listening to", addr)
	log.Fatal(http.ListenAndServe(addr, a.router))
}
