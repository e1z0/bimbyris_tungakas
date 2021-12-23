package main

import (
	"fmt"
        "io/ioutil"
        "os"
        "path/filepath"
	"strings"
        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
    Dummy_Messages []DummyMessage
)

type DummyMessage struct {
        Question string
        Answer string
}


func ReadMessages() {
	ProgramPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
                panic(err)
        }
        MessagesFile, err := os.Open(ProgramPath + "/messages.txt")
        if err != nil {
                fmt.Printf("messages.txt not found, please create one!\n")
                return
        }
        defer MessagesFile.Close()
        byteValue, err := ioutil.ReadAll(MessagesFile)
        if err != nil {
                fmt.Println(err)
        }
        content := string(byteValue)
        lines := strings.Split(content,"\n")
        count := 0
        for _,ln := range lines {
             if ln != "" {
             count++
             split := strings.Split(ln, "::")
             itm := DummyMessage{Question: split[0], Answer: split[1]}
             Dummy_Messages = append(Dummy_Messages,itm)
             }
        }
        fmt.Printf("%d Messages from messages.txt loaded to the memory\n",count)
}

func DummyMessages(bot *tgbotapi.BotAPI,chat_id int64, text string) {
	fmt.Printf("Parsing text: %s from %d\n",text,chat_id)
	var msg string
	for _, tpl := range Dummy_Messages {
	if tpl.Question == text {
		msg = tpl.Answer
	}
	}

	if (msg != "") {
		m := tgbotapi.NewMessage(chat_id, msg)
		bot.Send(m)
	}
}
