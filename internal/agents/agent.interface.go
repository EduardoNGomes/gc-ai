package agents

import (
	c "github.com/eduardongomes/gcai/internal/config"
)

const Prompt = "You are a senior software engineer. Analyze the given git diff and output only a single-line, professional conventional commit message following the Conventional Commits specification (type, scope if applicable, and short description): \n"

type AgentMethods interface {
	GetCommit(c c.ConfigMethods) (string, error)
	Edit(v string) string
	GetDiff() (string, error)
	MakeCommit(msg string) error
}
