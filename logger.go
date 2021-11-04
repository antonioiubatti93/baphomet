package baphomet

// Logger defines a minimal logger.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}
