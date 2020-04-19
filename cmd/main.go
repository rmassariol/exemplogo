//go get -v -u github.com/gorilla/mux
//go get -u github.com/jinzhu/gorm

// crtl shift p  (Ferramentas)

// install go  tools
// update go tools

// go mod init nome_projeto

// go run .\main.g

// go install .

package main

import (
	"exemplogo/pkg"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "teste")
}

// func HelloWord(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "chamou o hello word")
// }

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//clientes
	myRouter.HandleFunc("/", hello).Methods("GET")
	myRouter.HandleFunc("/hello", pkg.HelloWord).Methods("GET")

	//fornecedor

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("SERVIDOR VIGILANTE ONLINE")
	fmt.Println("SIMPLESTI.COM.BR")
	// Handle Subsequent requests
	handleRequests()
}
