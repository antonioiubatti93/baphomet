package baphomet

import (
	"github.com/antonioiubatti93/abraxas"
	abraxaslog "github.com/antonioiubatti93/abraxas/logger"
	"github.com/antonioiubatti93/baphomet/logger"
)

// Summon summons Baphomet.
func Summon(logger logger.Logger) {
	logger.Print("Hello, I'm Baphomet!")
}

func Mediate(logger abraxaslog.Logger) {
	abraxas.Invoke(logger)
	logger.Print("Summoning Baphomet now")
	Summon(logger)
}
