package db

import (
	"encoding/json"
	"errors"
	"log"
	"net-sender/internal/data"
	"os"
)

const directoryPath = "data/"
const dataPath = directoryPath + "data.json"
const directoryPermissions = 0777
const dataPermissions = 0644

func remove(s []data.Mailing, i int) []data.Mailing {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func Init() {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		log.Println("Data directory doesn't exist. Creating...")
		createDirectory()
	} else {
		log.Println("Data directory already exists")
	}

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		log.Println("Data file doensn't exist. Creating...")
		createDatafile()
	} else {
		log.Println("Data file already exists")
	}
}

func createDirectory() {
	err := os.Mkdir(directoryPath, os.FileMode(directoryPermissions))
	if (err != nil) {
		log.Println(err)
	}
}

func createDatafile() {
	file, err := os.Create(dataPath)
	if (err != nil) {
		log.Println(err)
	}
	defer file.Close()
}

func GetMailing(chatID int) (*data.Mailing, error) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Println(err)
	}
	mailings := []data.Mailing{}
	json.Unmarshal(file, &mailings)
	for _, mailing := range mailings {
		if mailing.ChatID == chatID {
			return &mailing, nil
		}
	}
	return &data.Mailing{}, errors.New(data.FailureGetText)
}

func UpdateMailing(updatedMailing *data.Mailing) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Println(err)
	}

	mailings := []data.Mailing{}
	json.Unmarshal(file, &mailings)
	for i := 0; i < len(mailings); i++ {
		if (mailings[i].ChatID == updatedMailing.ChatID) {
			mailings = remove(mailings, i)
		}
	}

	mailings = append(mailings, *updatedMailing)
	mailingsBytes, err := json.MarshalIndent(mailings, "", "    ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(dataPath, mailingsBytes, os.FileMode(dataPermissions))
	if err != nil {
		log.Println(err)
	}
}
