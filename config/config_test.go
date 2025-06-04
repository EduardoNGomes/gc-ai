package config

import "testing"

func TestConfig(t *testing.T) {
	t.Run("Should return empty config", func(t *testing.T) {
		conf := &Config{}

		r := conf.isEmpty()

		if r != true {
			t.Errorf("Should be empty but receive: %v", r)
		}
	})
}
