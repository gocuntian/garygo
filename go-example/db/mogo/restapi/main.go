package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

type Note struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
}

type NoteResource struct {
	Note Note `json:"note"`
}

type NotesResource struct {
	Notes []Note `json:"notes"`
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var noteResource NoteResource
	err := json.NewDecoder(r.Body).Decode(&noteResource)
	if err != nil {
		log.Fatalln(err)
	}
	note := noteResource.Note
	note_id := bson.NewObjectId()
	note.Id = note_id
	if err = collection.Insert(&note); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Inserted new Note %s with name %s", note.Id, note.Name)
	}
	data, err := json.Marshal(NoteResource{Note: note})
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	log.Println("start server")
	r := mux.NewRouter()
	r.HandleFunc("/api/notes", CreateNoteHandler).Methods("POST")
	http.Handle("/api/", r)
	http.Handle("/", http.FileServer(http.Dir(".")))
	//mgo
	log.Println("Starting mongo db session")
	var err error
	session, err = mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("xingcuntian").C("notes")
	fmt.Println("listen:9090....")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}
