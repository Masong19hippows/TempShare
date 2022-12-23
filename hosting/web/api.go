package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	testSes "github.com/masong19hippows/TempShare/hosting/session"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	log.Println("Starting the application")

	r := mux.NewRouter()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	r.HandleFunc("/api/user/login", userLogin).Methods("POST")
	r.HandleFunc("/api/user/signup", userSignup).Methods("POST")
	http.ListenAndServe("0.0.0.0:6000", r)
}

func get(response http.ResponseWriter, request *http.Request) {
	ses, err := testSes.Create("garten323@gmail.com", 15*time.Minute)
	if err != nil {
		panic(err)
	}
	jsonResponse, jsonError := json.Marshal(ses)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func create(response http.ResponseWriter, request *http.Request) {
	ses, err := testSes.Create("garten323@gmail.com", 15*time.Minute)
	if err != nil {
		panic(err)
	}
	jsonResponse, jsonError := json.Marshal(ses)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func Test() {
	fmt.Println("test")
}
