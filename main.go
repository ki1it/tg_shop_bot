package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/ki1it/tg_shop_bot/commands"
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

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	ok, msg := commands.HandleCommand(update)
	if ok {
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}


}

func main() {
	_ = godotenv.Load()
	log.Println("App started")

	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Println( "no telegram bot token")
		return
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println( "bot err ", err)
		return
	}
	_, ok = os.LookupEnv("BOT_MODE_UPDATES")
	if ok {
		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, _ := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			//
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//msg.ReplyToMessageID = update.Message.MessageID
			//
			//bot.Send(msg)
			go handleUpdate(update, bot)
		}
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got req", r.Body)
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
			go handleUpdate(update, bot)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Path[len("/greet/"):]
			fmt.Fprintf(w, "Hello %s\n", name)
		})

		http.ListenAndServe(":9990", nil)
	}
}
