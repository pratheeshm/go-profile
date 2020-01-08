package main

import (
	_ "github.com/lib/pq"
	"github.com/pkg/profile"
	"github.com/pratheeshm/go-profile/db"
	"log"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath("./data"), profile.MemProfileRate(1)).Stop()
	// dbConn, err := db.SQLConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var student = &db.Student{
	// 	Name: "John",
	// 	Age:  20,
	// }
	// for i := 0; i < 3000; i++ {
	// 	db.RowInsert(dbConn, student)
	// }
	dbConn, err := db.GormConnection()
	if err != nil {
		log.Fatal(err)
	}
	var student = &db.Student{
		Name: "John",
		Age:  20,
	}
	for i := 0; i < 3000; i++ {
		db.GormInsert(dbConn, student)
	}
}
