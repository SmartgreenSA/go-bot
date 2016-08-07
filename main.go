package main

import (
    "os"

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

func main() {
    filename := os.Args[1]
    
    var config Config
    err := config.LoadFromFile(filename)
    if err != nil {
        panic(err)
    }

    go irc.Run(config.IRC)
    go slack.Run(config.SlackToken)
    telegram.Run(config.TelegramToken, config.Debug)
}
