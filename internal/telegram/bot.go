package telegram

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	"github.com/cOsm1cs/World-Cards-master/internal/telegram/handlers"
)

func InitBot() error {
	apiToken, err := InitLoadEnvConfig()
	if err != nil {
		return err
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:  apiToken,
		Poller: &tele.LongPoller{ Timeout: 10 * time.Second },
	})

	if err != nil {
		log.Fatal(err)
	}
	
	bot.Handle("/start", 	   handlers.HandleStart)
	bot.Handle("/go",          handlers.HandleGo)
	bot.Handle("/add",         handlers.HandleAdd)
	bot.Handle("/alldelete",   handlers.HandleAllDelete)
	bot.Handle(tele.OnText,    handlers.HandleText)
	bot.Handle("/instruction", handlers.HandleInstruction)

	log.Println("Бот запущен")
	bot.Start()

	return nil
}

func InitLoadEnvConfig() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("Ошибка загрузки файла .env, предполагается, что переменные окружения установлены")
	}
	
	apiToken := os.Getenv("TOKEN")
	if apiToken == "" {
		return "", fmt.Errorf("Ошибка: Переменная окружения TOKEN не найдена.")
	}

	return apiToken, nil
}
