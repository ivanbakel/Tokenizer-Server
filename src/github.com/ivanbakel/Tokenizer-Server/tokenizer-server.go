package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

var dataBase *sql.DB

const databaseUsername = "tokenizerDB"
const databaseName = "tokenizer"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	_, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", databaseUsername, databaseName))

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers)
	r.HandleFunc("/users/{uid}", getUser)
	r.HandleFunc("/users/{uid}/tokens", getUserTokens)
	r.HandleFunc("/tokens", getTokens)
	r.HandleFunc("/tokens/{tid}", getToken)
	r.HandleFunc("/tokens/{tid}/grant-group", giveGroupTokens)
	r.HandleFunc("/tokens/{tid}/grant-user", giveUserTokens)
	r.HandleFunc("/tokens/create", createToken)
	r.HandleFunc("/users/{uid}/tokens/{tid}/receive", receiveTokens)
	r.HandleFunc("/users/{uid}/tokens/{tid}/spend", spendTokens)
	r.HandleFunc("/orgs", getOrgs)
	r.HandleFunc("/orgs/{oid}", getOrg)
	http.ListenAndServe(":8080", nil)
}
