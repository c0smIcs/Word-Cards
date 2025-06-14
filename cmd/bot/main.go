package main

import (
	"fmt"
	"sync"
	"github.com/cOsm1cs/World-Cards-master/internal/telegram"
)

var (
	userState = make(map[int64]string)
	stateMu   sync.RWMutex
)

func main() {
	err := telegram.InitBot()
	if err != nil {
		fmt.Println(err)
		return
	}
}
