package config

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Shoud return Gemini Key", func(t *testing.T) {
		key := "key"
		conf := &Config{open_ai_key: "", gemini_key: key}

		r := conf.GetGeminiKey()

		if r != key {
			t.Errorf("Expect %s, Receive %s", key, r)
		}
	})

	t.Run("Shoud return OpenAi Key", func(t *testing.T) {
		key := "key"
		conf := &Config{open_ai_key: key, gemini_key: ""}

		r := conf.GetOpenAIKey()

		checkAssert(t, r, key)

		if r != key {
			t.Errorf("Expect %s, Receive %s", key, r)
		}
	})

	t.Run("Should set Gemini Key", func(t *testing.T) {
		key := "key"

		conf := NewConfig()

		conf.setGeminiKey(key)

		r := conf.GetGeminiKey()

		checkAssert(t, r, key)

	})

	t.Run("Should set OpenAi Key", func(t *testing.T) {
		key := "key"

		conf := NewConfig()

		conf.setOpenAIKey(key)

		r := conf.GetOpenAIKey()

		checkAssert(t, r, key)

	})

	t.Run("Should return empty config", func(t *testing.T) {
		conf := NewConfig()

		r := conf.IsEmpty()

		if r != true {
			t.Errorf("Should be empty but receive: %v", r)
		}
	})

	t.Run("Should load envs with empty values", func(t *testing.T) {
		conf := NewConfig()

		conf.LoadEnvs()
		r := conf.GetGeminiKey()

		checkAssert(t, r, "")

	})

	t.Run("Should load envs with values", func(t *testing.T) {
		value := "Key"
		os.Setenv("OPEN_AI", value)

		conf := NewConfig()

		conf.LoadEnvs()
		r := conf.GetOpenAIKey()

		checkAssert(t, r, value)
	})

	t.Run("Should register user input", func(t *testing.T) {
		openAI := "openAIKey"
		gemini := "geminiKey"
		input := bytes.NewBufferString(fmt.Sprintf("%s\n%s\n", openAI, gemini))

		c := NewConfig()

		c.ConfigKey(input)

		checkAssert(t, c.GetGeminiKey(), gemini)
		checkAssert(t, c.GetOpenAIKey(), openAI)
	})
}

func checkAssert(t *testing.T, r, e string) {
	t.Helper()

	if e != r {
		t.Errorf("Receive: '%s', Expect: '%s'", r, e)
	}
}
