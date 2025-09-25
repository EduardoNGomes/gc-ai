package agents

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/eduardongomes/gcai/errs"
	c "github.com/eduardongomes/gcai/internal/config"
)

func TestGeminiAgent(t *testing.T) {
	t.Run("[GetDiff]should return erro if not has diff", func(t *testing.T) {
		agent := NewGeminiAgent()

		_, err := agent.GetDiff()

		if !errors.Is(err, errs.EmptyDiffError) {
			t.Errorf("Should not work without diff %v", err)
		}
	})

	t.Run("[GetDiff]should return err on get diff", func(t *testing.T) {
		agent := NewGeminiAgent()

		execCommand = func(name string, args ...string) *exec.Cmd {
			return exec.Command("false")
		}

		defer func() { execCommand = exec.Command }()

		_, err := agent.GetDiff()

		if err == nil {
			t.Error("[Diff]Expect error receive nil")
		}
	})

	t.Run("[GetDiff] Should not get error on get diff", func(t *testing.T) {
		agent := NewGeminiAgent()

		execCommand = func(name string, args ...string) *exec.Cmd {
			return exec.Command("echo", "fake-diff-success")
		}

		defer func() { execCommand = exec.Command }()

		_, err := agent.GetDiff()

		if err != nil {
			t.Errorf("Error on get diff %v", err)
		}

	})

	t.Run("[GetCommit]Should receive error if has wrong api_key", func(t *testing.T) {
		agent := NewGeminiAgent()
		c := c.NewConfSpy()

		c.SetGeminiKey("")

		_, err := agent.GetCommit(c)

		if err == nil {
			t.Errorf("Should not work with fake key")
		}

		if !errors.Is(err, errs.EmptyKeyError) {
			t.Errorf("Should not work without empty key %v", err)
		}

	})

	t.Run("[MakeCommit] Should  get error on try make commit", func(t *testing.T) {
		agent := NewGeminiAgent()

		execCommand = func(name string, args ...string) *exec.Cmd {
			return exec.Command("false")
		}

		defer func() { execCommand = exec.Command }()

		err := agent.MakeCommit("fake diff")
		if err == nil {
			t.Error("[MakeCommit] -> Expect error receive nil")
		}
	})

	t.Run("[MakeCommit] Should not get error on try make commit", func(t *testing.T) {
		agent := NewGeminiAgent()

		execCommand = func(name string, args ...string) *exec.Cmd {
			return exec.Command("echo", "fake-commit-success")
		}

		defer func() { execCommand = exec.Command }()

		err := agent.MakeCommit("fake diff")

		if err != nil {
			t.Errorf("Error on make commit %v", err)
		}

	})
}
