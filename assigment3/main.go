package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"path"
	"time"
)

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"waterStatus"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"windStatus"`
}

var (
	Water       int
	WaterStatus string
	Wind        int
	WindStatus  string
)

func main() {
	go changeWaterWind()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var template, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := Data{
			Status{
				Water,
				WaterStatus,
				Wind,
				WindStatus,
			},
		}

		err = template.Execute(w, res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		res := Data{
			Status{
				Water,
				WaterStatus,
				Wind,
				WindStatus,
			},
		}

		json.NewEncoder(w).Encode(res)
	})
	http.ListenAndServe(":8080", nil)
}

func changeWaterWind() {
	for {
		Water = rand.Intn(100)
		if Water < 5 {
			WaterStatus = "Aman"
		} else if Water > 5 && Water < 9 {
			WaterStatus = "Siaga"
		} else {
			WaterStatus = "Bahaya"
		}

		Wind = rand.Intn(100)

		if Wind < 6 {
			WindStatus = "Aman"
		} else if Wind > 5 && Water < 16 {
			WindStatus = "Siaga"
		} else {
			WindStatus = "Bahaya"
		}

		time.Sleep(15 * time.Second)
	}
}
