package callbacks

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *Manager) registerUser(userId int64) bool {
	err := m.backend.AddUser(userId)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		return false
	}
	return true
}
