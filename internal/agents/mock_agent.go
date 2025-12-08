package agents

import "github.com/eduardongomes/gcai/internal/config"

type MockAgent struct {
	GetCommitCalled  bool
	MakeCommitCalled bool
	EditCalled       bool
	GetCommitReturn  string
}

func (agent *MockAgent) GetDiff() (string, error) {
	return "", nil
}

func (m *MockAgent) GetCommit(c config.ConfigMethods) (string, error) {
	m.GetCommitCalled = true
	return "commit-msg", nil
}

func (m *MockAgent) Edit(msg string) (string, error) {
	m.EditCalled = true
	return msg + "-edited", nil
}

func (m *MockAgent) MakeCommit(msg string) error {
	m.MakeCommitCalled = true
	return nil
}

func NewMockAgent() *MockAgent {
	return &MockAgent{}
}
