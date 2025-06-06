package handlers

import (
	"strings"

	gt "github.com/bas24/googletranslatefree"
	"github.com/cOsm1cs/World-Cards-master/answer"
	"github.com/cOsm1cs/World-Cards-master/logger"

	tele "gopkg.in/telebot.v4"
)

func HandleAdd(c tele.Context) error {
	userID := c.Sender().ID  // Получаем ID пользователя
	text := c.Message().Text // Получаем текст сообщения
	parts := strings.Fields(text) // Разделяем текст на слова по пробелам
	if len(parts) < 2 {           // Если нет аргумента после /add
		return logger.SendWithLogError(c, answer.HandleAddCorrectUse)
	}

	word := strings.Join(parts[1:], " ") // Склеиваем все части после команды в одно слово
	translation, err := gt.Translate(word, "ru", "en") // Переводим на английский
	if err != nil { // Если ошибка перевода
		return logger.SendWithLogError(c, "Ошибка перевода: " + err.Error()) // Сообщаем пользователю
	}

	// Создаем пару слов
	pair := WordPair{Original: word, Translate: translation}

	stateMutex.Lock() // Блокируем доступ к данным (для безопасности)
	
	// Добавляем пару в словарь
	userWordPairs[userID] = append(userWordPairs[userID], pair)
	stateMutex.Unlock() // Разблокируем доступ
	
	logger.SendWithLogInfoTranslation(c, word, translation)
	
	// Подтверждаем добавление
	return logger.SendWithLogError(c, "Пара слов добавлена: " + word + " - " + translation)
}
