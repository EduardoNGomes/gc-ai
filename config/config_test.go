package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Shoud return Gemini Key", func(t *testing.T) {
		key := "key"
		conf := &Config{open_ai_key: "", gemini_key: key}

		r := conf.getGeminiKey()

		if r != key {
			t.Errorf("Expect %s, Receive %s", key, r)
		}
	})

	t.Run("Shoud return OpenAi Key", func(t *testing.T) {
		key := "key"
		conf := &Config{open_ai_key: key, gemini_key: ""}

		r := conf.getOpenIAKey()

		checkAssert(t, r, key)

		if r != key {
			t.Errorf("Expect %s, Receive %s", key, r)
		}
	})

	t.Run("Should set Gemini Key", func(t *testing.T) {
		key := "key"

		conf := &Config{}

		conf.setGeminiKey(key)

		r := conf.getGeminiKey()

		checkAssert(t, r, key)

	})

	t.Run("Should set OpenAi Key", func(t *testing.T) {
		key := "key"

		conf := &Config{}

		conf.setOpenIAKey(key)

		r := conf.getOpenIAKey()

		checkAssert(t, r, key)

	})

	t.Run("Should return empty config", func(t *testing.T) {
		conf := &Config{}

		r := conf.isEmpty()

		if r != true {
			t.Errorf("Should be empty but receive: %v", r)
		}
	})

	t.Run("Should load envs with empty values", func(t *testing.T) {
		conf := &Config{}

		conf.loadEnvs()
		r := conf.getGeminiKey()

		checkAssert(t, r, "")

	})

	t.Run("Should load envs with values", func(t *testing.T) {
		value := "Key"
		os.Setenv("OPEN_AI", value)
		conf := &Config{}

		conf.loadEnvs()
		r := conf.getOpenIAKey()

		checkAssert(t, r, value)

	})

}

func checkAssert(t *testing.T, r, e string) {
	t.Helper()

	if e != r {
		t.Errorf("Receive: '%s', Expect: '%s'", r, e)
	}
}
