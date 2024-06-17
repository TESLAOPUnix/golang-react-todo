package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/teslaopunix/golang-react-todo/router"
)

func main(){
	r:= router.Router()
	fmt.Println("starting in port 9000")
	log.Fatal(http.ListenAndServe(":9000",r))
}