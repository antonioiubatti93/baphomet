package baphomet

import (
	"testing"

	log "github.com/sirupsen/logrus"
	logt "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_Summon(t *testing.T) {
	logger := log.New()
	hook := logt.NewLocal(logger)

	Summon(logger)

	assert.Equal(t, "Hello, I'm Baphomet!", hook.LastEntry().Message)
}

func Test_Mediate(t *testing.T) {
	logger := log.New()
	hook := logt.NewLocal(logger)

	Mediate(logger)

	assert.Equal(t, "Hello, I'm Baphomet!", hook.LastEntry().Message)
}
