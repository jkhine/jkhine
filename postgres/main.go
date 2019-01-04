package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type player struct {
	firstName string
	lastName string
	country string
	age int
	rank int
	majors int
	masters int
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=Xplore1 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from players")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		player := player{}
		err = rows.Scan(&player.firstName, &player.lastName, &player.country, &player.age, &player.rank, &player.majors, &player.masters)
		if err != nil {
			panic(err)
		}
		fmt.Println(player)
	}
}