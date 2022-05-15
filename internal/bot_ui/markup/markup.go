package markup

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BoolToEmoji(value bool) rune {
	if value == true {
		return '✅'
	} else {
		return '☑'
	}
}

func SourceString(sourceName string, isActive bool) string {
	return fmt.Sprintf("%c %s", BoolToEmoji(isActive), sourceName)
}

func Keyboard(buttons [][]tgbotapi.InlineKeyboardButton) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func ButtonRow(button tgbotapi.InlineKeyboardButton) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(button)
}

func Button(text, data string) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(text, data)
}

func KeyboardBackToMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(ButtonBackToMenu())
}

func ButtonBackToMenu() []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(Button("🔙 Back to Menu", "menu"))
}

func ButtonBackToList() []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(Button("🔙 Back to Sources", "list"))
}

func ButtonBackToSource(sourceId int64) []tgbotapi.InlineKeyboardButton {
	data := fmt.Sprintf("source_menu %d", sourceId)
	return ButtonRow(Button("🔙 Back to Source", data))
}
