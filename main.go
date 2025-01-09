package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type data struct {
	login 		string
	password 	string
}

func panelRequest() {
	data := &loginData{
		host : os.Getenv("HOST"),
		port : os.Getenv("PORT"),
		login : os.Getenv("LOGIN"),
		password : os.Getenv("PASSWORD"),
	}

	url := "https://" + data.login +  

	byteBody, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(byteBody))
	if err != nil {
		log.Println(err)
	}
	
	resp, err := http.DefaultClient.Do(req)
	if (err != nil) {
		log.Println(err)
	}

	fmt.Println(resp.Body)
	

}

func main() {
	godotenv.Load()
	panelRequest()
}