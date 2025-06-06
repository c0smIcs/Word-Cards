package logger

import (
	"context"
	"log/slog"
	"os"

	tele "gopkg.in/telebot.v4"
)

func SendWithLogInfoStart(c tele.Context, message string) error {
	opts := &slog.HandlerOptions{
		AddSource: true,
	}
	Logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	userID := c.Sender().ID
	Logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Отправлено приветствие от бота пользователю",
		slog.Int64("id", userID),
		slog.Any("message", message),
	)

	return nil
}

func SendWithLogInfoTranslation(c tele.Context, word, translation string) error {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	Logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	userID := c.Sender().ID
	Logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Успешно добавлено пара слов",
		slog.Int64("id", userID),
		slog.Group(
			"Характеристика добавленного слова",
			slog.String("word", word),
			slog.String("translate", translation),
		),
	)

	return nil
}

func SendWithLogError(c tele.Context, message string) error {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true, // включаем источник журнала
	}
	Logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	userID := c.Sender().ID
	err := c.Send(message)
	if err != nil {
		Logger.LogAttrs(
			context.Background(),
			slog.LevelError,
			"Ошибка при отправке сообщения",
			slog.Int64("id", userID),
			slog.Group("Характеристика ошибки сообщения",
				slog.String("message", message),
				slog.Any("error", err),
			),
		)
	}

	return nil
}
