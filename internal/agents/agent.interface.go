package agents

import "github.com/eduardongomes/gcai/internal/config"

type AgentMethods interface {
	GetCommit(c config.ConfigMethods) (string, error)
	Edit(v string) (string, error)
	GetDiff() (string, error)
	MakeCommit(msg string) error
}
