package agents

import (
	"fmt"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/config"
	"github.com/eduardongomes/gcai/internal/menu"
)

type AgentMethods interface {
	GetCommit(c config.ConfigMethods) (string, error)
}

func GetDiff() (string, error) {
	diff := execCommand("git", "diff", "--cached")

	stdout, err := diff.Output()

	if err != nil {
		return "", fmt.Errorf("Error on get git diff: %w", err)
	}

	d := string(stdout)

	if len(d) == 0 {
		return "", errs.EmptyDiffError

	}

	return d, nil
}

func Edit(menu menu.MenuEditable, msg string) (string, error) {
	m := fmt.Sprintf("Please edit your commit message or press Enter to keep it unchanged:")

	menuEditable := menu

	line, err := menuEditable.Run(m, msg)

	if err != nil {
		return "", fmt.Errorf("error reading line: %v", err)
	}

	return line, nil
}

func MakeCommit(msg string) error {

	r := execCommand("git", "commit", "-m", msg)

	if _, err := r.Output(); err != nil {
		return fmt.Errorf("Erro on make commit: %v", err)
	}

	return nil
}
