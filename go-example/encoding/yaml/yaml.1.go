package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Pwd struct {
	Password int `yaml:"PASSWORD"`
}

type Env struct {
	GinMode	string  `yaml:"GIN_MODE"`
	Authors []string `yaml:"AUTHORS"`
	Age	    int		`yaml:"AGE"`
	Dev		Pwd    `yaml:"DEV"`
	Test	Pwd		`yaml:"TEST"`
}

func main(){
	content, _ := ioutil.ReadFile("env.yaml")
	env := Env{}
	err := yaml.Unmarshal(content,&env)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(env.Age)
	fmt.Println(env.Authors[0])
	fmt.Println(err,env)
}