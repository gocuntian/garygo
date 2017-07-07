package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type (
	Category struct {
		Id          bson.ObjectId `bson:"_id,omitempty"`
		Name        string
		Description string
	}
	DataStore struct {
		session *mgo.Session
	}
)

func (d *DataStore) Close() {
	d.session.Close()
}

func (d *DataStore) C(name string) *mgo.Collection {
	return d.session.DB("xingcuntian").C(name)
}

func NewDataStore() *DataStore {
	ds := &DataStore{
		session: session.Copy(),
	}
	return ds
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalln(err)
	}
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C("categories")
	err = c.Insert(&category)
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []Category
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C("categories")
	iter := c.Find(nil).Iter()
	result := Category{}
	for iter.Next(&result) {
		categories = append(categories, result)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(categories)
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	var err error
	session, err = mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		log.Fatalln(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/categories", GetCategories).Methods("GET")
	r.HandleFunc("/api/categories", PostCategory).Methods("POST")
	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}
	fmt.Println("listening ....")
	server.ListenAndServe()

}
