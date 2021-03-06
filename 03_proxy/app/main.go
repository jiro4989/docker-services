package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	driverName     = "mysql"
	dataSourceName = "root:password@tcp(db:3306)/my_db"
)

type InsertResult struct {
	InsertedID int64 `json:"insertedID"`
	Money      int   `json:"money"`
}

type DeleteResult struct {
	DeletedID int64 `json:"deletedID"`
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // <- Added
}

// moneyをDBに追加する
func Insert(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	setHeader(w)

	log.Println("Insert: ", param)
	errData := InsertResult{
		InsertedID: -1,
		Money:      -1,
	}

	moneyStr := param.ByName("money")
	money, err := strconv.Atoi(moneyStr)
	if err != nil {
		json.NewEncoder(w).Encode(errData)
		return
	}

	db := newDB(driverName, dataSourceName)
	result, err := db.insertMoney(money)
	if err != nil {
		json.NewEncoder(w).Encode(errData)
		return
	}
	data := InsertResult{
		InsertedID: result,
		Money:      money,
	}
	json.NewEncoder(w).Encode(data)
}

// idでレコードを削除する
func Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	log.Println("Delete: ", param)
	setHeader(w)
	errData := DeleteResult{
		DeletedID: -1,
	}

	idStr := param.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(errData)
		return
	}

	db := newDB(driverName, dataSourceName)
	result, err := db.deleteMoney(id)
	if err != nil {
		json.NewEncoder(w).Encode(errData)
		return
	}
	data := DeleteResult{
		DeletedID: result,
	}
	json.NewEncoder(w).Encode(data)
}

// moneyの集計結果を返却する
func APIMoney(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	log.Println("Money: ", param)
	setHeader(w)
	db := newDB(driverName, dataSourceName)
	money, err := db.fetchMoney()
	if err != nil {
		var empty []Money
		json.NewEncoder(w).Encode(empty)
	}
	json.NewEncoder(w).Encode(money)
}

func main() {
	router := httprouter.New()
	router.GET("/api/insert/:money", Insert)
	router.GET("/api/delete/:id", Delete)
	router.GET("/api/money", APIMoney)

	log.Fatal(http.ListenAndServe(":8888", router))
}
