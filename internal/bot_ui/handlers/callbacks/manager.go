package callbacks

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/grpc_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
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
			"menu": m.cbMenu,

			"list":       m.cbList,
			"source_add": m.cbSourceAdd,
			"filter":     m.cbFilter,

			"source_menu":           m.cbSourceMenu,
			"source_rename":         m.cbSourceRename,
			"source_active_enable":  m.cbSourceActiveEnable,
			"source_active_disable": m.cbSourceActiveDisable,
			"source_remove":         m.cbSourceRemove,

			"filter_menu":   m.cbFilterMenu,
			"filter_remove": m.cbFilterRemove,
		},
	}
	return &m
}

func (m *Manager) Execute(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	var cmdName string
	if cmdEndIndex := strings.Index(c.Text, " "); cmdEndIndex == -1 {
		cmdName = c.Text
	} else {
		cmdName = c.Text[:cmdEndIndex]
	}

	cmdHandler, ok := m.callbackHandlers[cmdName]
	if ok {
		reply, markup := cmdHandler(c)
		return reply, markup
	} else {
		log.Printf("callbacks.manager.Execute: Unknown query: \"%s\"", c.Text)
		return m.ExecuteMenu(c)
	}
}

func (m *Manager) ExecuteMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	return m.callbackHandlers["menu"](c)
}
