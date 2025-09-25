package agents

import (
	c "github.com/eduardongomes/gcai/internal/config"
)

const Prompt = `You are a senior software engineer. Analyze the given git diff and output only a single-line, professional conventional commit message following the Conventional Commits specification. Ensure the type and scope are in lowercase, and the message is clear and concise.

Structure:
	- type
	- message

Examples:
- feat: implement clas Login to handler member login
- feat: add internal link suggestion
- fix: handle 404 errors for user profiles
- refactor: abstract parsing logic into Parser class ` + "\n"

type AgentMethods interface {
	GetCommit(c c.ConfigMethods) (string, error)
	Edit(v string) string
	GetDiff() (string, error)
	MakeCommit(msg string) error
}
