package handlers

import (
	"github.com/cOsm1cs/World-Cards-master/answer"
	"github.com/cOsm1cs/World-Cards-master/logger"
	tele "gopkg.in/telebot.v4"
)

func HandleInstruction(c tele.Context) error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	return logger.SendWithLogError(c, answer.HandleInstruction)
}
