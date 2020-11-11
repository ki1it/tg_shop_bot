package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token, ok := os.LookupEnv("BOT_TOKEN")
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
