package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	src, err := os.Open("../table.csv")
	if err != nil {
		log.Fatalln("Couldn't open file", err.Error())
	}
	defer src.Close()

	rdr := csv.NewReader(src)

	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("couldn't marshall", err.Error())
	}
	os.Stdout.Write(b)
}
