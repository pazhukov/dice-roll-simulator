package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type dice struct {
	Title string `json:"dice"`
	Value int    `json:"value"`
}

type errMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {

	var errMsg errMessage
	errMsg.Code = 200
	errMsg.Msg = "Use /dice/{value} for get random result. For Example /dice/20 - for d20"

}

func getDice(w http.ResponseWriter, r *http.Request) {
	diceType := mux.Vars(r)["diceType"]

	min := 1
	max, err := strconv.Atoi(diceType)
	if err != nil {
		var errMsg errMessage
		errMsg.Code = 500
		errMsg.Msg = err.Error()
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	value := rand.Intn(max-min+1) + min

	var result dice
	result.Title = "d" + diceType
	result.Value = value

	json.NewEncoder(w).Encode(result)

}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Golang REST API service - Dice Roll Simulator")
	fmt.Println("")
	fmt.Println("Roll d20 > http://localhost:11001/dice/20")
	fmt.Println("Roll d6 > http://localhost:11001/dice/6")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/dice/{diceType}", getDice).Methods("GET")
	log.Fatal(http.ListenAndServe(":11001", router))
}
