package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <title>Chat</title>
        </head>
        <body>
          Let's start chatting!
        </body>
      </html>
    `))
	})

	fmt.Println("Before if")
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
		fmt.Println("Error block")
	} else {
		fmt.Println("Server Listening")
		fmt.Println("Else block")
	}
}