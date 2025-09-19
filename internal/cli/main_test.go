package cli

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/agents"
	c "github.com/eduardongomes/gcai/internal/config"
)

func TestCLI(t *testing.T) {
	t.Run("Should call config method on start method when config is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		cli := NewCLI()
		agent := agents.NewMockAgent()

		cli.Run(false, confSpy, agent, &bytes.Buffer{})

		expect := true
		result := confSpy.IsEmptyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})

	t.Run("Should call config key if confif is empty", func(t *testing.T) {
		cli := NewCLI()
		confSpy := c.NewConfSpy()

		agent := agents.NewMockAgent()
		cli.Run(false, confSpy, agent, &bytes.Buffer{})
		expect := true
		result := confSpy.IsConfigKeyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})
}

func checkAssert(t *testing.T, r, e string) {
	t.Helper()

	if e != r {
		t.Errorf("Receive: '%s', Expect: '%s'", r, e)
	}
}

func createTestPath(t *testing.T) string {
	t.Helper()

	hash := md5.Sum([]byte(t.Name()))
	id := hex.EncodeToString(hash[:8])

	return fmt.Sprintf("./.config-test-%s.json", id)
}

func createTestFileWithEnvs(t *testing.T, p, openai, gemini string) {
	f, err := os.Create(p)
	if err != nil {
		t.Fatalf(errs.CannotOpenFileErr+" -> %v", err)
	}
	defer f.Close()

	data := struct {
		GeminiKey string
		OpenAIKey string
	}{
		GeminiKey: gemini,
		OpenAIKey: openai,
	}

	dataByte, err := json.Marshal(data)
	if err != nil {
		t.Fatalf(errs.ErrorOnConvertDataToByte+" -> %v", err)
	}

	if _, err = f.Write(dataByte); err != nil {
		t.Fatalf(errs.ErrorOnWriteFileConfig+" -> %v", err)
	}
}
