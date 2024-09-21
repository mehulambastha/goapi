package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mehulambastha/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO Api Service...")
	fmt.Println(`
     _/_/_/    _/_/          _/_/    _/_/_/    _/_/_/
  _/        _/    _/      _/    _/  _/    _/    _/
 _/  _/_/  _/    _/      _/_/_/_/  _/_/_/      _/
_/    _/  _/    _/      _/    _/  _/          _/
 _/_/_/    _/_/        _/    _/  _/        _/_/_/
    
  `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
