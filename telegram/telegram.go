package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	gaylangPartyChatID = -1002462478085
	garageChatID       = -1001787973763
)
const botToken = "6463773243:AAGw6eg9Hvt3DoVJbHHfLX9kYwn-kTIn2SM"
const (
	SemenID  = 939464313
	RuslanID = 717917653
	PetrID   = 1059516310
	DanID    = 304396307
)

type Telegram struct {
	pollID        string
	pollMessageID int
}

func NewTelegram(ctx context.Context) error {
	telegram := Telegram{}

	opts := []bot.Option{
		bot.WithDefaultHandler(telegram.handler),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		return err
	}

	b.Start(ctx)

	return err
}

func (t *Telegram) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %v\n", err)
		}
	}()

	t.poll(ctx, b, update)
	t.stopPoll(ctx, b, update)

	if update != nil && update.Message != nil && update.Message.From != nil {
		fmt.Println(
			update.Message.From.ID,
			update.Message.From.FirstName,
			update.Message.From.LastName,
			update.Message.From.Username,
		)
	}
}

func (t *Telegram) stopPoll(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil {
		return
	}
	if update.Message == nil {
		return
	}
	if update.Message.Text != "/stoppoll" {
		return
	}

	params := &bot.StopPollParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: t.pollMessageID,
	}
	respStopPoll, err := b.StopPoll(ctx, params)
	if err != nil {
		fmt.Printf("stop poll: %v", err)
		return
	}

	pollOption := models.PollOption{}

	for _, option := range respStopPoll.Options {
		if option.VoterCount > pollOption.VoterCount {
			pollOption = option
		}
	}

	reqSendMessage := bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Сегодня будем ебать - %s!", pollOption.Text),
	}
	_, err = b.SendMessage(ctx, &reqSendMessage)
	if err != nil {
		fmt.Printf("send poll result: %v", err)
		return
	}
}

func (t *Telegram) poll(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil {
		return
	}
	if update.Message == nil {
		return
	}
	if update.Message.Text != "/poll" {
		return
	}

	sendPollParams := &bot.SendPollParams{
		ChatID:     update.Message.Chat.ID,
		Question:   "Кого сегодня будем ебать?",
		OpenPeriod: 600,
		Options: []models.InputPollOption{
			{
				Text: "Семен",
				TextEntities: []models.MessageEntity{
					{
						Type: models.MessageEntityTypeBold,
						User: &models.User{
							ID: SemenID,
						},
					},
				},
			},
			{
				Text: "Пека",
				TextEntities: []models.MessageEntity{
					{
						Type: models.MessageEntityTypeBold,
						User: &models.User{
							ID: PetrID,
						},
					},
				},
			},
			{
				Text: "Данил",
				TextEntities: []models.MessageEntity{
					{
						Type: models.MessageEntityTypeBold,
						User: &models.User{
							ID: DanID,
						},
					},
				},
			},
		},
	}

	respPollParams, err := b.SendPoll(ctx, sendPollParams)
	if err != nil {
		fmt.Printf("send poll: %v", err)
		return
	}

	t.pollID = respPollParams.Poll.ID
	t.pollMessageID = respPollParams.ID
	fmt.Println("pollID: ", respPollParams.Poll.ID, "messageID", respPollParams.ID)
}
