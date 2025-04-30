package bot

import (
	"log"
	"time"

	"github.com/awwwm4n/expense_tracker/src/config"
	"github.com/awwwm4n/expense_tracker/src/sheets"
	tb "gopkg.in/telebot.v3"
)

func Start(cfg config.Config) {
	pref := tb.Settings{
		Token:  cfg.TelegramBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	srv, err := sheets.GetSheetsService(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to sheet: %v", err)
	}

	categories, err := sheets.GetCategories(srv, cfg.SpreadsheetID)
	if err != nil || len(categories) == 0 {
		log.Fatalf("Failed to load categories: %v", err)
	}

	paymentMethods, err := sheets.GetPaymentMethods(srv, cfg.SpreadsheetID)
	if err != nil || len(paymentMethods) == 0 {
		log.Fatalf("Failed to load payment methods: %v", err)
	}

	InitMarkup(b, categories, paymentMethods)
	RegisterHandlers(b, cfg, srv)

	log.Println("✅ Bot started.")
	b.Start()
}
