package notifier

import (
	"encoding/json"

	"github.com/joaosoft/web"
)

type notifierSlack struct {
	name    string
	service *Notifier
	config  *SlackConfig
}

type SlackConfig struct {
	Webhook string `json:"webhook"`
}

type slackMessage struct {
	Text string `json:"text"`
}

func newNotifierSlack(service *Notifier, config *SlackConfig) *notifierSlack {
	return &notifierSlack{
		name:    notifierSlackName,
		service: service,
		config:  config,
	}
}

func (n *notifierSlack) Name() string {
	return n.name
}

func (n *notifierSlack) Notify(message string) error {
	request, err := n.service.webClient.NewRequest(web.MethodPost, n.config.Webhook)
	if err != nil {
		return ErrorNotifierSendMessage.Format(n.name, err)
	}

	body, err := json.Marshal(&slackMessage{
		Text: message,
	})

	if err != nil {
		return ErrorMarshalMessage.Format(n.name, message)
	}

	response, err := request.WithBody(body, web.ContentTypeApplicationJSON).Send()
	if err != nil {
		return ErrorNotifierSendMessage.Format(n.name, err)
	}

	if response.Status >= web.StatusBadRequest {
		return ErrorNotifierStatus.Format(n.name, response.Status, response.StatusText)
	}

	return nil
}
