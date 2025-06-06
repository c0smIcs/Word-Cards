package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cOsm1cs/World-Cards-master/handlers"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

var(
	userState = make(map[int64]string)
	stateMu   sync.RWMutex
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки файла .env, предполагается, что переменные окружения установлены:", err)
		return
	}

	apiToken := os.Getenv("TOKEN")
	
	if apiToken == "" {
		fmt.Println("Ошибка: Переменная окружения TOKEN не найдена")
		return
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:  apiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", 	 handlers.HandleStart)
	bot.Handle("/go",   	 handlers.HandleGo)
	bot.Handle("/add",       handlers.HandleAdd)
	bot.Handle("/alldelete", handlers.HandleAllDelete)
	bot.Handle(tele.OnText,  handlers.HandleText)
	bot.Handle("/instruction", handlers.HandleInstruction)

	log.Println("Бот запущен")
	bot.Start()
}
