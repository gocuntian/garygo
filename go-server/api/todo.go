package api

import (
	"encoding/json"
	"fmt"
	"github.com/xingcuntian/go_test/go-server/database"
	"github.com/xingcuntian/go_test/go-server/helpers"
	"github.com/xingcuntian/go_test/go-server/models"
	"github.com/xingcuntian/go_test/httprouter"
	"net/http"
)

type TodoApi struct{}

func NewTodoApi() *TodoApi {
	return &TodoApi{}
}

// base endpoint
func (ta TodoApi) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Welcome from Go-Server!")
}

// base todo endpoint
// returns array of todos
func (ta TodoApi) GetTodos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	// connect to db
	db := database.DbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM t_todo")

	todo := make([]models.Todo, 0)

	helpers.CheckErr(err)

	for rows.Next() {

		var id int
		var name string
		var description string
		var completed bool
		var created []byte
		var updated []byte
		var due []byte
		var date_completed []byte

		if err := rows.Scan(&id, &name, &description, &completed, &created, &updated, &due, &date_completed); err != nil {
			panic(err)
		}
		todo = append(todo, models.Todo{
			Id:            id,
			Name:          name,
			Description:   description,
			Completed:     completed,
			Created:       created,
			Updated:       updated,
			Due:           due,
			DateCompleted: date_completed,
		})
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	js, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// Get a todo by id
func (ta TodoApi) GetTodoById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var todo models.Todo

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	err = db.QueryRow("SELECT * FROM t_todo where id = ?", p.ByName("id")).Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Completed, &todo.Created, &todo.Updated, &todo.Due, &todo.DateCompleted)

	if err == nil {
		// 1 row
		js, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		// error
		fmt.Printf("Error is %s\n", err)
		helpers.CheckErr(err)
	}

}

// create a todo
func (ta TodoApi) CreateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// stub a new todo
	todo := models.Todo{}

	// populate the todo with the data from post req
	json.NewDecoder(r.Body).Decode(&todo)

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	// Insert data to database
	stmt, err := db.Prepare("INSERT INTO t_todo(name, description, due) VALUES ($1, $2, $3)")
	helpers.CheckErr(err)

	_, err = stmt.Exec(todo.Name, todo.Description, todo.Due)
	helpers.CheckErr(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

}

// delete a todo
func (ta TodoApi) DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	// Delete todo from database
	stmt, err := db.Prepare("DELETE FROM t_todo where name=?")
	helpers.CheckErr(err)

	_, err = stmt.Exec(p.ByName("name"))
	helpers.CheckErr(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
