package states

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/user_state"
	"AggreBot/internal/pkg/grpc_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Manager struct {
	backend       *grpc_client.Client
	stateHandlers map[int]stateHandler
}

type stateHandler func(*command.Command) (string, *tgbotapi.InlineKeyboardMarkup)

func NewManager(grpcClient *grpc_client.Client) *Manager {
	var m Manager
	m = Manager{
		backend: grpcClient,
		stateHandlers: map[int]stateHandler{
			user_state.SourceAdd:    m.stAdd,
			user_state.SourceRename: m.stRename,
			user_state.FilterSet:    m.stFilter,
		},
	}
	return &m
}

func (m *Manager) Execute(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	stHandler, ok := m.stateHandlers[c.UserState.State]
	if ok {
		reply, markup := stHandler(c)
		return reply, markup
	} else {
		log.Printf("states.Execute: can't handle state <%v>", c.UserState.State)
		return errors.ErrInternalError, nil
	}
}
