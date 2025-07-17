// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 2 chanel
// 	ping := make(chan string)
// 	pong := make(chan string)

// 	go func() {
// 		for {
// 			msg := <-ping
// 			fmt.Println("ping")
// 			time.Sleep(1 * time.Second)
// 			pong <- msg
// 		}

// 		for {
// 			msg := <-pong
// 			fmt.Println("pong")
// 			ping <- msg

// 		}

// 	}()
// 	ping <- "hello"

// 	time.Sleep(10 * time.Second)

// 	for {

//		}
//	}
package main

import "net/http"

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /greet", greetUser)
	http.ListenAndServe(":7080", router)
}
func greetUser(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome user!!"))

}
