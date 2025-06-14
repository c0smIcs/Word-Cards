package handlers

import (
	answer "github.com/cOsm1cs/World-Cards-master/internal/telegram/service"
	"github.com/cOsm1cs/World-Cards-master/internal/logger"
	tele "gopkg.in/telebot.v4"
)

func HandleInstruction(c tele.Context) error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	return logger.SendWithLogError(c, answer.HandleInstruction)
}
