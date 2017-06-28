package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data := []User{
		User{
			FirstName: "xingcuntian",
			LastName:  "gary",
			Email:     "gary@qq.com",
		},
		User{
			FirstName: "xingcuntian2",
			LastName:  "gary2",
			Email:     "gary@qq.com2",
		},
	}
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":9090", r)
}
