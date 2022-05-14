package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/pkg/grpc_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Manager struct {
	backend          *grpc_client.Client
	callbackHandlers map[string]callbackHandler
}

type callbackHandler func(*command.Command) (string, *tgbotapi.InlineKeyboardMarkup)

func NewManager(grpcClient *grpc_client.Client) *Manager {
	var m Manager
	m = Manager{
		backend: grpcClient,
		callbackHandlers: map[string]callbackHandler{
			"menu":          m.cbMenu,
			"list":          m.cbList,
			"add":           m.cbAdd,
			"filter":        m.cbFilter,
			"filter_menu":   m.cbFilterMenu,
			"filter_remove": m.cbFilterRemove,
			"source_menu":   m.cbSourceMenu,
			"source_rename": m.cbSourceRename,
			//"source_enable": m.cbSourceEnable,
			//"source_disable": m.cbSourceDisable,
			//"source_remove": m.cbSourceRemove,
		},
	}
	return &m
}

func (m *Manager) Execute(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	cmdHandler, ok := m.callbackHandlers[c.Cmd]
	if ok {
		reply, markup := cmdHandler(c)
		return reply, markup
	} else {
		return m.ExecuteMenu(c)
	}
}

func (m *Manager) ExecuteMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	return m.callbackHandlers["menu"](c)
}
