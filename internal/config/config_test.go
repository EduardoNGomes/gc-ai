package config

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/providers"
)

func TestConfig(t *testing.T) {
	allowEdit := "n"
	var output bytes.Buffer

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
		cpath := createTestPath(t)
		conf := NewConfig()

		createTestFile(t, cpath, "", "", providers.GEMINI)

		if err := conf.LoadEnvs(cpath); err != nil {
			t.Errorf("Error on load env ->  %v", err)
		}

		r := conf.GetGeminiKey()

		checkAssert(t, r, "")
	})

	t.Run("Should load envs with values", func(t *testing.T) {
		value := "Key"
		cpath := createTestPath(t)

		createTestFile(t, cpath, value, value, providers.GEMINI)

		conf := NewConfig()

		if err := conf.LoadEnvs(cpath); err != nil {
			t.Errorf("Error on load env ->  %v", err)
		}
		r := conf.GetOpenAIKey()

		checkAssert(t, r, value)
	})

	t.Run("[ConfigGemini]Should register user input GEMINI", func(t *testing.T) {
		gemini := "geminiKey"
		input := bytes.NewBufferString(fmt.Sprintf("%s\n%s", gemini, allowEdit))

		cpath := createTestPath(t)
		createTestFile(t, cpath, "", "", providers.GEMINI)
		c := NewConfig()

		if err := c.LoadEnvs(cpath); err != nil {
			t.Errorf("Error on load env ->  %v", err)
		}

		agents := providers.NewSelectAgentSpy()
		agents.SetAgent(providers.GEMINI)

		c.ConfigKey(input, agents, &output)

		checkAssert(t, c.GetGeminiKey(), gemini)

	})

	t.Run("[ConfigOpenAI]Should register user input OPEN AI", func(t *testing.T) {
		openAI := "openAIKey"
		input := bytes.NewBufferString(fmt.Sprintf("%s\n%s", openAI, allowEdit))

		cpath := createTestPath(t)
		createTestFile(t, cpath, "", "", providers.OPEN_AI)
		c := NewConfig()

		if err := c.LoadEnvs(cpath); err != nil {
			t.Errorf("Error on load env ->  %v", err)
		}

		agents := providers.NewSelectAgentSpy()
		agents.SetAgent(providers.OPEN_AI)
		c.ConfigKey(input, agents, &output)

		checkAssert(t, c.GetOpenAIKey(), openAI)
	})

	t.Run("[SetAllowEdit] should alter edit config", func(t *testing.T) {
		c := NewConfig()

		c.setAllowEdit(true)

		r := c.GetAllowEdit()

		if r != true {
			t.Errorf("Unexpect value, expect %t, receive %t", true, r)
		}
	})
	t.Run("[SetAllowEdit] should alter edit config rewrite file", func(t *testing.T) {
		c := NewConfig()

		cpath := createTestPath(t)
		createTestFile(t, cpath, "", "", providers.GEMINI)

		if err := c.LoadEnvs(cpath); err != nil {
			t.Errorf("Error on load env ->  %v", err)
		}

		firstValue := c.GetAllowEdit()

		if firstValue != false {
			t.Errorf("Unexpect value, expect %t, receive %t", false, firstValue)
		}

		newValue := true

		c.RewriteConfig(RewriteConfigOptions{
			PromptType: nil,
			Prompt:     nil,
			AllowEdit:  &newValue,
			Agent:      nil,
		})
		f, err := os.ReadFile(cpath)

		if err != nil {
			t.Errorf("Err on read file -> %v", err)
		}

		var eTest envStruct

		err = json.Unmarshal(f, &eTest)

		if err := json.Unmarshal(f, &eTest); err != nil {
			t.Errorf("Err decode JSON -> %v", err)
		}

		r := eTest.AllowEdit

		if r != newValue {
			t.Errorf("Err on set new config, expect %t receive %t ", newValue, r)
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

func createTestFile(t *testing.T, p, geminiV, openAIV string, provider providers.Provider) {
	f, err := os.Create(p)

	if err != nil {
		t.Errorf(errs.CannotOpenFileErr+" -> %v", err)
	}

	defer f.Close()

	if err = writeConfig(f, geminiV, openAIV, false, provider); err != nil {
		t.Error(err)
	}
}
