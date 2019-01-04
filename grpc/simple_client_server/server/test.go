// comparing method vs. function as field

package main

import "fmt"

type player struct {
	name    string
	age     int32
	rank    int32
	country string
	getRank func(*player) int
}

func (p *player) getRanking() int32 {
	return p.rank
}

func getRank(p *player) int32 {
	return p.rank
}

func main() {
	roger := player{
		name:    "Roger Federer",
		age:     37,
		rank:    3,
		country: "switzerland",
	}

	fmt.Println(roger.getRanking())

	fmt.Println(getRank(&roger))
}
