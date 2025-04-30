package main

import (
	"github.com/awwwm4n/expense_tracker/src/bot"
	"github.com/awwwm4n/expense_tracker/src/config"
)

func main() {
	cfg := config.LoadEnv()

	bot.Start(cfg)
}
