package main

import (
    "fmt"
    "encoding/json"
)

var jsonStr = []byte(
`{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        }
    ]
}`)

type Person struct {
    Name string
    Age  int
}

type Place struct {
    City    string
    Country string
}

func main(){
    personsA, placesA := solutionA(jsonStr)
	fmt.Printf("%d %d\n", len(personsA), len(placesA))

    fmt.Println(personsA)
    fmt.Println(placesA)
}

//方法A: map and type assert

func solutionA(jsonStr []byte) ([]Person,[]Place){
    persons := []Person{}
    places := []Place{}
    var data map[string][]map[string]interface{}
    fmt.Println(data)
    err :=json.Unmarshal(jsonStr,&data)
    if err != nil {
        fmt.Println(err)
        return persons,places
    }
    fmt.Println(data)
    for i := range data["things"] {
        item := data["things"][i]
        if item["name"] != nil {
            persons = addPerson(persons,item)
        }else{
            places = addPlace(places,item)
        }
    }
    return persons, places
}

func addPerson(persons []Person, item map[string]interface{}) []Person{
    name, _ := item["name"].(string)
    age, _ := item["age"].(int)
    person := Person{name,age}
    persons = append(persons,person)
    return persons
}

func addPlace(places []Place, item map[string]interface{}) []Place{
    city, _ := item["city"].(string)
    country, _ := item["country"].(string)
    place := Place{city,country}
    places = append(places,place)
    return places
}