package main

import (
	"fmt"
	"log"
	"net/http"
)

var exercises = []map[string]interface{}{
	{
		"id": "1cf5cdbd-f6c5-47a3-a9e9-117339034f16",
		"name": "bench press", 
		"description": "horizontal lift moving weight perpendicular to chest",
		"muscle_group": "chest",
	},
	{
		"id": "4e9b4f4e-0df0-4665-8777-c93acf678229",
		"name": "back squat",
		"description": "vertical lift moving weight on back",
		"muscle_group": "legs",
	},
	{
		"id": "b7a28ef5-3775-469f-828c-d19433fd9b7c",
		"name": "pull ups",
		"description": "vertical lift moving body weight up and down",
		"muscle_group": "back",
	},
}

func main () {
	fmt.Println("This workout tracker is working!!!!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /exercises", getExercises)
	mux.HandleFunc("GET /exercises/{id}", getExercise)


	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}