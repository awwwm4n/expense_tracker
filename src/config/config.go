package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	SpreadsheetID    string
	GoogleCredsJSON  string
}

var AppConfig Config

func LoadEnv() Config {
	_ = godotenv.Load()

	AppConfig = Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		SpreadsheetID:    os.Getenv("SPREADSHEET_ID"),
		GoogleCredsJSON:  os.Getenv("GOOGLE_CREDENTIALS_JSON"),
	}

	if AppConfig.TelegramBotToken == "" || AppConfig.SpreadsheetID == "" || AppConfig.GoogleCredsJSON == "" {
		log.Fatal("Missing required environment variables")
	}

	return AppConfig
}
