package handlers

import (
	answer "github.com/cOsm1cs/World-Cards-master/internal/telegram/service"
	"github.com/cOsm1cs/World-Cards-master/internal/logger"	
	tele "gopkg.in/telebot.v4"
)

func HandleStart(c tele.Context) error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	err := logger.SendWithLogInfoStart(c, answer.HandleStart)
	if err != nil {
		logger.SendWithLogError(c, answer.HandleStart)
	}
	return err
}
