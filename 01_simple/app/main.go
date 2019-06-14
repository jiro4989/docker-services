package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JSONData struct {
	Data string `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	data := JSONData{
		Data: "Index",
	}
	json.NewEncoder(w).Encode(data)
}

func Param(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	data := JSONData{
		Data: "Hello World. Param = " + param.ByName("param"),
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/:param", Param)

	log.Fatal(http.ListenAndServe(":8888", router))
}
