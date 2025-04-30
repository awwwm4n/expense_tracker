package bot

import (
	"fmt"
	"strings"

	tb "gopkg.in/telebot.v3"
)

var (
	userCategory      = make(map[int64]string)
	userPaymentMethod = make(map[int64]string)
	categoryButtons   = map[string]string{}
	paymentButtons    = map[string]string{}
)

func InitMarkup(b *tb.Bot, categories, paymentMethods []string) {
	for _, cat := range categories {
		uid := "cat_" + strings.ToLower(strings.ReplaceAll(cat, " ", "_"))
		categoryButtons[uid] = cat

		btn := tb.InlineButton{Unique: uid, Text: cat}
		b.Handle(&btn, func(c tb.Context) error {
			_ = c.Respond()
			userID := c.Sender().ID
			userCategory[userID] = categoryButtons[c.Callback().Unique]
			return c.Send(fmt.Sprintf("Category: *%s*\nNow select payment method:", categoryButtons[c.Callback().Unique]), GetPaymentMarkup())
		})
	}

	for _, method := range paymentMethods {
		uid := "pay_" + strings.ToLower(strings.ReplaceAll(method, " ", "_"))
		paymentButtons[uid] = method

		btn := tb.InlineButton{Unique: uid, Text: method}
		b.Handle(&btn, func(c tb.Context) error {
			_ = c.Respond()
			userID := c.Sender().ID
			userPaymentMethod[userID] = paymentButtons[c.Callback().Unique]

			category := userCategory[userID]
			return c.Send(fmt.Sprintf("✅ *%s* via *%s*\nNow send amount and note like:\n`250 Coffee at Starbucks`", category, paymentButtons[c.Callback().Unique]), tb.ModeMarkdown)
		})
	}
}

func GetCategoryMarkup() *tb.ReplyMarkup {
	markup := &tb.ReplyMarkup{}
	var rows [][]tb.InlineButton
	for uid, label := range categoryButtons {
		rows = append(rows, []tb.InlineButton{{Unique: uid, Text: label}})
	}
	markup.InlineKeyboard = rows
	return markup
}

func GetPaymentMarkup() *tb.ReplyMarkup {
	markup := &tb.ReplyMarkup{}
	var rows [][]tb.InlineButton
	for uid, label := range paymentButtons {
		rows = append(rows, []tb.InlineButton{{Unique: uid, Text: label}})
	}
	markup.InlineKeyboard = rows
	return markup
}
