package main

import "fmt"

type player struct {
	name    string
	age     int
	ranking int
	country string
}

func main() {
	rankList := make(map[string]player)
	p := player{"Federer", 37, 3, "Switzerland"}
	rankList[p.name] = p

	p = player{"Djokovic", 31, 1, "Serbia"}
	rankList[p.name] = p

	fmt.Println(rankList)
}
