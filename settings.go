package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "os"
        "path/filepath"
)

var (
        settings         Settings
)

type Settings struct {
        Bot_token string `json:"bot_token"`
	News_source string `json:"news_source"`
}

func LoadSettings() {
	ProgramPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
                panic(err)
        }
        jsonFile, err := os.Open(ProgramPath + "/settings.json")
        if err != nil {
                fmt.Printf("Settings file not found, please create one!\n")
        }
        defer jsonFile.Close()
        byteValue, err := ioutil.ReadAll(jsonFile)
        if err != nil {
                fmt.Println(err)
        }
        err = json.Unmarshal(byteValue, &settings)
        if err != nil {
          fmt.Printf("Unable to read program settings: %s\n",err)
          os.Exit(1)
        }
}
