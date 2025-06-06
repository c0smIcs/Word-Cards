package handlers

import (
	"github.com/cOsm1cs/World-Cards-master/answer"
	"github.com/cOsm1cs/World-Cards-master/logger"
	tele "gopkg.in/telebot.v4"
)

// Команда, которая удаляет все добавленные слова
func HandleAllDelete(c tele.Context) error {
	userID := c.Sender().ID
	stateMutex.Lock()
	defer stateMutex.Unlock()

	// Проверяем, есть ли у пользователя добавленные слова
	if len(userWordPairs[userID]) == 0 {
		return logger.SendWithLogError(c, answer.HandleAllDeleteNotWorks)
	}
	delete(userWordPairs, userID)
	return logger.SendWithLogError(c, answer.HandleAllDeleteWorks)
}
