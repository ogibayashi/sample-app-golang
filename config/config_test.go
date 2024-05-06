package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setTestConfig() error {
	var testConfig = []byte(`
kafka:
  bootstrap: localhost:9092
  user: testuser
`)
	return c.ReadConfig(bytes.NewBuffer(testConfig))
}

func TestConfig(t *testing.T) {
	err := setTestConfig()
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, "localhost:9092", GetString("kafka.bootstrap"))

	t.Setenv("APP_FOO_BAR", "foo")
	assert.Equal(t, "foo", GetString("foo.bar"))

}
