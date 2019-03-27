Notifier
================

[![Build Status](https://travis-ci.org/joaosoft/notifier.svg?branch=master)](https://travis-ci.org/joaosoft/notifier) | [![codecov](https://codecov.io/gh/joaosoft/notifier/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/notifier) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/notifier)](https://goreportcard.com/report/github.com/joaosoft/notifier) | [![GoDoc](https://godoc.org/github.com/joaosoft/notifier?status.svg)](https://godoc.org/github.com/joaosoft/notifier)

A simple notifier tool.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* Slack 

## Dependecy Management
>### Dependency

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Get dependency manager: `go get github.com/joaosoft/dependency`
* Install dependencies: `dependency get`

>### Go
```
go get github.com/joaosoft/notifier
```

>### Configuration
```
{
  "notifier": {
    "log": {
      "level": "info"
    },
    "slack": {
      "webhook": "https://hooks.slack.com/services/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"
    }
  }
}
```

## Usage 
This examples are available in the project at [notifier/examples](https://github.com/joaosoft/notifier/tree/master/examples)

```go
func main() {
	myNotifier, err := notifier.New()
	if err != nil {
		panic(err)
	}

	// notifiers
	slack := myNotifier.NewSlackNotifier()

	// send message with slack notifier
	sendMessage(slack, "hello slack")
}

func sendMessage(notifier notifier.INotifier, message string) {
	fmt.Printf("\nSending message to %s...\n\n", notifier.Name())

	if err := notifier.Notify(message); err != nil {
		panic(err)
	}

	fmt.Printf("\nMessage sent to %s!", notifier.Name())
}
```

> ##### Result:
```
Sending message to slack...

[IN] Method[POST] Url[/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXXXXX] on Start[2019-03-26 21:51:21.943464 +0000 WET m=+0.005795099]
[OUT] Status[200] Method[POST] Url[/services/T9NSHKVA7/BH8H6K15W/UnVtECaobipqK8MmkvC0mJav] on Start[2019-03-26 21:51:21.943464 +0000 WET m=+0.005795099] Elapsed[565.239383ms]

Message send to slack!
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
