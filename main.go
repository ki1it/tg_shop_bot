package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonSwitch("2sw", "open 2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func getEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func gracefulExit(w http.ResponseWriter, text string) {
	log.Println(text)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("App started")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got req", r.Body)
		token, ok := os.LookupEnv("BOT_TOKEN")
		log.Println(token)
		if !ok {
			gracefulExit(w, "no telegram bot token")
			return
		}
		bot, err := tgbotapi.NewBotAPI(token)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			gracefulExit(w, "no body")
			return
		}
		var update tgbotapi.Update

		if err := json.Unmarshal(body, &update); err != nil {
			gracefulExit(w, "body parse fucking error")
			return
		}
		MsgHandler(update.Message.Text)

		log.Println("текст из сообщения", update.Message.Text)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// todo: looks like send message can be done via response on this webhook request
		// but let's use this API for now
		bot.Send(msg)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greet/"):]
		fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe(":9990", nil)

}
