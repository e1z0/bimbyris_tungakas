all: build
deps:
	go get github.com/go-telegram-bot-api/telegram-bot-api
	go get github.com/ssimunic/gosensors
	go get github.com/mmcdole/gofeed
	go get github.com/mmcdole/gofeed/atom
	go get github.com/mmcdole/gofeed/rss
build:
	go build -o tg-bot main.go utils.go rss.go settings.go cmd.go dummy_messages.go
