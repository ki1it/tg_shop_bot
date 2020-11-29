package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func HandleCommand(update tgbotapi.Update) (bool, tgbotapi.MessageConfig){
	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			msg.Text = "type /sayhi or /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "withArgument":
			msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
		case "html":
			msg.ParseMode = "html"
			msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
		default:
			msg.Text = "I don't know that command"
		}
		return true, msg
	}
	return false, tgbotapi.MessageConfig{}
}