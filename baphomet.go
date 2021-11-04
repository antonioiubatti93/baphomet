package baphomet

import (
	abraxaslog "github.com/antonioiubatti93/abraxas/logger"
	"github.com/antonioiubatti93/baphomet/logger"
)

// Summon summons Baphomet.
func Summon(logger logger.Logger) {
	logger.Print("Hello, I'm Baphomet!")
}

func Mediate(logger abraxaslog.Logger) {
	logger.Print("Summoning Baphomet now")
	Summon(logger)
}
