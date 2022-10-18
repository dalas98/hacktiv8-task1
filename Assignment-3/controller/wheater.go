package controller

import (
	"dalas98/golang-lesson/Assignment-3/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func WriteToJson() {
	randomWater := randNumber(100)
	randomWind := randNumber(100)

	data := models.Payload{
		Status: models.Status{
			Water: randomWater,
			Wind:  randomWind,
		},
	}
	log.Println("Generated ", data)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data.json", file, 0644)
}

func LoadJson(w http.ResponseWriter, r *http.Request) {
	// var status models.Payload
	WriteToJson()
	file, err := os.OpenFile("data.json", os.O_RDONLY, 0666)
	if err != nil {
		// return nil, err
		w.Write([]byte("failed load data.json"))
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte("failed read data.json"))
		// return nil, err
	}
	// json.Unmarshal(b, &status)

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Write([]byte(b))
	// return &status, nil
}

func randNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
