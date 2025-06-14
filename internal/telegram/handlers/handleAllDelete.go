package handlers

import (
	answer "github.com/cOsm1cs/World-Cards-master/internal/telegram/service"
	"github.com/cOsm1cs/World-Cards-master/internal/logger"	
	tele "gopkg.in/telebot.v4"
)

// Команда, которая удаляет все добавленные слова
func HandleAllDelete(c tele.Context) error {
	userID := c.Sender().ID
	stateMutex.Lock()
	defer stateMutex.Unlock()

	// Проверяем, есть ли у пользователя добавленные слова
	err := handleDeleteCheckingAddedWords(c, userID)
	if err != nil {
		return err
	}
	
	return logger.SendWithLogError(c, answer.HandleAllDeleteWorks)
}

func handleDeleteCheckingAddedWords(c tele.Context, userID int64) error {
	if len(userWordPairs[userID]) == 0 {
		return logger.SendWithLogError(c, answer.HandleAllDeleteNotWorks)
	}
	delete(userWordPairs, userID)

	return nil
}
