package bot_ui

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/handlers/callbacks"
	"AggreBot/internal/bot_ui/handlers/states"
	"AggreBot/internal/bot_ui/user_state"
	"AggreBot/internal/pkg/grpc_client"
	"AggreBot/internal/pkg/tg_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	tgClient        *tgbotapi.BotAPI
	callbackManager *callbacks.Manager
	stateManager    *states.Manager
	userState       map[int64]*user_state.UserState
}

func NewBot(tgClient *tgbotapi.BotAPI, grpcClient *grpc_client.Client) *Bot {
	return &Bot{
		tgClient:        tgClient,
		callbackManager: callbacks.NewManager(grpcClient),
		stateManager:    states.NewManager(grpcClient),
		userState:       make(map[int64]*user_state.UserState),
	}
}

func (bot *Bot) RunBotLoop() {
	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60
	var userId int64
	var replyText string
	var markup *tgbotapi.InlineKeyboardMarkup
	for u := range bot.tgClient.GetUpdatesChan(uConfig) {
		if u.Message != nil && u.Message.Text != "" {
			userId = u.Message.From.ID
			bot.checkNilUserState(userId)
			replyText, markup = bot.handleMessage(u.Message)
			bot.resetUserState(userId)
			_ = tg_client.SendMessage(bot.tgClient, userId, replyText, markup)
		} else if u.CallbackQuery != nil {
			userId = u.CallbackQuery.From.ID
			bot.checkNilUserState(userId)
			replyText, markup = bot.handleCallback(u.CallbackQuery)
			_ = tg_client.SendCallbackAnswer(bot.tgClient, u.CallbackQuery.ID, "")
			_ = tg_client.UpdateMessage(
				bot.tgClient,
				userId,
				u.CallbackQuery.Message.MessageID,
				replyText,
				markup,
			)
		}
	}
}

func (bot *Bot) handleMessage(msg *tgbotapi.Message) (string, *tgbotapi.InlineKeyboardMarkup) {
	log.Printf("%s: %s", msg.From.UserName, msg.Text)

	userState := bot.userState[msg.From.ID]
	cmd := command.FromMessage(msg, userState)

	switch {
	case cmd.Text == "/start":
		return bot.callbackManager.ExecuteMenu(cmd)
	case userState.State != user_state.Empty:
		return bot.stateManager.Execute(cmd)
	}
	return "", nil
}

func (bot *Bot) handleCallback(query *tgbotapi.CallbackQuery) (string, *tgbotapi.InlineKeyboardMarkup) {
	log.Printf("%s> %s", query.From.UserName, query.Data)

	userState := bot.userState[query.From.ID]
	cmd := command.FromCallbackQuery(query, userState)
	return bot.callbackManager.Execute(cmd)
}

func (bot *Bot) checkNilUserState(userId int64) {
	if _, ok := bot.userState[userId]; !ok {
		bot.userState[userId] = &user_state.UserState{}
	}
}

func (bot *Bot) resetUserState(userId int64) {
	bot.userState[userId].State = user_state.Empty
	bot.userState[userId].Value = 0
}
