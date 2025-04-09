package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func home(w  http.ResponseWriter, r *http.Request) {
	fmt.Println("Home route being hittt!!!!")

	_, err := w.Write([]byte("Welcome home to the workout app"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getExercises(w  http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all exercises")

	js, err := json.Marshal(exercises)
	if err != nil {
		fmt.Println(err)
	}
	
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getExercise(w  http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting singular exercise")

	exerciseID := r.PathValue("id")

	exercise, ok := findExerciseById(exerciseID)
	if !ok {
		http.Error(w, fmt.Sprintf("could not find exercise id: %s", exerciseID), http.StatusNotFound)
	}

	js, err := json.Marshal(exercise)
	if err != nil {
		fmt.Println(err)
	}
	
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func findExerciseById(id string) (map[string]interface{}, bool) {
	for _, exMap := range exercises {
		if exMap["id"] == id {
			return exMap, true
		}
	}

	return nil, false
}