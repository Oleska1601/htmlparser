package logging

//go:generate go run github.com/vektra/mockery/v2@latest --name=LoggerInterface
type LoggerInterface interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
}
