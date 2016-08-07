package main

import (
    "log"
    "os"
    "io/ioutil"
    "gopkg.in/yaml.v2"

    "github.com/go-chat-bot/bot/irc"
    "github.com/go-chat-bot/bot/slack"
    "github.com/go-chat-bot/bot/telegram"
    _ "github.com/go-chat-bot/plugins-br/cnpj"
    _ "github.com/go-chat-bot/plugins-br/cotacao"
    _ "github.com/go-chat-bot/plugins-br/cpf"
    _ "github.com/go-chat-bot/plugins-br/dilma"
    _ "github.com/go-chat-bot/plugins-br/lula"
    _ "github.com/go-chat-bot/plugins-br/megasena"
    _ "github.com/go-chat-bot/plugins/9gag"
    _ "github.com/go-chat-bot/plugins/catfacts"
    _ "github.com/go-chat-bot/plugins/catgif"
    _ "github.com/go-chat-bot/plugins/chucknorris"
    _ "github.com/go-chat-bot/plugins/crypto"
    _ "github.com/go-chat-bot/plugins/encoding"
    _ "github.com/go-chat-bot/plugins/example"
    _ "github.com/go-chat-bot/plugins/gif"
    _ "github.com/go-chat-bot/plugins/godoc"
    _ "github.com/go-chat-bot/plugins/guid"
    _ "github.com/go-chat-bot/plugins/puppet"
    _ "github.com/go-chat-bot/plugins/treta"
    _ "github.com/go-chat-bot/plugins/url"
)

type Config struct {
    Debug bool
    SlackToken string
    TelegramToken string
    IRC *irc.Config
}

func main() {
    config := loadConfig()

    go irc.Run(config.IRC)
    go slack.Run(config.SlackToken)
    telegram.Run(config.TelegramToken, config.Debug)
}

func loadConfig() Config {
    filename := os.Args[1]
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    log.Printf("%s\n", source)

    var config Config
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }
    log.Printf("Configuration loaded: %v\n", config)

    return config
}
