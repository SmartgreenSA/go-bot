package main

import (
	"fmt"
	"github.com/go-chat-bot/bot/irc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configuration interface {
	LoadFromFile(filename string) error
}

type Config struct {
	Debug         bool
	SlackToken    string
	TelegramToken string
	IRC           *irc.Config
}

func (c *Config) LoadFromFile(filename string) error {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(source, c)
	if err != nil {
		return err
	}

	log.Printf("Configuration loaded: %s\n", c)

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("{%t %s %s %v}", c.Debug, c.SlackToken, c.TelegramToken, c.IRC)
}
