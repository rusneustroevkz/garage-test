package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const botToken = "6463773243:AAGw6eg9Hvt3DoVJbHHfLX9kYwn-kTIn2SM"
const SemenUserID = 939464313

func NewTelegram(ctx context.Context) error {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		return err
	}

	b.Start(ctx)

	return err
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.From.ID == SemenUserID {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Semen evanay"),
		})
	}

}
