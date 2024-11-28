package main

import (
	"database/sql"
	"fmt"
)

type DBconn struct {
	conn *sql.DB
}

var DB *DBconn

func New(host, user, password, dbname string) (*DBconn, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &DBconn{conn: conn}, nil
}

func (db *DBconn) Close() error {
	return db.conn.Close()
}

func (db *DBconn) Begin() (*sql.Tx, error) {
	return db.conn.Begin()
}

func init() {
	db, err := New("localhost", "postgres", "password", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}
