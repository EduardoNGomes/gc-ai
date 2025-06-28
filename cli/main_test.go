package cli

import (
	"bytes"
	"os"
	"testing"

	c "github.com/eduardongomes/gcai/config"
)

func TestCLI(t *testing.T) {
	t.Run("Should load envs on Start method", func(t *testing.T) {
		value := "Key"
		os.Setenv("OPEN_AI", value)

		conf := c.NewConfig()
		cli := NewCLI()

		cli.Run(conf, &bytes.Buffer{})

		conf.GetOpenAIKey()

		r := conf.GetOpenAIKey()

		checkAssert(t, r, value)
	})

	t.Run("Should call config method on start method when env is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		cli := NewCLI()

		cli.Run(confSpy, &bytes.Buffer{})

		expect := true
		result := confSpy.IsEmptyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})

	t.Run("Should call config key if confif is empty", func(t *testing.T) {
		cli := NewCLI()
		confSpy := c.NewConfSpy()

		cli.Run(confSpy, &bytes.Buffer{})
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
