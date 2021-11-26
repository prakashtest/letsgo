package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	logger, _ := zap.NewProduction()
	logger.Info("succefully performed http request ")

}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8084", router))
}
