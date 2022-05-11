package commands

import "AggreBot/internal/pkg/grpc_client"

type Manager struct {
	backend         *grpc_client.Client
	commandHandlers map[string]commandHandler
}

type commandHandler func(*command) string

func NewManager(grpcClient *grpc_client.Client) *Manager {
	var m Manager
	m = Manager{
		backend: grpcClient,
		commandHandlers: map[string]commandHandler{
			"/start":  m.cmdStart,
			"/add":    m.cmdAdd,
			"/list":   m.cmdList,
			"/filter": m.cmdFilter,
			"/rename": m.cmdRename,
			"/toggle": m.cmdToggle,
			"/delete": m.cmdDelete,
		},
	}
	return &m
}

func (m *Manager) Execute(c *command) string {
	cmdHandler, ok := m.commandHandlers[c.cmdName]
	if ok {
		reply := cmdHandler(c)
		return reply
	} else {
		return ""
	}
}
