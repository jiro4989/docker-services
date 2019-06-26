package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	driverName     string
	dataSourceName string
}

type Money struct {
	ID, Data int
}

func newDB(driverName, dataSourceName string) DB {
	return DB{driverName: driverName, dataSourceName: dataSourceName}
}

func (d DB) insertMoney(money int) (int64, error) {
	db, err := sql.Open(d.driverName, d.dataSourceName)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO money (data) VALUES (?)", money)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (d DB) deleteMoney(moneyID int) (int64, error) {
	db, err := sql.Open(d.driverName, d.dataSourceName)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM money WHERE id = ?", moneyID)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (d DB) fetchMoney() ([]Money, error) {
	// db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
	db, err := sql.Open(d.driverName, d.dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, data FROM money") //
	if err != nil {
		return nil, err
	}

	var money []Money
	for rows.Next() {
		var id int
		var data int
		if err := rows.Scan(&id, &data); err != nil {
			return nil, err
		}
		log.Println("id:", id, ",data:", data)
		m := Money{ID: id, Data: data}
		money = append(money, m)
	}

	return money, nil
}
