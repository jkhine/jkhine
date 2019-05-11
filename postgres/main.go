package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "joek"
	password = "joek"
	dbname   = "postgres"
)

type Mroute struct {
	Vrf string
	Grp string
	Src string
	Iif string
	Oif string
	Id int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, openErr := sql.Open("postgres", psqlInfo)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer db.Close()

	dbPing(db)
	dbQuerySingleRow(db)
	dbQueryMultiRows(db)
}

func dbQueryMultiRows(db *sql.DB) {
	sqlStatement := "select * from mroutes"
	rows, queryErr := db.Query(sqlStatement)
	if queryErr != nil {
		log.Fatal(queryErr)
	}

	for rows.Next() {
		var mroute Mroute
		scanErr := rows.Scan(&mroute.Vrf, &mroute.Grp, &mroute.Src, &mroute.Iif, &mroute.Oif, &mroute.Id)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		fmt.Println(mroute)
	}
	rowErr := rows.Err()
	if rowErr != nil {
		log.Fatal(rowErr)
	}
}

func dbQuerySingleRow(db *sql.DB) {
	sqlStatement := "select * from mroutes"
	row := db.QueryRow(sqlStatement)

	var mroute Mroute
	scanErr := row.Scan(&mroute.Vrf, &mroute.Grp, &mroute.Src, &mroute.Iif, &mroute.Oif, &mroute.Id)
	switch scanErr {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(mroute)
	default:
		log.Fatal(scanErr)
	}
}

func dbPing(db *sql.DB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}