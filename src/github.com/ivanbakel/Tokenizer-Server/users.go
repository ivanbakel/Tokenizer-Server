package main

import (
	"net/http"
	"encoding/json"
	"github.com/ivanbakel/Tokenizer-Server/models"
	"github.com/vattle/sqlboiler/boil"
	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users,err = models.Users(boil.GetDB()).All()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var encoder *json.Encoder = json.NewEncoder(w)

	for _,user := range users {
		encoder.Encode(user)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var userID = mux.Vars(r)["uid"]

	var user,err = models.FindUser(boil.GetDB(), userID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var encoder *json.Encoder = json.NewEncoder(w)

	encoder.Encode(user)
}

func getUserTokens(w http.ResponseWriter, r *http.Request) {
	var userID = mux.Vars(r)["uid"]

	var user,err = models.FindUser(boil.GetDB(), userID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user.L.LoadUserTokens(boil.GetDB(), true, user)

	var encoder *json.Encoder = json.NewEncoder(w)

	for _,token := range user.R.UserTokens {
		encoder.Encode(token)
	}
}
