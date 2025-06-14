package handlers

import (
	"strings"

	gt "github.com/bas24/googletranslatefree"
	tele "gopkg.in/telebot.v4"

	answer "github.com/cOsm1cs/World-Cards-master/internal/telegram/service"
	"github.com/cOsm1cs/World-Cards-master/internal/logger"
)

func HandleAdd(c tele.Context) error {
	userID := c.Sender().ID  // Получаем ID пользователя
	text := c.Message().Text // Получаем текст сообщения

	word, err := addFieldsText(c, text)
	if err != nil {
		return err
	}
	
	translation, err := addTranslationWord(c, word)
	if err != nil {
		return err
	}

	// Создаем пару слов
	pair, err := addPairs(word, translation)
	if err != nil {
		return err
	}

	stateMutex.Lock() // Блокируем доступ к данным для безопасности
	defer stateMutex.Unlock() // Разблокируем доступ

	// Добавляем пару в словарь
	userWordPairs[userID] = append(userWordPairs[userID], pair)

	logger.SendWithLogInfoTranslation(c, word, translation)

	// Подтверждаем добавление
	return logger.SendWithLogError(c, "Пара слов добавлена: "+word+" - "+translation)
}

func addFieldsText(c tele.Context, text string) (string, error) {
	parts := strings.Fields(text) // Разделяем текст на слова по пробелам
	if len(parts) < 2 {           // Если нет аргумента после /add
		return "", logger.SendWithLogError(c, answer.HandleAddCorrectUse) //const HandleAddCorrectUse = "Используй: /add <слово>"
	}
	word := strings.Join(parts[1:], " ") // Склеиваем все части после команды в одно слово

	return word, nil
}

func addTranslationWord(c tele.Context, word string) (string, error) {
	translation, err := gt.Translate(word, "ru", "en") // Переводим на английский
	if err != nil {                                    // Если ошибка перевода
		return "", logger.SendWithLogError(c, "Ошибка перевода: "+err.Error()) // Сообщаем пользователю
	}

	return translation, nil
}

func addPairs(word, translation string) (WordPair, error) {
	return WordPair{
		Original: word,
		Translate: translation,
	}, nil
}
