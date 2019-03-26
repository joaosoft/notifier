package notifier

import (
	"sync"

	"github.com/joaosoft/web"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Notifier struct {
	config        *NotifierConfig
	webClient     *web.Client
	logger        logger.ILogger
	isLogExternal bool
	pm            *manager.Manager
	mux           sync.Mutex
}

type INotifier interface {
	Name() string
	Notify(message string) error
}

// New ...
func New(options ...NotificationsOption) (*Notifier, error) {
	config, simpleConfig, err := NewConfig()

	webClient, err := web.NewClient()
	if err != nil {
		return nil, err
	}

	service := &Notifier{
		pm:        manager.NewManager(manager.WithRunInBackground(true)),
		logger:    logger.NewLogDefault("notifier", logger.WarnLevel),
		config:    config.Notifier,
		webClient: webClient,
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Notifier != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Notifier.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	service.Reconfigure(options...)

	return service, nil
}

func (n *Notifier) NewSlackNotifier(config ...*SlackConfig) INotifier {
	c := n.config.Slack
	if len(config) > 0 {
		c = config[0]
	}

	return newNotifierSlack(n, c)
}
