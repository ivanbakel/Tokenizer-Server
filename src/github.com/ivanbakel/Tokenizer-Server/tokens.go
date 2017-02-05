package main

import (
	"net/http"
	"encoding/json"
	"github.com/ivanbakel/Tokenizer-Server/models"
	"github.com/vattle/sqlboiler/boil"
)


func getTokens(w http.ResponseWriter, r *http.Request) {
	var encoder *json.Encoder = json.NewEncoder(w)

	var tokens,err = models.Tokens(boil.GetDB()).All()

	if err == nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _,token := range tokens {
		encoder.Encode(token)
	}

}

func getToken(w http.ResponseWriter, r *http.Request) {

}

func giveGroupTokens(w http.ResponseWriter, r *http.Request) {

}

func giveUserTokens(w http.ResponseWriter, r *http.Request) {

}

func createToken(w http.ResponseWriter, r *http.Request) {

}

func receiveTokens(w http.ResponseWriter, r *http.Request) {
	
}

func spendTokens(w http.ResponseWriter, r *http.Request) {

}
