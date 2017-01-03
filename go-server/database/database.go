package database

import (
	"bytes"
	"database/sql"
	"github.com/xingcuntian/go_test/go-server/helpers"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

type DataBase struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

var dbinfo DataBase

// get database info from exported environment variables
func getDbInfo() {//"mysql","xingcuntian:xingcuntian@2016@tcp(192.168.8.70:3306)/gotest?charset=utf8"
	dbinfo.Name = "gotest"//os.Getenv("DB_NAME")
	dbinfo.Host = "192.168.8.70"//os.Getenv("DB_HOST")
	dbinfo.Port = "3306"//os.Getenv("DB_PORT")
	dbinfo.User = "xingcuntian"//os.Getenv("DB_USER")
	dbinfo.Password = "xingcuntian@2016"//os.Getenv("DB_PASSWORD")
}

// build the db connection as a string
func buildConnection(c DataBase) string {
	var buffer bytes.Buffer
	buffer.WriteString(c.User)
	buffer.WriteString(":")
	buffer.WriteString(c.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(c.Host)
	buffer.WriteString(":")
	buffer.WriteString(c.Port)
	buffer.WriteString(")/")
	buffer.WriteString(c.Name)
	return buffer.String()
}

// connect to database
func DbConnect() *sql.DB {

	getDbInfo()
	dbConnection := buildConnection(dbinfo)
	db, err := sql.Open("mysql", dbConnection)
	helpers.CheckErr(err)

	return db
}
