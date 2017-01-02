package main

import (
    "fmt"
    "sort"
)

type earthMass float64
type au float64

type Planet struct{
    name     string
    mass     earthMass
    distance au
}

type By func(p1,p2 *Planet)bool
func (by By)Sort(planets []Planet){
    ps:=&planetSorter{
        planets:planets,
        by:by,
    }
    sort.Sort(ps)
}

type planetSorter struct{
    planets []Planet
    by func(p1,p2 *Planet) bool
}

func (s *planetSorter)Len()int{
    return len(s.planets)
}

func (s *planetSorter)Swap(i,j int){
    s.planets[i],s.planets[j] = s.planets[j],s.planets[i]
}

func (s *planetSorter)Less(i,j int) bool{
    return s.by(&s.planets[i],&s.planets[j])
}

var planets=[]Planet{
    {"Xie",0.12,0.56},
    {"Mercury",0.0023,0.2},
    {"Venus",0.9877,0.7},
    {"Mars",0.127,0.7},
    {"Sddsd",0.565,0.43},
    {"Aers",0.2344,0.76},
}


func main(){
   name:=func(p1,p2 *Planet)bool{
       return p1.name < p2.name
   }
   By(name).Sort(planets)
   fmt.Println("by name:",planets)
}