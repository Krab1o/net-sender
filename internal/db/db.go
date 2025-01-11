package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net-sender/internal/data"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB
var directoryPath = "data/"
var directoryPermissions = 0777

func createDirectory() {
	err := os.Mkdir("data", os.FileMode(directoryPermissions))
	if err != nil {
		log.Println(err)
	}
}

func InitDB() {
	var err error
	_, err = os.Stat("data/")
	if errors.Is(err, os.ErrNotExist) {
		createDirectory()
	}
	DB, err = sql.Open("sqlite3", "data/app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
	 	log.Println(err) // Log an error and stop the program if the database can't be opened
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
CREATE TABLE IF NOT EXISTS Mailings (
	chat_id INTEGER NOT NULL PRIMARY KEY,
	login STRING NOT NULL,
	time DATETIME NOT NULL,
	download BIGINT NOT NULL,
	upload BIGINT NOT NULL
);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}
}

func GetMailing(chatID int) (*data.Mailing, error){
	stmt, err := DB.Prepare("SELECT * FROM Mailings WHERE chat_id = $1")
	if err != nil {
		log.Println("Failed to prepare query", err)
		return &data.Mailing{}, errors.New("Не удалось получить данные")
	}
	defer stmt.Close()

	mailing := data.Mailing{}
	row := stmt.QueryRow(chatID)
	err = row.Scan(
		&mailing.ChatID, 
		&mailing.Login,
		&mailing.LastTime,
		&mailing.Download,
		&mailing.Upload,
	)
	if err != nil {
		return &mailing, err
	}
	return &mailing, nil
}

func InsertMailing(newMailing *data.Mailing) error {
	insertStmt, err := DB.Prepare(
`INSERT INTO Mailings (chat_id, login, time, download, upload) VALUES
	($1, $2, $3, $4, $5)
`)
	if err != nil {
		log.Println("Failed to prepare statement")
		return err
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(
		newMailing.ChatID,
		newMailing.Login,
		newMailing.LastTime,
		newMailing.Download,
		newMailing.Upload,
	)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: Mailings.chat_id" {
			log.Printf("User %d already exists. " + 
			"Trying to update existing row...", newMailing.ChatID)
			UpdateMailing(newMailing)		
		} else {
			log.Printf("Failed to insert user %s: %v", newMailing.Login, err)
			return err
		}		
	}
	return nil
}

func UpdateMailing(updatedMailing *data.Mailing) error {
	updateStmt, err := DB.Prepare(
`UPDATE Mailings SET
	chat_id = $1,
	login = $2,
	time = $3,
	download = $4,
	upload = $5
WHERE
	chat_id = $1
`)
	if err != nil {
		log.Println("Failed to prepare statement:\n " + err.Error())
		return err
	}

	_, err = updateStmt.Exec(
		&updatedMailing.ChatID, 
		&updatedMailing.Login, 
		&updatedMailing.LastTime, 
		&updatedMailing.Download, 
		&updatedMailing.Upload, 
	)
	if err != nil {
		log.Printf("Failed to update user %s: %v", updatedMailing.Login, err)
		return err
	}
	return nil
}