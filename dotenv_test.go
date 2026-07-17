package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDotEnv(t *testing.T) {
	// should set env to "development" and only load ".env"
	LoadDotEnv(`fixtures`)
	assert.Equal(t, `foo`, os.Getenv(`VALUE1`), `VALUE1 should be "foo"`)
	assert.Equal(t, `world`, os.Getenv(`SECTION2_VALUE`), `SECTION2_VALUE should be "world"`)

	unsertDotEnvVars()

	SetEnv(true, Testing)
	// should set env to "testing" and load ".env" & ".env.testing"
	LoadDotEnv(`fixtures`)
	assert.Equal(t, `dlrow`, os.Getenv(`SECTION2_VALUE`), `SECTION2_VALUE should be "dlrow"`)

	unsertDotEnvVars()
}

func TestLoadDotEnvPanic(t *testing.T) {
	assert.Panics(t, func() {
		LoadDotEnv(`fixtures`, `env.ini`)
	})

	unsertDotEnvVars()
}

func unsertDotEnvVars() {
	os.Unsetenv(`ENV`)
	os.Unsetenv(`VALUE1`)
	os.Unsetenv(`VALUE2`)
	os.Unsetenv(`SECTION1_VALUE`)
	os.Unsetenv(`SECTION2_VALUE`)
}
