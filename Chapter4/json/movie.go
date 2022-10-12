package main

//Names of JSON fields must be CAPITALS

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct{
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main(){
	var movies = []Movie{
		{Title: "casa", Year: 1942, Color: false, Actors: []string{"hump","bob"}},
		{Title: "22s", Year: 1942, Color: true, Actors: []string{"hump","bob"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("failed")
	}
	fmt.Printf("%s\n", data)

	data1, err1 := json.MarshalIndent(movies, "", "  ")
	if err1 != nil {
		log.Fatalf("failed")
	}
	fmt.Printf("%s\n", data1)

}