package main

import "fmt"

type Player struct {
	Country, Racquet, Shoe string
	Age int
}

func main() {
	players := make(map[string]Player)
	players["Roger Federer"] = Player{"Switzerland", "Wilson", "Nike", 37}
	players["Rafael Nadal"] = Player{"Spain", "Babolat", "Nike", 33}

	for k,v := range players {
		fmt.Println(k, v)
	}
}