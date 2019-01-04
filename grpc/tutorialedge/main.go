package main

import (
	fmt "fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	joe := &Person{
		Name: "Joe",
		Age:  47,
	}

	data, _ := proto.Marshal(joe)
	fmt.Println(data)

	newJoe := &Person{}
	err := proto.Unmarshal(data, newJoe)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	fmt.Println(newJoe.GetName())
	fmt.Println(newJoe.GetAge())
}
