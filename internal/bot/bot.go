package bot

import (
	"fmt"
	"log"
	"net-sender/internal/data"
	"net-sender/internal/db"
	"net-sender/internal/requests"
	"time"
)

func InitBot(config *data.Config, commands []data.Command) chan data.Update{
	requests.SetCommands(commands)
	return requests.GetUpdatesChan(config)
}

func SendMessage(update data.Update, msg string) {
	requests.SendMessage(update.Message.Chat.ID, msg, data.None)
}

func SetLogin(update data.Update, site *requests.Site) {
	login, ok := parseLogin(update.Message.Text)
	if !ok {
		log.Println("Не удалось распознать логин")
		SendMessage(update, "Не удалось распознать логин")
		return
	}
	mailing := &data.Mailing{
		ChatID: update.Message.Chat.ID,
		Login: login,
		LastTime: time.Time{},
		Download: 0,
		Upload: 0,
	}
	
	db.UpdateMailing(mailing)
}

func GetLogin(update data.Update) {
	mailing, err := db.GetMailing(update.Message.Chat.ID)
	if err != nil {
		SendMessage(update, "В этом чате ещё не был указан логин")
		return
	}
	text := fmt.Sprintf(data.SuccessGetLoginText, mailing.Login)
	SendMessage(update, text)
}

func GetDiff(update data.Update, site *requests.Site) {
	mailing, err := db.GetMailing(update.Message.Chat.ID)
	if err != nil {
		SendMessage(update, err.Error())
		return
	}
	reqData, err := site.GetClientRequest(mailing.Login)
	if err != nil {
		SendMessage(update, err.Error())
		return
	}

	download := formatFileSize(float64(reqData.Obj.Download - mailing.Download))
	upload := formatFileSize(float64(reqData.Obj.Upload - mailing.Upload))
	//TODO: логика на пустую дату 
	parsedTime := mailing.LastTime.Format(time.RFC822)
	msg := fmt.Sprintf(data.SuccessGetDiffText, parsedTime, download, upload)
	SendMessage(update, msg)

	mailing = &data.Mailing{
		ChatID : mailing.ChatID,
		Login : mailing.Login,
		LastTime : time.Now(),
		Download : reqData.Obj.Download,
		Upload : reqData.Obj.Upload,
	}
	db.UpdateMailing(mailing)
}

func GetStat(update data.Update, site *requests.Site) {
	mailing, err := db.GetMailing(update.Message.Chat.ID)
	if err != nil {
		SendMessage(update, err.Error())
		return
	}
	reqData, err := site.GetClientRequest(mailing.Login)
	if err != nil {
		SendMessage(update, err.Error())
		return
	}

	download := formatFileSize(float64(reqData.Obj.Download))
	upload := formatFileSize(float64(reqData.Obj.Upload))
	msg := fmt.Sprintf(data.SuccessGetStatText, download, upload)
	SendMessage(update, msg)
}

func Help(update data.Update) {
	requests.SendMessage(update.Message.Chat.ID, data.StartText, data.MarkdownV2)
}

func Start(update data.Update) {
	requests.SendMessage(update.Message.Chat.ID, data.StartText, data.MarkdownV2)
}