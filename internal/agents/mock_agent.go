package agents

import "github.com/eduardongomes/gcai/internal/config"

type MockAgent struct{}

func (agent *MockAgent) GetCommit(config config.ConfigMethods) (string, error) {
	return "", nil
}

func (agent *MockAgent) GetDiff() (string, error) {
	return "", nil
}

func (agent *MockAgent) MakeCommit(msg string) error {
	return nil
}

func (agent *MockAgent) Edit(v string) (string, error) {
	return "", nil
}

func NewMockAgent() *MockAgent {
	return &MockAgent{}
}
