package notifier

import "github.com/joaosoft/errors"

var (
	ErrorNotifierSendMessage = errors.New(errors.LevelError, 1, "error sending message [notifier: %s, error: %s]")
	ErrorNotifierStatus      = errors.New(errors.LevelError, 1, "error sending message [notifier: %s, status: %d, error: %s]")
	ErrorMarshalMessage      = errors.New(errors.LevelError, 1, "error marshall message [notifier: %s, message: %s]")
)
