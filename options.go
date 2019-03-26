package notifier

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// NotificationsOption ...
type NotificationsOption func(client *Notifier)

// Reconfigure ...
func (n *Notifier) Reconfigure(options ...NotificationsOption) {
	for _, option := range options {
		option(n)
	}
}

// WithConfiguration ...
func WithConfiguration(config *NotifierConfig) NotificationsOption {
	return func(n *Notifier) {
		n.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) NotificationsOption {
	return func(n *Notifier) {
		n.logger = logger
		n.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) NotificationsOption {
	return func(n *Notifier) {
		n.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) NotificationsOption {
	return func(n *Notifier) {
		n.pm = mgr
	}
}