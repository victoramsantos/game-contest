package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080"
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	if !isAppRunning(client, url) {
		log.Fatal("application is not running")
	}

	characters := createRandomCharacters(10)
	startFights(characters)

}

func isAppRunning(client http.Client, url string) bool {
	path := url + "/health"
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	return res.StatusCode == 200
}

func createRandomCharacters(amountOfChars int) {

}

func createCharacter(client http.Client, url string) {
	path := url + "/character"

	body, err := json.Marshal(map[string]string{
		"character_name": buildName(),
		"class_name":     randomClass(),
	})

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
}

func buildName() string {
	pseudoNames := []string{"john", "elvis", "ster", "yolo", "xablau"}
	name := pseudoNames[randomizer().Intn(len(pseudoNames))]
	return name + string(randomizer().Intn(100))
}

func randomClass() string {
	classes := []string{"Warrior", "Thief", "Mage"}
	return classes[randomizer().Intn(len(classes))]
}

func startFights(characters []string) {

}

func randomizer() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
