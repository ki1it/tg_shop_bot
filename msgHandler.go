package main

import "log"

func MsgHandler(msg string) {
	if msg == "/start" {
		log.Println("123")
	}
}
