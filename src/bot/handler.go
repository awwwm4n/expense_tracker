package bot

import (
	"strconv"
	"strings"

	"github.com/awwwm4n/expense_tracker/src/config"
	sh "github.com/awwwm4n/expense_tracker/src/sheets"
	"google.golang.org/api/sheets/v4"
	tb "gopkg.in/telebot.v3"
)

func RegisterHandlers(b *tb.Bot, cfg config.Config, srv *sheets.Service) {
	b.Handle("/start", func(c tb.Context) error {
		return c.Send("Welcome! Type /expense to log a new expense.")
	})

	b.Handle("/expense", func(c tb.Context) error {
		return c.Send("Choose a category:", GetCategoryMarkup())
	})

	b.Handle(tb.OnText, func(c tb.Context) error {
		userID := c.Sender().ID
		category, hasCat := userCategory[userID]
		method, hasPay := userPaymentMethod[userID]

		if !hasCat || !hasPay {
			return nil // User hasn't selected both category and method
		}

		text := c.Text()
		parts := strings.SplitN(text, " ", 2)
		if len(parts) != 2 {
			return c.Send("Invalid format. Send like: `250 Snacks`", tb.ModeMarkdown)
		}

		amount, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return c.Send("Invalid amount. Use something like: `250 Snacks`")
		}
		description := parts[1]

		paidBy := "Pari"
		if c.Sender().Username == "awwwm4n" {
			paidBy = "Aman"
		}

		err = sh.AppendExpense(srv, cfg.SpreadsheetID, category, amount, description, paidBy, method)
		if err != nil {
			return c.Send("❌ Failed to log expense. Try again with /expense")
		}

		delete(userCategory, userID)
		delete(userPaymentMethod, userID)

		return c.Send("✅ Expense logged! Use /expense to add another.")
	})
}
