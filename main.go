package main

import (
	"log"
	"net-sender/internal/bot"
	"net-sender/internal/data"
	"net-sender/internal/db"
	"net-sender/internal/requests"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func parseCommand(message string) (string, bool) {
	rg, err := regexp.Compile(`^\/[a-zA-Z_]{1,}\b`)
	if err != nil {
		log.Println(err)
	}
	if index := rg.FindStringIndex(message); index != nil {
		return message[index[0]:index[1]], true
	}
	return data.None.String(), false
}

//TODO: docker compose
//TODO: add database/redis

func main() {
	err := godotenv.Load(".env")
  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}

	site := &requests.Site{
		Host : os.Getenv("HOST"),
		Port : os.Getenv("PORT"),
		Panel : os.Getenv("PANEL"),
	}
	site.LoginRequest()

	conf := &data.Config{
		Offset: 0,
		Timeout: 60,
	}

	commands := []data.Command{
		{Title: "/start", Desc:	data.StartDescription},
		{Title: "/set_login", Desc: data.SetLoginDescription},
		{Title: "/get_diff", Desc: data.GetDiffDescription},
		{Title: "/get_login", Desc: data.GetLoginDescription},
		{Title: "/get_stat", Desc: data.GetStatDescription},
		{Title: "/help", Desc: data.HelpDescription},
	}

	db.InitDB()
	updates := bot.InitBot(conf, commands)

	for update := range updates {
		if command, ok := parseCommand(update.Message.Text); ok != false {
			log.Println(command)
			switch command {
			case "/start":
				bot.Start(update)
			case "/set_login":
				bot.SetLogin(update, site)
			case "/get_diff":
				bot.GetDiff(update, site)
			case "/get_login":
				bot.GetLogin(update)
			case "/get_stat":
				bot.GetStat(update, site)
			case "/help":
				bot.Help(update)
			default:
				bot.SendMessage(update, "Неизвестная команда")
				//do nothing
			}
		} else {
			bot.SendMessage(update, "Я понимаю только команды :(")
		}
	}
}