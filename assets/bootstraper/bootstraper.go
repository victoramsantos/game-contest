package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	url := "http://localhost:8080"
	client := http.Client{
		Timeout: time.Second * 2,
	}

	if !isAppRunning(client, url) {
		log.Fatal("application is not running")
	}

	characters := createRandomCharacters(client, url, 100)
	startFights(client, url, characters, 30)

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

func createRandomCharacters(client http.Client, url string, amountOfChars int) []string {
	characters := make([]string, 0)
	for i := 0; i < amountOfChars; i++ {
		characters = append(characters, createCharacter(client, url))
	}

	return characters
}

func createCharacter(client http.Client, url string) string {
	path := url + "/character"
	name := buildName()
	content := map[string]string{
		"character_name": name,
		"class_name":     randomClass(),
	}

	body, _ := json.Marshal(content)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("/create %v", content)

	client.Do(req)

	return name
}

func buildName() string {
	pseudoNames := []string{"john", "elvis", "ster", "yolo", "xablau", "checo", "verstapen", "neymar", "airton", "allan"}
	pseudoSurname := []string{"wich", "presley", "onell", "newton", "hamilton", "perez", "senna", "prost"}
	return pseudoNames[randomizer().Intn(len(pseudoNames))] + "_" + pseudoNames[randomizer().Intn(len(pseudoSurname))]
}

func randomClass() string {
	classes := []string{"Warrior", "Thief", "Mage"}
	return classes[randomizer().Intn(len(classes))]
}

func startFights(client http.Client, url string, characters []string, amountOfFights int) {
	for i := 0; i < amountOfFights; i++ {
		attacker := randomizer().Intn(len(characters))
		opponent := randomizer().Intn(len(characters))

		fight(client, url, characters[attacker], characters[opponent])
		time.Sleep(time.Duration(randomizer().Intn(5)) * time.Second)
	}

}

func fight(client http.Client, url string, attacker string, opponent string) {
	path := url + "/game/start"

	content := map[string]string{
		"character_a": attacker,
		"character_b": opponent,
	}
	body, _ := json.Marshal(content)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	logger, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("/game/start %v", content)

	fmt.Println(string(logger))
}

func randomizer() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
